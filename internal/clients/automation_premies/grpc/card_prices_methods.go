package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"premiesPortal/internal/app/grpc/gen/go/card_prices"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func (c *Client) UploadCardPricesData(ctx context.Context, in *card_prices.CardPricesUploadRequest) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.UploadCardPricesData"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.CardPricesApi.UploadCardPricesData(ctx, in)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			// Это gRPC-ошибка
			logger.Error.Printf("[%s] gRPC error code: %v, message: %v\n", op, st.Code(), st.Message())
			if st.Code() == codes.NotFound {
				return nil, errs.ErrFileNotFound
			}
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp.Resp = respGrpc.Status

	return &resp, nil
}
