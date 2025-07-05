package grpc

import (
	"context"
	"fmt"
	accountantServ "premiesPortal/internal/app/grpc/gen/accountant"
	cardPricesServ "premiesPortal/internal/app/grpc/gen/card_prices"
	cardsServ "premiesPortal/internal/app/grpc/gen/cards"
	mobileBankServ "premiesPortal/internal/app/grpc/gen/mobile_bank"
	reportServ "premiesPortal/internal/app/grpc/gen/reports"
	tusServ "premiesPortal/internal/app/grpc/gen/tus"
	uplServ "premiesPortal/internal/app/grpc/gen/upload_file"
	"premiesPortal/pkg/logger"

	"google.golang.org/grpc/codes"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	CardsApi      cardsServ.CardsServiceClient
	MobileBankApi mobileBankServ.MobileBankServiceClient
	TusApi        tusServ.TusServiceClient
	CardPricesApi cardPricesServ.CardPricesServiceClient
	ReportApi     reportServ.ReportsServiceClient
	UploadFileApi uplServ.UploadFileServiceClient
	AccountantApi accountantServ.AccountantsServiceClient
}

var client *Client
var conn *grpc.ClientConn

func New(ctx context.Context,
	addr string,
	timeout time.Duration,
	retriesCount int,
) (err error) {
	const op = "grpc.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	conn, err = grpc.DialContext(ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOpts...)),
	)
	if err != nil {
		logger.Error.Printf("[%s]: dial failed: %s", op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	client = &Client{
		CardsApi:      cardsServ.NewCardsServiceClient(conn),
		MobileBankApi: mobileBankServ.NewMobileBankServiceClient(conn),
		TusApi:        tusServ.NewTusServiceClient(conn),
		CardPricesApi: cardPricesServ.NewCardPricesServiceClient(conn),
		ReportApi:     reportServ.NewReportsServiceClient(conn),
		UploadFileApi: uplServ.NewUploadFileServiceClient(conn),
		AccountantApi: accountantServ.NewAccountantsServiceClient(conn),
	}

	return nil
}

func GetClient() *Client {
	return client
}

func GrpcConnClose() (err error) {
	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}
