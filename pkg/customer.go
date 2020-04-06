package pkg

import (
	"context"
	paymentProto "github.com/ianspire/amazing-payments/proto"
	"github.com/stripe/stripe-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ps *PaymentService) HealthCheck(ctx context.Context, req *paymentProto.HealthCheckRequest) (
	*paymentProto.HealthCheckResponse, error) {
	ps.Logger.Infow("healthcheck request",
		"context", &ctx,
		"request", &req,
	)
	var healthy bool
	err := ps.PGDB.HealthCheck()
	if err != nil {
		healthy = false
	} else {
		healthy = true
	}
	return &paymentProto.HealthCheckResponse{
		IsHealthy: healthy,
	}, err
}

/*
CreateCustomer method handles requests to /v1/customer, calls stripe to register a new customer by email address,
and returns the complete record, including Stripe customer key.
*/
func (ps *PaymentService) CreateCustomer(ctx context.Context, req *paymentProto.CreateCustomerRequest) (
	*paymentProto.Customer, error) {

	ps.Logger.Infow("create customer request",
		"context", &ctx,
		"request", &req,
	)

	if req.Email == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "must provide email and name")
	}

	stripeCustomer, err := ps.StripeClient.Customers.New(&stripe.CustomerParams{
		Email: &req.Email,
		Name:  &req.Name,
	})
	if err != nil {
		ps.Logger.Errorw("error sending stripe client new customer request",
			"email", &req.Email,
			"name", &req.Email,
		)
		return nil, status.Error(codes.Internal, "stripe client call failed")
	}

	cust, err := ps.PGDB.InsertCustomer(ctx, req.Name, req.Email, req.StripeChargeDate, stripeCustomer.ID)
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

/*
GetCustomer method handles requests to /v1/customer/{customerID}, obtains our internal record by customerID,
and returns all customer record data including stripe_customer_key, which can be used to make payment requests.
*/
func (ps *PaymentService) GetCustomer(ctx context.Context, req *paymentProto.GetCustomerRequest) (
	*paymentProto.Customer, error) {

	ps.Logger.Infow("get customer request",
		"context", &ctx,
		"request", &req,
	)

	if req.CustomerID == 0 {
		return nil, status.Error(codes.InvalidArgument, "must provide a valid customer ID")
	}

	cust, err := ps.PGDB.GetCustomer(ctx, req.CustomerID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve customer from DB")
	}

	return &paymentProto.Customer{
		CustomerID:        cust.CustomerID,
		Name:              cust.Name,
		Email:             cust.Email,
		StripeCustomerKey: cust.StripeCustomerKey,
		StripeChargeDate:  cust.StripeChargeDate,
	}, err
}

/*
Consideration should be taken with updating our local customer information, and how it should be updated
with Stripe.  Because we're under a time crunch, I'm going to eliminate this method for now.
*/
//func (ps *PaymentService) UpdateCustomer(ctx context.Context, req *paymentProto.UpdateCustomerRequest) (
//	*paymentProto.Customer, error) {
//
//	if req.Email == "" && req.Name == "" && req.StripeChargeDate == "" {
//		return nil, status.Error(codes.InvalidArgument, "must provide a field to update - email, name," +
//			"stripe_charge_date")
//	}
//
//	cust, err := ps.pgdb.UpdateCustomer(ctx, &req.Name, &req.Email, &req.StripeChargeDate)
//	if err != nil {
//		return nil, status.Error(codes.Internal, "failed to update customer in DB")
//	}
//
//	return &paymentProto.Customer{
//		CustomerID:           cust.CustomerID,
//		Name:                 cust.Name,
//		Email:                cust.Email,
//		StripeCustomerKey:    cust.StripeCustomerKey,
//		StripeChargeDate:     cust.StripeChargeDate,
//	}, err
//}
