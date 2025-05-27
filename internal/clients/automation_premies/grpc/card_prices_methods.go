package grpc

import (
	"context"
	"fmt"
	"premiesPortal/internal/app/grpc/gen/go/card_prices"
)

func (c *Client) UploadCardPricesData(ctx context.Context, in *card_prices.CardPricesUploadRequest) (*card_prices.CardPricesUploadResponse, error) {
	const op = "grpc.UploadCardPricesData"

	resp, err := c.CardPricesApi.UploadCardPricesData(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
