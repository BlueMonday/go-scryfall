package scryfall

import (
	"context"
	"fmt"
)

// BulkData is a Scryfall bulk data item.
type BulkData struct {
	// ID is a unique ID for this bulk item.
	ID string `json:"id"`

	// Type is a computer-readable string for the kind of bulk item.
	Type string `json:"type"`

	// UpdatedAt is the time when this file was last updated.
	UpdatedAt Timestamp `json:"updated_at"`

	// Name is a human-readable name for this file.
	Name string `json:"name"`

	// URI is a link to this bulk object on Scryfall's API.
	URI string `json:"uri"`

	// Description is a human-readable name for this file.
	Description string `json:"description"`

	// CompressedSize is the compressed size of this file in integer bytes.
	CompressedSize int `json:"compressed_size"`

	// DownloadURI is the URL that hosts this bulk file.
	DownloadURI string `json:"download_uri"`

	// ContentType is the MIME type of this file.
	ContentType string `json:"content_type"`

	// ContentEncoding is the Content-Encoding encoding that will be used
	// to transmit this file when you download it.
	ContentEncoding string `json:"content_encoding"`
}

// ListBulkData returns a list of all bulk data items on Scryfall.
func (c *Client) ListBulkData(ctx context.Context) ([]BulkData, error) {
	bulkDataItems := []BulkData{}
	err := c.listGet(ctx, "bulk-data", &bulkDataItems)
	if err != nil {
		return nil, err
	}

	return bulkDataItems, nil
}

// GetBulkDataByID gets a bulk data item by ID.
func (c *Client) GetBulkDataByID(ctx context.Context, id string) (BulkData, error) {
	bulkDataURL := fmt.Sprintf("bulk-data/%s", id)
	bulkData := BulkData{}
	err := c.get(ctx, bulkDataURL, &bulkData)
	if err != nil {
		return BulkData{}, err
	}

	return bulkData, nil
}

// GetBulkDataByType gets a bulk data item by type.
func (c *Client) GetBulkDataByType(ctx context.Context, typ string) (BulkData, error) {
	bulkDataURL := fmt.Sprintf("bulk-data/%s", typ)
	bulkData := BulkData{}
	err := c.get(ctx, bulkDataURL, &bulkData)
	if err != nil {
		return BulkData{}, err
	}

	return bulkData, nil
}
