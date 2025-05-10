package grpc

import (
	"context"
	"fmt"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"premiesPortal/internal/app/grpc/gen/go/cards"
)

func (c *Client) UploadCards(ctx context.Context, in *cards.CardsUploadRequest) (*cards.CardsUploadResponse, error) {
	const op = "grpc.UploadCards"

	resp, err := c.CardsApi.UploadCardsData(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (c *Client) CleanCards(ctx context.Context, in *emptypb.Empty) (*cards.CardsCleanResponse, error) {
	const op = "grpc.CleanCards"

	resp, err := c.CardsApi.CleanCardsTable(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
