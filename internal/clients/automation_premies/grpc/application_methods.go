package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"premiesPortal/internal/app/grpc/gen/application"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func (c *Client) CreateXLSXApplicationReport(ctx context.Context, in *application.CreateXLSXApplicationRequest) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.CreateXLSXApplicationReport"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.ApplicationApi.CreateXLSXApplications(ctx, in)
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

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp.Resp = respGrpc.GetXlsxPath()

	return &resp, nil
}
