package grpc

import (
	"context"
	"fmt"
	cardPricesServ "premiesPortal/internal/app/grpc/gen/go/card_prices"
	cardsServ "premiesPortal/internal/app/grpc/gen/go/cards"
	mobileBankServ "premiesPortal/internal/app/grpc/gen/go/mobile_bank"
	reportServ "premiesPortal/internal/app/grpc/gen/go/reports"
	tusServ "premiesPortal/internal/app/grpc/gen/go/tus"
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
}

var client *Client
var conn *grpc.ClientConn

func New(ctx context.Context,
	addr string,
	timeout time.Duration,
	retriesCount int,
) error {
	const op = "grpc.New"

	var err error

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
