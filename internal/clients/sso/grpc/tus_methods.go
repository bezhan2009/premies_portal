package grpc

import (
	"context"
	"fmt"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"premiesPortal/internal/app/grpc/gen/go/tus"
)

func (c *Client) UploadTusData(ctx context.Context, in *tus.TusUploadRequest) (*tus.TusUploadResponse, error) {
	const op = "grpc.UploadTusData"

	resp, err := c.TusApi.UploadTusData(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (c *Client) CleanTusTable(ctx context.Context, in *emptypb.Empty) (*tus.TusCleanResponse, error) {
	const op = "grpc.CleanTusTable"

	resp, err := c.TusApi.CleanTusTable(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
