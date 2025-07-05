package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	upl "premiesPortal/internal/app/grpc/gen/upload_file"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func (c *Client) UploadFile(ctx context.Context, in *upl.UploadFileRequest) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.UploadFile"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.UploadFileApi.UploadFile(ctx, in)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			// Это gRPC-ошибка
			logger.Error.Printf("[%s] gRPC error code: %v, message: %v\n", op, st.Code(), st.Message())
			if st.Code() == codes.Internal {
				return nil, errs.ErrSomethingWentWrong
			}

			if st.Code() == codes.NotFound {
				return nil, errs.ErrTempReportNotFound
			}
		}

		logger.Error.Printf("[%s] Error calling UploadFile: %v", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp.Resp = respGrpc.Path

	return &resp, nil
}

func (c *Client) DownloadFile(ctx context.Context, in *upl.DownloadFileRequest) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.DownloadFile"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.UploadFileApi.DownloadFile(ctx, in)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			// Это gRPC-ошибка
			logger.Error.Printf("[%s] gRPC error code: %v, message: %v\n", op, st.Code(), st.Message())
			if st.Code() == codes.Internal {
				return nil, errs.ErrSomethingWentWrong
			}

			if st.Code() == codes.NotFound {
				return nil, errs.ErrTempReportNotFound
			}
		}

		logger.Error.Printf("[%s] Error calling UploadFile: %v", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp.Resp = respGrpc.GetFileContent()

	return &resp, nil
}
