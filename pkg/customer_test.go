package pkg

import (
	"context"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/client"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"testing"
	"time"
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

func setupDatastore() *MockDatastore {
	setLogger()
	var pgdb *MockDatastore
	var ctx context.Context

	pgdb.EXPECT().
		HealthCheck().
		Return(nil)

	var customerID int64
	var name, email, stripeChargeDate, customerKey string
	customerKey = RandomString(12)

	pgdb.EXPECT().
		InsertCustomer(ctx, name, email, stripeChargeDate, customerKey).
		Do(func() int64 {customerID++; return customerID}).
		Return(customerID, name, email, customerKey, stripeChargeDate)

	pgdb.EXPECT().
		GetCustomer(ctx, customerID).
		Return(customerID, RandomString(12), RandomString(12), RandomString(12),
			RandomString(12))

	return pgdb
}

func TestPaymentService_HealthCheck(t *testing.T) {
	ctx := context.Background()
	pgdb := setupDatastore()

	s := PaymentService{
		Cfg:          &Config{},
		PGDB:         pgdb,
		Logger:       logger,
		StripeClient: &client.API{},
	}
	req := paymentProto.HealthCheckRequest{}
	resp, err := s.HealthCheck(ctx, &req)
	assert.Nil(t, resp)
	assert.Errorf(t, err, "healthcheck failed; expected no error")
}
