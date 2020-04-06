package pkg

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"github.com/rs/cors"
	"github.com/stripe/stripe-go/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type PaymentService struct {
	Cfg          *Config
	PGDB         *pgdb
	Logger       *zap.SugaredLogger
	StripeClient *client.API
}

// NewServer is a factory for creating a struct containing our critical service elements
func NewPaymentService(logger *zap.SugaredLogger, cfg *Config) PaymentService {

	// Initialize database connection
	db, err := NewPGDB(cfg, logger)
	FatalIfError(logger, err, "failed to connect to pgdb")

	// Initialize Stripe client
	sc, err := NewStripeClient(cfg)
	FatalIfError(logger, err, "failed to validate Stripe client connection")

	return PaymentService{
		Cfg:          cfg,
		PGDB:         db,
		Logger:       logger,
		StripeClient: sc,
	}
}

func FatalIfError(appLogger *zap.SugaredLogger, err error, msg string) {
	if err != nil {
		appLogger.Fatalw(msg,
			"error", err.Error())
	}
}

func (ps *PaymentService) RunEndpoints() error {
	ps.Logger.Infow("RunEndpoints Port",
		"ps.Cfg.RPCPort", ps.Cfg.RPCPort)
	lis, err := net.Listen("tcp", ps.Cfg.RPCPort)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	paymentProto.RegisterPaymentServiceServer(server, ps)

	go server.Serve(lis)
	return nil
}

func (ps *PaymentService) RunGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"Authorization", "Content-Type", "Accept"},
	}).Handler(mux)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	ps.Logger.Infow("RunGateway Port",
		"ps.Cfg.RPCPort", ps.Cfg.RPCPort)
	err := paymentProto.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, ps.Cfg.RPCPort, opts)
	if err != nil {
		return err
	}

	ps.Logger.Infow("RunEndpoints ListenAndServe Port",
		"ps.Cfg.JSONPort", ps.Cfg.JSONPort)
	http.ListenAndServe(ps.Cfg.JSONPort, handler)
	return nil
}
