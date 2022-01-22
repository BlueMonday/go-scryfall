package scryfall

import (
	"context"
	"fmt"
)

// Catalog contains an array of Magic datapoints (words, card values,
// etc). Catalog objects are provided by the API as aids for building other Magic
// software and understanding possible values for a field on Card objects.
type Catalog struct {
	// URI is a link to the current catalog on Scryfall's API.
	URI string `json:"uri"`

	// TotalValues is the number of items in the data array.
	TotalValues int `json:"total_values"`

	// Data is an array of datapoints, as strings.
	Data []string `json:"data"`
}

func (c *Client) getCatalog(ctx context.Context, name string) (Catalog, error) {
	catalogURL := fmt.Sprintf("catalog/%s", name)
	catalog := Catalog{}
	err := c.get(ctx, catalogURL, &catalog)
	if err != nil {
		return Catalog{}, err
	}

	return catalog, nil
}

// GetCardNamesCatalog returns a list of all nontoken English card names in
// Scryfall's database. Values are updated as soon as a new card is entered for
// spoiler seasons.
func (c *Client) GetCardNamesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "card-names")
}

// GetArtistNamesCatalog returns a list of all canonical artist names in
// Scryfall's database. This catalog won't include duplicate, misspelled, or funny
// names for artists. Values are updated as soon as a new card is entered for
// spoiler seasons.
func (c *Client) GetArtistNamesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "artist-names")
}

// GetWordBankCatalog returns a Catalog of all English words, of length 2 or
// more, that could appear in a card name. Values are drawn from cards currently
// in Scryfall's database. Values are updated as soon as a new card is entered for
// spoiler seasons.
func (c *Client) GetWordBankCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "word-bank")
}

// GetCreatureTypesCatalog returns a Catalog of all creature types in
// Scryfall's database. Values are updated as soon as a new card is entered for
// spoiler seasons.
func (c *Client) GetCreatureTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "creature-types")
}

// GetPlaneswalkerTypesCatalog returns a Catalog of all Planeswalker types in
// Scryfall's database. Values are updated as soon as a new card is entered for
// spoiler seasons.
func (c *Client) GetPlaneswalkerTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "planeswalker-types")
}

// GetLandTypesCatalog returns a Catalog of all Land types in Scryfall's
// database. Values are updated as soon as a new card is entered for spoiler
// seasons.
func (c *Client) GetLandTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "land-types")
}

// GetArtifactTypesCatalog returns a Catalog of all artifact types in
// Scryfall's database. Values are updated as soon as a new card is entered for
// spoiler seasons.
func (c *Client) GetArtifactTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "artifact-types")
}

// GetEnchantmentTypesCatalog returns a Catalog of all enchantment types in
// Scryfall's database. Values are updated as soon as a new card is entered for
// spoiler seasons.
func (c *Client) GetEnchantmentTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "enchantment-types")
}

// GetSpellTypesCatalog returns a Catalog of all spell types in Scryfall's
// database. Values are updated as soon as a new card is entered for spoiler
// seasons.
func (c *Client) GetSpellTypesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "spell-types")
}

// GetPowersCatalog returns a Catalog of all possible values for a creature or
// vehicle's power in Scryfall's database. Values are updated as soon as a new
// card is entered for spoiler seasons.
func (c *Client) GetPowersCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "powers")
}

// GetToughnessesCatalog returns a Catalog of all possible values for a
// creature or vehicle's toughness in Scryfall's database. Values are updated as
// soon as a new card is entered for spoiler seasons.
func (c *Client) GetToughnessesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "toughnesses")
}

// GetLoyaltiesCatalog returns a Catalog of all possible values for a
// Planeswalker's loyalty in Scryfall's database. Values are updated as soon as a
// new card is entered for spoiler seasons.
func (c *Client) GetLoyaltiesCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "loyalties")
}

// GetWatermarksCatalog returns a Catalog of all card watermarks in Scryfall's
// database. Values are updated as soon as a new card is entered for spoiler
// seasons.
func (c *Client) GetWatermarksCatalog(ctx context.Context) (Catalog, error) {
	return c.getCatalog(ctx, "watermarks")
}
