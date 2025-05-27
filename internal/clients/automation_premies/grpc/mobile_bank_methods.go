package grpc

import (
	"context"
	"fmt"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"premiesPortal/internal/app/grpc/gen/go/mobile_bank"
)

func (c *Client) UploadMobileBankData(ctx context.Context, in *mobile_bank.MobileBankUploadRequest) (*mobile_bank.MobileBankUploadResponse, error) {
	const op = "grpc.UploadMobileBankData"

	resp, err := c.MobileBankApi.UploadMobileBankData(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (c *Client) CleanMobileBankTable(ctx context.Context, in *emptypb.Empty) (*mobile_bank.MobileBankCleanResponse, error) {
	const op = "grpc.CleanMobileBankTable"

	resp, err := c.MobileBankApi.CleanMobileBankTable(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
