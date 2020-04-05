package main

import (
	"github.com/ianspire/amazing-payments/pkg"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"google.golang.org/grpc"
	"log"

	"go.uber.org/zap"
)

func main() {

	// Initialize logging
	baseLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}
	defer baseLogger.Sync()
	appLogger := baseLogger.Sugar()

	// Obtain config from env variables
	cfg := pkg.NewConfig()
	appLogger.Infow("loaded config",
		"config", cfg,
	)

	// Initialize database connection
	db, err := pkg.NewDB(cfg, appLogger)
	if err != nil {
		appLogger.Fatalw("failed to connect to PGDB",
			"error", err,
		)
	}

	// Initialize Stripe client
	sc, err := pkg.NewStripeClient(cfg)
	if err != nil {
		appLogger.Fatalw("failed to validate Stripe client connection",
			"error", err,
		)
	}

	// Build the service
	s := pkg.NewServer(db, appLogger, sc)
	server := grpc.NewServer()
	paymentProto.RegisterPaymentServiceServer(server, &s)

	if err := pkg.StartRESTServer(cfg.JSONPort, cfg.RPCPort); err != nil {
		appLogger.Fatalw("failed to start REST Server",
			"error", err,
			)
	}
}
