package pkg

import (
	"context"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/client"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func setLogger() {
	baseLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}
	logger = baseLogger.Sugar()
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

func setupDatastore(t *testing.T) *MockDatastore {
	setLogger()

	ctrl := gomock.NewController(t)
	pgdb := NewMockDatastore(ctrl)
	pgdb.EXPECT().
		HealthCheck().
		Return(nil)

	return pgdb
}

func TestPaymentService_HealthCheck(t *testing.T) {
	ctx := context.Background()
	pgdb := setupDatastore(t)

	s := PaymentService{
		Cfg:          &Config{},
		PGDB:         pgdb,
		Logger:       logger,
		StripeClient: &client.API{},
	}
	req := paymentProto.HealthCheckRequest{}
	resp, err := s.HealthCheck(ctx, &req)
	assert.Nil(t, err)
	assert.Equal(t, &paymentProto.HealthCheckResponse{IsHealthy: true}, resp)
}
