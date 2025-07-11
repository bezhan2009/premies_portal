package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"premiesPortal/internal/app/grpc/gen/reports"
	"premiesPortal/internal/app/models"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
)

func (c *Client) CreateZIPReports(ctx context.Context, in *reports.CreateZIPReportsRequest) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.CreateZIPReports"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.ReportApi.CreateZIPReports(ctx, in)
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

	resp.Resp = respGrpc.GetZipPath()

	return &resp, nil
}

func (c *Client) CreateExcelReport(ctx context.Context, in *reports.CreateExcelReportRequest) (*models.ResponseWithStatusCode, error) {
	const op = "grpc.CreateExcelReport"

	resp := models.ResponseWithStatusCode{
		StatusCode: http.StatusOK,
	}

	respGrpc, err := c.ReportApi.CreateExcelReport(ctx, in)
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

	resp.Resp = respGrpc.GetZipPath()

	return &resp, nil
}
