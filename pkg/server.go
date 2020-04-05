package pkg

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"github.com/stripe/stripe-go/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type PaymentService struct {
	PGDB         *PGDB
	Logger       *zap.SugaredLogger
	StripeClient *client.API
}

// NewServer is a factory for creating a struct containing our critical service elements
func NewServer(d *PGDB, l *zap.SugaredLogger, s *client.API) PaymentService {
	return PaymentService{
		PGDB:         d,
		Logger:       l,
		StripeClient: s,
	}
}

// starts the REST Service handler
func StartRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()

	// Setup the client gRPC options
	var opts []grpc.DialOption

	// Register Payment Service Handler
	err := paymentProto.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register payment service: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

// RunEndpoints starts the gRPC server endpoint - in case you'd like to make that available
func RunEndpoints(ps *PaymentService, jsonPort string) error {
	lis, err := net.Listen("tcp", jsonPort)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	paymentProto.RegisterPaymentServiceServer(server, ps)

	go server.Serve(lis)
	return nil
}
