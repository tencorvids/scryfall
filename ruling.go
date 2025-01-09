package scryfall

import (
	"context"
	"fmt"
)

type Source string

const (
	SourceWOTC     Source = "wotc"
	SourceScryfall Source = "scryfall"
)

type Ruling struct {
	OracleID    string `json:"oracle_id"`
	Source      Source `json:"source"`
	PublishedAt Date   `json:"published_at"`
	Comment     string `json:"comment"`
}

func (c *Client) getRulings(ctx context.Context, url string) ([]Ruling, error) {
	rulings := []Ruling{}
	err := c.listGet(ctx, url, &rulings)
	if err != nil {
		return nil, err
	}
	return rulings, nil
}

func (c *Client) GetRulingsByMultiverseID(ctx context.Context, multiverseID int) ([]Ruling, error) {
	rulingsURI := fmt.Sprintf("cards/multiverse/%d/rulings", multiverseID)
	return c.getRulings(ctx, rulingsURI)
}

func (c *Client) GetRulingsByMTGOID(ctx context.Context, mtgoID int) ([]Ruling, error) {
	rulingsURI := fmt.Sprintf("cards/mtgo/%d/rulings", mtgoID)
	return c.getRulings(ctx, rulingsURI)
}

func (c *Client) GetRulingsByArenaID(ctx context.Context, arenaID int) ([]Ruling, error) {
	rulingsURI := fmt.Sprintf("cards/arena/%d/rulings", arenaID)
	return c.getRulings(ctx, rulingsURI)
}

func (c *Client) GetRulingsBySetCodeAndCollectorNumber(ctx context.Context, setCode string, collectorNumber int) ([]Ruling, error) {
	rulingsURI := fmt.Sprintf("cards/%s/%d/rulings", setCode, collectorNumber)
	return c.getRulings(ctx, rulingsURI)
}

func (c *Client) GetRulings(ctx context.Context, id string) ([]Ruling, error) {
	rulingsURI := fmt.Sprintf("cards/%s/rulings", id)
	return c.getRulings(ctx, rulingsURI)
}
