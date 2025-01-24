package scryfall

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	qs "github.com/google/go-querystring/query"
)

type Lang string

const (
	LangEnglish            Lang = "en"
	LangSpanish            Lang = "es"
	LangFrench             Lang = "fr"
	LangGerman             Lang = "de"
	LangItalian            Lang = "it"
	LangPortuguese         Lang = "pt"
	LangJapanese           Lang = "ja"
	LangKorean             Lang = "ko"
	LangRussian            Lang = "ru"
	LangSimplifiedChinese  Lang = "zhs"
	LangTraditionalChinese Lang = "zht"
	LangHebrew             Lang = "he"
	LangLatin              Lang = "la"
	LangAncientGreek       Lang = "grc"
	LangArabic             Lang = "ar"
	LangSanskrit           Lang = "sa"
	LangPhyrexian          Lang = "ph"
)

type Layout string

const (
	LayoutNormal           Layout = "normal"
	LayoutSplit            Layout = "split"
	LayoutFlip             Layout = "flip"
	LayoutTransform        Layout = "transform"
	LayoutModalDFC         Layout = "modal_dfc"
	LayoutMeld             Layout = "meld"
	LayoutLeveler          Layout = "leveler"
	LayoutClass            Layout = "class"
	LayoutCase             Layout = "case"
	LayoutSaga             Layout = "saga"
	LayoutAdventure        Layout = "adventure"
	LayoutMutate           Layout = "mutate"
	LayoutPrototype        Layout = "prototype"
	LayoutBattle           Layout = "battle"
	LayoutPlanar           Layout = "planar"
	LayoutScheme           Layout = "scheme"
	LayoutVanguard         Layout = "vanguard"
	LayoutToken            Layout = "token"
	LayoutDoubleFacedToken Layout = "double_faced_token"
	LayoutEmblem           Layout = "emblem"
	LayoutAugment          Layout = "augment"
	LayoutHost             Layout = "host"
	LayoutArtSeries        Layout = "art_series"
	LayoutReversible       Layout = "reversible_card"
)

type Legality string

const (
	LegalityLegal      Legality = "legal"
	LegalityNotLegal   Legality = "not_legal"
	LegalityBanned     Legality = "banned"
	LegalityRestricted Legality = "restricted"
)

type Frame string

const (
	Frame1993   Frame = "1993"
	Frame1997   Frame = "1997"
	Frame2003   Frame = "2003"
	Frame2015   Frame = "2015"
	FrameFuture Frame = "future"
)

type FrameEffect string

const (
	FrameEffectLegendary              FrameEffect = "legendary"
	FrameEffectMiracle                FrameEffect = "miracle"
	FrameEffectEnchantment            FrameEffect = "enchantment"
	FrameEffectDraft                  FrameEffect = "draft"
	FrameEffectDevoid                 FrameEffect = "devoid"
	FrameEffectTombstone              FrameEffect = "tombstone"
	FrameEffectColorShifted           FrameEffect = "colorshifted"
	FrameEffectInverted               FrameEffect = "inverted"
	FrameEffectSunMoonDFC             FrameEffect = "sunmoondfc"
	FrameEffectCompassLandDFC         FrameEffect = "compasslanddfc"
	FrameEffectOriginPWDFC            FrameEffect = "originpwdfc"
	FrameEffectMoonEldraziDFC         FrameEffect = "mooneldrazidfc"
	FrameEffectWaxingAndWaningMoonDFC FrameEffect = "waxingandwaningmoondfc"
	FrameEffectShowcase               FrameEffect = "showcase"
	FrameEffectExtendedArt            FrameEffect = "extendedart"
	FrameEffectCompanion              FrameEffect = "companion"
	FrameEffectEtched                 FrameEffect = "etched"
	FrameEffectSnow                   FrameEffect = "snow"
	FrameEffectLesson                 FrameEffect = "lesson"
	FrameEffectShatteredGlass         FrameEffect = "shatteredglass"
	FrameEffectConvertDFC             FrameEffect = "convertdfc"
	FrameEffectFanDFC                 FrameEffect = "fandfc"
	FrameEffectUpsideDownDFC          FrameEffect = "upsidedowndfc"
	FrameEffectSpree                  FrameEffect = "spree"
)

type Preview struct {
	PreviewedAt Date   `json:"previewed_at"`
	SourceURI   string `json:"source_uri"`
	Source      string `json:"source"`
}

type Component string

const (
	ComponentToken      Component = "token"
	ComponentMeldPart   Component = "meld_part"
	ComponentMeldResult Component = "meld_result"
	ComponentComboPiece Component = "combo_piece"
)

type Rarity string

const (
	RarityCommon   Rarity = "common"
	RarityUncommon Rarity = "uncommon"
	RarityRare     Rarity = "rare"
	RaritySpecial  Rarity = "special"
	RarityMythic   Rarity = "mythic"
	RarityBonus    Rarity = "bonus"
)

type Finish string

const (
	FinishFoil    Finish = "foil"
	FinishNonFoil Finish = "nonfoil"
	FinishEtched  Finish = "etched"
)

type ImageStatus string

const (
	ImageStatusMissing    ImageStatus = "missing"
	ImageStatusPlaceholer ImageStatus = "placeholder"
	ImageStatusLowres     ImageStatus = "lowres"
	ImageStatusHighres    ImageStatus = "highres_scan"
)

type Card struct {
	ArenaID              *int          `json:"arena_id,omitempty"`
	ID                   string        `json:"id"`
	Lang                 Lang          `json:"lang"`
	OracleID             string        `json:"oracle_id"`
	MultiverseIDs        []int         `json:"multiverse_ids"`
	MTGOID               *int          `json:"mtgo_id,omitempty"`
	MTGOFoilID           *int          `json:"mtgo_foil_id,omitempty"`
	URI                  string        `json:"uri"`
	ScryfallURI          string        `json:"scryfall_uri"`
	TCGPlayerID          *int          `json:"tcgplayer_id,omitempty"`
	TCGPlayerEtchedID    *int          `json:"tcgplayer_etched_id,omitempty"`
	CardMarketID         *int          `json:"Integer,omitempty"`
	PrintsSearchURI      string        `json:"prints_search_uri"`
	RulingsURI           string        `json:"rulings_uri"`
	Name                 string        `json:"name"`
	PrintedName          *string       `json:"printed_name"`
	Layout               Layout        `json:"layout"`
	CMC                  float64       `json:"cmc"`
	TypeLine             string        `json:"type_line"`
	PrintedTypeLine      *string       `json:"printed_type_line"`
	OracleText           string        `json:"oracle_text"`
	PrintedText          *string       `json:"printed_text"`
	ManaCost             string        `json:"mana_cost"`
	Power                *string       `json:"power"`
	Toughness            *string       `json:"toughness"`
	Loyalty              *string       `json:"loyalty"`
	Defense              *string       `json:"defense"`
	LifeModifier         *string       `json:"life_modifier"`
	HandModifier         *string       `json:"hand_modifier"`
	Colors               []Color       `json:"colors"`
	ColorIndicator       []Color       `json:"color_indicator"`
	ColorIdentity        []Color       `json:"color_identity"`
	AllParts             []RelatedCard `json:"all_parts"`
	CardFaces            []CardFace    `json:"card_faces"`
	Legalities           Legalities    `json:"legalities"`
	Reserved             bool          `json:"reserved"`
	Foil                 bool          `json:"foil"`
	NonFoil              bool          `json:"nonfoil"`
	Oversized            bool          `json:"oversized"`
	Promo                bool          `json:"promo"`
	EDHRECRank           *int          `json:"edhrec_rank"`
	Set                  string        `json:"set"`
	SetName              string        `json:"set_name"`
	SetID                string        `json:"set_id"`
	CollectorNumber      string        `json:"collector_number"`
	SetURI               string        `json:"set_uri"`
	SetSearchURI         string        `json:"set_search_uri"`
	ScryfallSetURI       string        `json:"scryfall_set_uri"`
	ImageURIs            *ImageURIs    `json:"image_uris"`
	Prices               Prices        `json:"prices"`
	ReleasedAt           Date          `json:"released_at"`
	HighresImage         bool          `json:"highres_image"`
	Reprint              bool          `json:"reprint"`
	Digital              bool          `json:"digital"`
	Rarity               string        `json:"rarity"`
	FlavorText           *string       `json:"flavor_text"`
	Artist               *string       `json:"artist"`
	IllustrationID       *string       `json:"illustration_id"`
	Frame                Frame         `json:"frame"`
	FrameEffects         []FrameEffect `json:"frame_effects"`
	FullArt              bool          `json:"full_art"`
	Watermark            *string       `json:"watermark"`
	Preview              Preview       `json:"preview"`
	PromoTypes           []string      `json:"promo_types,omitempty"`
	BorderColor          string        `json:"border_color"`
	StorySpotlightNumber *int          `json:"story_spotlight_number"`
	StorySpotlightURI    *string       `json:"story_spotlight_uri"`
	RelatedURIs          RelatedURIs   `json:"related_uris"`
	PurchaseURIs         PurchaseURIs  `json:"purchase_uris"`
	Keywords             []string      `json:"keywords"`
	ProducedMana         []Color       `json:"produced_mana"`
	Booster              bool          `json:"booster"`
	Finishes             []Finish      `json:"finishes"`
	ImageStatus          *ImageStatus  `json:"image_status"`
	AttractionLights     []int         `json:"attraction_lights,omitempty"`
	ContentWarning       *bool         `json:"content_warning,omitempty"`
	FlavorName           *string       `json:"flavor_name,omitempty"`
}

type RelatedCard struct {
	ID        string    `json:"id"`
	Component Component `json:"component"`
	Name      string    `json:"name"`
	TypeLine  string    `json:"type_line"`
	URI       string    `json:"uri"`
}

type CardFace struct {
	Name            string    `json:"name"`
	PrintedName     *string   `json:"printed_name"`
	TypeLine        string    `json:"type_line"`
	PrintedTypeLine *string   `json:"printed_type_line"`
	OracleText      *string   `json:"oracle_text"`
	PrintedText     *string   `json:"printed_text"`
	ManaCost        string    `json:"mana_cost"`
	Colors          []Color   `json:"colors"`
	ColorIndicator  []Color   `json:"color_indicator"`
	Power           *string   `json:"power"`
	Toughness       *string   `json:"toughness"`
	Layout          *Layout   `json:"layout"`
	Loyalty         *string   `json:"loyalty"`
	OracleID        *string   `json:"oracle_id,omitempty"`
	Defense         *string   `json:"defense"`
	FlavorText      *string   `json:"flavor_text"`
	IllustrationID  *string   `json:"illustration_id"`
	ImageURIs       ImageURIs `json:"image_uris"`
}

type ImageURIs struct {
	Small      *string `json:"small"`
	Normal     *string `json:"normal"`
	Large      *string `json:"large"`
	PNG        *string `json:"png"`
	ArtCrop    *string `json:"art_crop"`
	BorderCrop *string `json:"border_crop"`
}

type Prices struct {
	USD       string `json:"usd"`
	USDFoil   string `json:"usd_foil"`
	USDEtched string `json:"usd_etched"`
	EUR       string `json:"eur"`
	EURFoil   string `json:"eur_foil"`
	Tix       string `json:"tix"`
}

type Legalities struct {
	Standard  Legality `json:"standard"`
	Modern    Legality `json:"modern"`
	Pauper    Legality `json:"pauper"`
	Pioneer   Legality `json:"pioneer"`
	Legacy    Legality `json:"legacy"`
	Penny     Legality `json:"penny"`
	Vintage   Legality `json:"vintage"`
	Duel      Legality `json:"duel"`
	Commander Legality `json:"commander"`
	Future    Legality `json:"future"`
}

type RelatedURIs struct {
	Gatherer       string `json:"gatherer"`
	TCGPlayerDecks string `json:"tcgplayer_decks"`
	EDHREC         string `json:"edhrec"`
	MTGTop8        string `json:"mtgtop8"`
}

type PurchaseURIs struct {
	TCGPlayer   string `json:"tcgplayer"`
	CardMarket  string `json:"cardmarket"`
	CardHoarder string `json:"cardhoarder"`
}

type UniqueMode string

const (
	UniqueModeCards  UniqueMode = "cards"
	UniqueModeArt    UniqueMode = "art"
	UniqueModePrints UniqueMode = "prints"
)

type Order string

const (
	OrderName      Order = "name"
	OrderSet       Order = "set"
	OrderRarity    Order = "rarity"
	OrderColor     Order = "color"
	OrderUSD       Order = "usd"
	OrderTix       Order = "tix"
	OrderEUR       Order = "eur"
	OrderCMC       Order = "cmc"
	OrderPower     Order = "power"
	OrderToughness Order = "toughness"
	OrderEDHREC    Order = "edhrec"
	OrderArtist    Order = "artist"
)

type Dir string

const (
	DirAuto Dir = "auto"
	DirAsc  Dir = "asc"
	DirDesc Dir = "desc"
)

type SearchCardsOptions struct {
	Unique              UniqueMode `url:"unique,omitempty"`
	Order               Order      `url:"order,omitempty"`
	Dir                 Dir        `url:"dir,omitempty"`
	IncludeExtras       bool       `url:"include_extras,omitempty"`
	IncludeMultilingual bool       `url:"include_multilingual,omitempty"`
	IncludeVariations   bool       `url:"include_variations,omitempty"`
	Page                int        `url:"page,omitempty"`
}
type CardListResponse struct {
	Cards      []Card   `json:"data"`
	HasMore    bool     `json:"has_more"`
	NextPage   *string  `json:"next_page"`
	TotalCards int      `json:"total_cards"`
	Warnings   []string `json:"warnings"`
}

func (c *Client) SearchCards(ctx context.Context, query string, opts SearchCardsOptions) (CardListResponse, error) {
	values, err := qs.Values(opts)
	if err != nil {
		return CardListResponse{}, err
	}
	values.Set("q", query)
	cardsURI := fmt.Sprintf("cards/search?%s", values.Encode())
	result := CardListResponse{}
	err = c.get(ctx, cardsURI, &result)
	if err != nil {
		return CardListResponse{}, err
	}
	return result, nil
}

func (c *Client) getCard(ctx context.Context, url string) (Card, error) {
	card := Card{}
	err := c.get(ctx, url, &card)
	if err != nil {
		return Card{}, err
	}
	return card, nil
}

type GetCardByNameOptions struct {
	Set string `url:"set,omitempty"`
}

func (c *Client) GetCardByName(ctx context.Context, name string, exact bool, opts GetCardByNameOptions) (Card, error) {
	values, err := qs.Values(opts)
	if err != nil {
		return Card{}, err
	}
	if exact {
		values.Set("exact", name)
	} else {
		values.Set("fuzzy", name)
	}
	cardURI := fmt.Sprintf("cards/named?%s", values.Encode())
	return c.getCard(ctx, cardURI)
}

func (c *Client) AutocompleteCard(ctx context.Context, s string) ([]string, error) {
	values := url.Values{}
	values.Set("q", s)
	autocompleteCardURI := fmt.Sprintf("cards/autocomplete?%s", values.Encode())
	catalog := Catalog{}
	err := c.get(ctx, autocompleteCardURI, &catalog)
	if err != nil {
		return nil, err
	}
	return catalog.Data, nil
}

func (c *Client) GetRandomCard(ctx context.Context) (Card, error) {
	return c.getCard(ctx, "cards/random")
}

type CardIdentifier struct {
	ID              string `json:"id,omitempty"`
	MTGOID          int    `json:"mtgo_id,omitempty"`
	MultiverseID    int    `json:"multiverse_id,omitempty"`
	Name            string `json:"name,omitempty"`
	Set             string `json:"set,omitempty"`
	CollectorNumber string `json:"collector_number,omitempty"`
}

type GetCardsByIdentifiersRequest struct {
	Identifiers []CardIdentifier `json:"identifiers"`
}

type GetCardsByIdentifiersResponse struct {
	NotFound []CardIdentifier `json:"not_found"`
	Data     []Card           `json:"data"`
}

func (c *Client) GetCardsByIdentifiers(ctx context.Context, identifiers []CardIdentifier) (GetCardsByIdentifiersResponse, error) {
	getCardsByIdentifiersRequest := GetCardsByIdentifiersRequest{
		Identifiers: identifiers,
	}
	getCardsByIdentifiersResponse := GetCardsByIdentifiersResponse{}
	err := c.post(ctx, "cards/collection", &getCardsByIdentifiersRequest, &getCardsByIdentifiersResponse)
	if err != nil {
		return GetCardsByIdentifiersResponse{}, err
	}
	return getCardsByIdentifiersResponse, nil
}

func (c *Client) GetCardBySetCodeAndCollectorNumber(ctx context.Context, setCode string, collectorNumber string) (Card, error) {
	cardURI := fmt.Sprintf("cards/%s/%s", setCode, collectorNumber)
	return c.getCard(ctx, cardURI)
}

func (c *Client) GetCardBySetCodeAndCollectorNumberInLang(ctx context.Context, setCode string, collectorNumber string, lang Lang) (Card, error) {
	cardURI := fmt.Sprintf("cards/%s/%s/%s", setCode, collectorNumber, lang)
	return c.getCard(ctx, cardURI)
}

func (c *Client) GetCardByMultiverseID(ctx context.Context, multiverseID int) (Card, error) {
	cardURI := fmt.Sprintf("cards/multiverse/%d", multiverseID)
	return c.getCard(ctx, cardURI)
}

func (c *Client) GetCardByMTGOID(ctx context.Context, mtgoID int) (Card, error) {
	cardURI := fmt.Sprintf("cards/mtgo/%d", mtgoID)
	return c.getCard(ctx, cardURI)
}

func (c *Client) GetCardByArenaID(ctx context.Context, arenaID int) (Card, error) {
	cardURI := fmt.Sprintf("cards/arena/%d", arenaID)
	return c.getCard(ctx, cardURI)
}

func (c *Client) GetCardByTCGPlayerID(ctx context.Context, tcgPlayerID int) (Card, error) {
	cardURI := fmt.Sprintf("cards/tcgplayer/%d", tcgPlayerID)
	return c.getCard(ctx, cardURI)
}

func (c *Client) GetCard(ctx context.Context, id string) (Card, error) {
	cardURI := fmt.Sprintf("cards/%s", id)
	return c.getCard(ctx, cardURI)
}

func (c *Client) GetCardFromURI(ctx context.Context, uri string) (Card, error) {
	parts := strings.Split(uri, "/")
	if len(parts) < 6 {
		return Card{}, fmt.Errorf("invalid card url format")
	}
	setCode := parts[len(parts)-3]
	collectorNumber := parts[len(parts)-2]

	cardURI := fmt.Sprintf("cards/%s/%s", setCode, collectorNumber)
	return c.getCard(ctx, cardURI)
}
