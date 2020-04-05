package pkg

import (
	"context"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"github.com/stripe/stripe-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *PaymentService) HealthCheck(ctx context.Context, req *paymentProto.HealthCheckRequest) (
	*paymentProto.HealthCheckResponse, error) {
	var healthy bool
	err := p.PGDB.HealthCheck()
	if err != nil {
		healthy = false
	} else {
		healthy = true
	}
	return &paymentProto.HealthCheckResponse{
		IsHealthy:            healthy,
	}, err
}

func (p *PaymentService) CreateCustomer(ctx context.Context, req *paymentProto.CreateCustomerRequest) (
	*paymentProto.Customer, error) {

	if req.Email == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "must provide email and name")
	}

	stripeCustomer, err := p.StripeClient.Customers.New(&stripe.CustomerParams{
		Email:               &req.Email,
		Name:                &req.Name,
	})
	if err != nil {
		p.Logger.Errorw("error sending stripe client new customer request",
			"email", &req.Email,
			"name", &req.Email,
		)
		return nil, status.Error(codes.Internal, "stripe client call failed")
	}

	cust, err := p.PGDB.InsertCustomer(ctx, req.Name, req.Email, req.StripeChargeDate, stripeCustomer.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to store customer")
	}

	return &paymentProto.Customer{
		CustomerID:        cust.CustomerID,
		Name:              cust.Name,
		Email:             cust.Email,
		StripeCustomerKey: cust.StripeCustomerKey,
		StripeChargeDate:  cust.StripeChargeDate,
	}, err
}

func (p *PaymentService) GetCustomer(ctx context.Context, req *paymentProto.GetCustomerRequest) (
	*paymentProto.Customer, error) {

	if req.CustomerID == 0 {
		return nil, status.Error(codes.InvalidArgument, "must provide a valid customer ID")
	}

	cust, err := p.PGDB.GetCustomer(ctx, req.CustomerID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve customer from DB")
	}

	return &paymentProto.Customer{
		CustomerID:           cust.CustomerID,
		Name:                 cust.Name,
		Email:                cust.Email,
		StripeCustomerKey:    cust.StripeCustomerKey,
		StripeChargeDate:     cust.StripeChargeDate,
	}, err
}

func (p *PaymentService) UpdateCustomer(ctx context.Context, req *paymentProto.UpdateCustomerRequest) (
	*paymentProto.Customer, error) {

	if req.Email == "" && req.Name == "" && req.StripeChargeDate == "" {
		return nil, status.Error(codes.InvalidArgument, "must provide a field to update - email, name," +
			"stripe_charge_date")
	}

	cust, err := p.PGDB.UpdateCustomer(ctx, &req.Name, &req.Email, &req.StripeChargeDate)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update customer in DB")
	}

	return &paymentProto.Customer{
		CustomerID:           cust.CustomerID,
		Name:                 cust.Name,
		Email:                cust.Email,
		StripeCustomerKey:    cust.StripeCustomerKey,
		StripeChargeDate:     cust.StripeChargeDate,
	}, err
}
