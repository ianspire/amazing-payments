package main

import (
	"log"

	"github.com/ianspire/amazing-payments/pkg/v1/config"

	"go.uber.org/zap"
)

func main() {

	// Initialize logging
	baseLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}
	defer baseLogger.Sync()
	logger := baseLogger.Sugar()

	cfg := config.InitializeConfig()
	logger.Infow("loaded config",
		"config", cfg,
	)

	logger.Info("Payment Service is running! Hello world!")

}
