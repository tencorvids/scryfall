package scryfall

import (
	"context"
	"fmt"
)

type BulkData struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	UpdatedAt       Timestamp `json:"updated_at"`
	Name            string    `json:"name"`
	URI             string    `json:"uri"`
	Description     string    `json:"description"`
	Size            int       `json:"size"`
	DownloadURI     string    `json:"download_uri"`
	ContentType     string    `json:"content_type"`
	ContentEncoding string    `json:"content_encoding"`
}

func (c *Client) ListBulkData(ctx context.Context) ([]BulkData, error) {
	bulkDataItems := []BulkData{}
	err := c.listGet(ctx, "bulk-data", &bulkDataItems)
	if err != nil {
		return nil, err
	}
	return bulkDataItems, nil
}

func (c *Client) GetBulkDataByID(ctx context.Context, id string) (BulkData, error) {
	bulkDataURI := fmt.Sprintf("bulk-data/%s", id)
	bulkData := BulkData{}
	err := c.get(ctx, bulkDataURI, &bulkData)
	if err != nil {
		return BulkData{}, err
	}
	return bulkData, nil
}

func (c *Client) GetBulkDataByType(ctx context.Context, typ string) (BulkData, error) {
	bulkDataURI := fmt.Sprintf("bulk-data/%s", typ)
	bulkData := BulkData{}
	err := c.get(ctx, bulkDataURI, &bulkData)
	if err != nil {
		return BulkData{}, err
	}
	return bulkData, nil
}
