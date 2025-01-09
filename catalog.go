package scryfall

import (
	"context"
	"fmt"
)

type Catalog struct {
	URI         string   `json:"uri"`
	TotalValues int      `json:"total_values"`
	Data        []string `json:"data"`
}

func (c *Client) getCatalog(ctx context.Context, name string) (Catalog, error) {
	catalogURI := fmt.Sprintf("catalog/%s", name)
	catalog := Catalog{}
	err := c.get(ctx, catalogURI, &catalog)
	if err != nil {
		return Catalog{}, err
	}
	return catalog, nil
}

func (c *Client) GetCardNamesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "card-names")
}

func (c *Client) GetArtistNamesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "artist-names")
}

func (c *Client) GetWordBankCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "word-bank")
}

func (c *Client) GetSuperTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "supertypes")
}

func (c *Client) GetCardTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "card-types")
}

func (c *Client) GetCreatureTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "creature-types")
}

func (c *Client) GetPlaneswalkerTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "planeswalker-types")
}

func (c *Client) GetLandTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "land-types")
}

func (c *Client) GetArtifactTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "artifact-types")
}

func (c *Client) GetBattleTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "battle-types")
}

func (c *Client) GetEnchantmentTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "enchantment-types")
}

func (c *Client) GetSpellTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "spell-types")
}

func (c *Client) GetPowersCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "powers")
}

func (c *Client) GetToughnessesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "toughnesses")
}

func (c *Client) GetLoyaltiesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "loyalties")
}

func (c *Client) GetKeywordAbilitiesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "keyword-abilities")
}

func (c *Client) GetKeywordActionsCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "keyword-actions")
}

func (c *Client) GetAbilityWordsCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "ability-words")
}

func (c *Client) GetFlavorWordsCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "flavor-words")
}

func (c *Client) GetWatermarksCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "watermarks")
}
