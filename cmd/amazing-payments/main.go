package main

import (
	"fmt"
	"log"

	"github.com/ianspire/amazing-payments/pkg"
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
	defer PanicRecovery(appLogger)

	// Obtain config from env variables
	cfg := pkg.NewConfig()
	appLogger.Infow("loaded config",
		"config", cfg,
	)

	service := pkg.NewPaymentService(appLogger, cfg)

	appLogger.Infof("starting grpc on port %s", cfg.RPCPort)
	err = service.RunEndpoints()
	pkg.FatalIfError(appLogger, err, "failed to start grpc server")

	appLogger.Infof("starting grpc-gateway on port %s", cfg.JSONPort)
	err = service.RunGateway()
	pkg.FatalIfError(appLogger, err, "failed to start grpc-gateway")
}

func PanicRecovery(appLogger *zap.SugaredLogger) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("#{r}")
		}
		appLogger.Errorf("failed due to uncaught error",
			"error", err)
		panic(err)

	}
}
