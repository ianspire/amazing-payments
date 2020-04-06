package pkg

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

func NewStripeClient(cfg *Config) (*client.API, error) {

	stripeClient := &client.API{}
	stripeClient.Init(cfg.StripeKey, nil)

	// test the client to be sure it's working
	_, err := stripeClient.Balance.Get(&stripe.BalanceParams{})
	if err != nil {
		return nil, err
	}

	return stripeClient, nil
}
