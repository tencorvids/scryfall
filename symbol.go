package scryfall

import (
	"context"
	"fmt"
	"net/url"
)

type CardSymbol struct {
	Symbol               string      `json:"symbol"`
	SVGURI               string      `json:"svg_uri"`
	ManaValue            float64     `json:"mana_value"`
	Hybrid               bool        `json:"hybrid"`
	Phyrexian            bool        `json:"phyrexian"`
	GathererAlternatives interface{} `json:"gatherer_alternatives"`
	LooseVariant         *string     `json:"loose_variant"`
	English              string      `json:"english"`
	Transposable         bool        `json:"transposable"`
	RepresentsMana       bool        `json:"represents_mana"`
	CMC                  float64     `json:"cmc"`
	AppearsInManaCosts   bool        `json:"appears_in_mana_costs"`
	Funny                bool        `json:"funny"`
	Colors               []Color     `json:"colors"`
}

type ManaCost struct {
	Cost         string  `json:"cost"`
	CMC          float64 `json:"cmc"`
	Colors       []Color `json:"colors"`
	Colorless    bool    `json:"colorless"`
	Monocolored  bool    `json:"monocolored"`
	Multicolored bool    `json:"multicolored"`
}

func (c *Client) ListCardSymbols(ctx context.Context) ([]CardSymbol, error) {
	symbols := []CardSymbol{}
	err := c.listGet(ctx, "symbology", &symbols)
	if err != nil {
		return nil, err
	}
	return symbols, nil
}

func (c *Client) ParseManaCost(ctx context.Context, cost string) (ManaCost, error) {
	values := url.Values{}
	values.Set("cost", cost)
	parseManaURI := fmt.Sprintf("symbology/parse-mana?%s", values.Encode())
	manaCost := ManaCost{}
	err := c.get(ctx, parseManaURI, &manaCost)
	if err != nil {
		return ManaCost{}, err
	}
	return manaCost, nil
}
