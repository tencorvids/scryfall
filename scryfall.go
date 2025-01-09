package scryfall

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/ratelimit"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Version             = "0.0.0"
	defaultBaseURI      = "https://api.scryfall.com"
	defaultUserAgent    = "github.com/tencorvids/scryfall/" + Version
	defaultTimeout      = 30 * time.Second
	defaultReqPerSecond = 10
	dateFormat          = "2006-01-02"
	timestampFormat     = "2006-01-02T15:04:05.999Z07:00"
)

var ErrMultipleSecrets = errors.New("multiple secrets configured")

type Color string

const (
	ColorWhite Color = "W"
	ColorBlue  Color = "U"
	ColorBlack Color = "B"
	ColorRed   Color = "R"
	ColorGreen Color = "G"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	parsedTime, err := time.ParseInLocation(dateFormat, s, time.FixedZone("UTC-8", -8*60*60))
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Format(dateFormat))), nil
}

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	parsedTime, err := time.Parse(timestampFormat, s)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

type Error struct {
	Status   int      `json:"status"`
	Code     string   `json:"code"`
	Details  string   `json:"details"`
	Type     *string  `json:"type"`
	Warnings []string `json:"warnings"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Details)
}

type clientOptions struct {
	baseURI      string
	userAgent    string
	clientSecret string
	grantSecret  string
	client       *http.Client
	limiter      ratelimit.Limiter
}

type ClientOption func(*clientOptions)

func WithBaseURI(baseURI string) ClientOption {
	return func(o *clientOptions) {
		o.baseURI = baseURI
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(o *clientOptions) {
		o.userAgent = userAgent
	}
}

func WithClientSecret(clientSecret string) ClientOption {
	return func(o *clientOptions) {
		o.clientSecret = clientSecret
	}
}

func WithGrantSecret(grantSecret string) ClientOption {
	return func(o *clientOptions) {
		o.grantSecret = grantSecret
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(o *clientOptions) {
		o.client = client
	}
}

func WithLimiter(limiter ratelimit.Limiter) ClientOption {
	return func(o *clientOptions) {
		o.limiter = limiter
	}
}

type Client struct {
	baseURI       *url.URL
	userAgent     string
	authorization string
	client        *http.Client
	limiter       ratelimit.Limiter
}

func NewClient(options ...ClientOption) (*Client, error) {
	co := &clientOptions{
		baseURI:   defaultBaseURI,
		userAgent: defaultUserAgent,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		limiter: ratelimit.New(defaultReqPerSecond),
	}
	for _, option := range options {
		option(co)
	}
	if len(co.clientSecret) != 0 && len(co.grantSecret) != 0 {
		return nil, ErrMultipleSecrets
	}
	var authorization string
	if len(co.clientSecret) != 0 {
		authorization = "Bearer " + co.clientSecret
	}
	if len(co.grantSecret) != 0 {
		authorization = "Bearer " + co.grantSecret
	}
	baseURI, err := url.Parse(co.baseURI)
	if err != nil {
		return nil, err
	}
	c := &Client{
		baseURI:       baseURI,
		userAgent:     co.userAgent,
		authorization: authorization,
		client:        co.client,
		limiter:       co.limiter,
	}
	return c, nil
}
func (c *Client) doReq(ctx context.Context, req *http.Request, respBody interface{}) error {
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")
	if len(c.authorization) != 0 {
		req.Header.Set("Authorization", c.authorization)
	}
	reqWithContext := req.WithContext(ctx)
	if c.limiter != nil {
		c.limiter.Take()
	}
	resp, err := c.client.Do(reqWithContext)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		scryfallErr := &Error{}
		err = decoder.Decode(scryfallErr)
		if err != nil {
			return err
		}
		return scryfallErr
	}
	return decoder.Decode(respBody)
}
func (c *Client) get(ctx context.Context, relativeURI string, respBody interface{}) error {
	absoluteURI, err := c.baseURI.Parse(relativeURI)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodGet, absoluteURI.String(), nil)
	if err != nil {
		return err
	}
	return c.doReq(ctx, req, respBody)
}
func (c *Client) post(ctx context.Context, relativeURI string, reqBody interface{}, respBody interface{}) error {
	absoluteURI, err := c.baseURI.Parse(relativeURI)
	if err != nil {
		return err
	}
	var body io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest(http.MethodPost, absoluteURI.String(), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.doReq(ctx, req, respBody)
}

type listResponse struct {
	Data       json.RawMessage `json:"data"`
	HasMore    bool            `json:"has_more"`
	NextPage   *string         `json:"next_page"`
	TotalCards *int            `json:"total_cards"`
	Warnings   []string        `json:"warnings"`
}

func (c *Client) listGet(ctx context.Context, url string, v interface{}) error {
	response := &listResponse{}
	err := c.get(ctx, url, response)
	if err != nil {
		return err
	}
	return json.Unmarshal(response.Data, v)
}
