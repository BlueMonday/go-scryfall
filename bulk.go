package scryfall

import "context"

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
//
// Note: Card objects in bulk data do not contain prices, and will omit the
// USD, EUR, Tix, and purchase URIs properties.
func (c *Client) ListBulkData(ctx context.Context) ([]BulkData, error) {
	bulkDataItems := []BulkData{}
	err := c.listGet(ctx, "bulk-data", &bulkDataItems)
	if err != nil {
		return nil, err
	}

	return bulkDataItems, nil
}
