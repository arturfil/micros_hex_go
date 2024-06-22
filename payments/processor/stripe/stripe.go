package stripe

import (
	"fmt"
	"log"

	common "github.com/arturfil/m_commons"
	pb "github.com/arturfil/m_commons/api"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

var gatewayHTTPAddr = common.EnvString("GATEWAY_HTTP_ADDRESS", "localhost:8080")

type Stripe struct {
}

func NewStripeProcessor() *Stripe {
	return &Stripe{}
}

func (s *Stripe) CreatePaymentLink(o *pb.Order) (string, error) {
	log.Printf("Creating Payment link for order %v", o)

	gatewaySuccessURL := fmt.Sprintf("%s/success.html", gatewayHTTPAddr)

	items := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range o.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID), // "price_1PU0koIAKhlfjLC8hQpbqfXk"
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		LineItems:  items,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gatewaySuccessURL),
	}

	result, err := session.New(params)
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
