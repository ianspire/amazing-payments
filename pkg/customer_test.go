package pkg

import (
	"context"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentService_HealthCheck(t *testing.T) {
	var s PaymentService
	ctx := context.Background()
	req := paymentProto.HealthCheckRequest{}
	resp, err := s.HealthCheck(ctx, &req)
	assert.Nil(t, resp)
	assert.Errorf(t, err, "healthcheck failed; expected no error")
}
