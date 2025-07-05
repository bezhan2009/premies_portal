package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"premiesPortal/internal/app/grpc/gen/mobile_bank"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func (c *Client) UploadMobileBankData(ctx context.Context, in *mobile_bank.MobileBankUploadRequest) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.UploadMobileBankData"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.MobileBankApi.UploadMobileBankData(ctx, in)
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

	resp.Resp = respGrpc.GetStatus()

	return &resp, nil
}

func (c *Client) CleanMobileBankTable(ctx context.Context, in *emptypb.Empty) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.CleanMobileBankTable"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.MobileBankApi.CleanMobileBankTable(ctx, in)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			// Это gRPC-ошибка
			logger.Error.Printf("[%s] gRPC error code: %v, message: %v\n", op, st.Code(), st.Message())
			if st.Code() == codes.Internal {
				return nil, errs.ErrSomethingWentWrong
			}
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp.Resp = respGrpc.GetStatus()

	return &resp, nil
}
