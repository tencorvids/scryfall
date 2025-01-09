package scryfall

import (
	"context"
	"fmt"
)

type SetType string

const (
	SetTypeCore            SetType = "core"
	SetTypeExpansion       SetType = "expansion"
	SetTypeAlchemy         SetType = "alchemy"
	SetTypeMasters         SetType = "masters"
	SetTypeMasterpiece     SetType = "masterpiece"
	SetTypeArsenal         SetType = "arsenal"
	SetTypeFromTheVault    SetType = "from_the_vault"
	SetTypeSpellbook       SetType = "spellbook"
	SetTypePremiumDeck     SetType = "premium_deck"
	SetTypeDuelDeck        SetType = "duel_deck"
	SetTypeDraftInnovation SetType = "draft_innovation"
	SetTypeTreasureChest   SetType = "treasure_chest"
	SetTypeCommander       SetType = "commander"
	SetTypePlanechase      SetType = "planechase"
	SetTypeArchenemy       SetType = "archenemy"
	SetTypeVanguard        SetType = "vanguard"
	SetTypeFunny           SetType = "funny"
	SetTypeStarter         SetType = "starter"
	SetTypeBox             SetType = "box"
	SetTypePromo           SetType = "promo"
	SetTypeToken           SetType = "token"
	SetTypeMemorabilia     SetType = "memorabilia"
	SetTypeMiniGame        SetType = "minigame"
)

type Set struct {
	ID            string  `json:"id"`
	Code          string  `json:"code"`
	MTGOCode      *string `json:"mtgo_code"`
	ArenaCode     *string `json:"arena_code"`
	TCGplayerID   *int    `json:"tcgplayer_id"`
	Name          string  `json:"name"`
	URI           string  `json:"uri"`
	ScryfallURI   string  `json:"scryfall_uri"`
	SetType       SetType `json:"set_type"`
	ReleasedAt    *Date   `json:"released_at"`
	BlockCode     *string `json:"block_code"`
	Block         *string `json:"block"`
	ParentSetCode string  `json:"parent_set_code"`
	CardCount     int     `json:"card_count"`
	PrintedSize   *int    `json:"printed_size"`
	Digital       bool    `json:"digital"`
	FoilOnly      bool    `json:"foil_only"`
	NonfoilOnly   bool    `json:"nonfoil_only"`
	IconSVGURI    string  `json:"icon_svg_uri"`
	SearchURI     string  `json:"search_uri"`
}

func (c *Client) ListSets(ctx context.Context) ([]Set, error) {
	sets := []Set{}
	err := c.listGet(ctx, "sets", &sets)
	if err != nil {
		return nil, err
	}
	return sets, nil
}

func (c *Client) GetSet(ctx context.Context, code string) (Set, error) {
	setURI := fmt.Sprintf("sets/%s", code)
	set := Set{}
	err := c.get(ctx, setURI, &set)
	if err != nil {
		return Set{}, err
	}
	return set, nil
}
