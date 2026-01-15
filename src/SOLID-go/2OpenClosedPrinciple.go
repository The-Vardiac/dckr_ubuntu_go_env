package main

/* open for extension closed for modification */

/* bad example */
func Pay(method string, amount float64) {
	if method == "paypal" {}
	if method == "stripe" {}
}


/* good example */
type PaymentGateway interface {
	Pay(amount float64) error
}

type paypal struct {
	paypal *paypal.Paypal
}
func NewPaypal(p *paypal.Paypal) PaymentGateway {
	return &paypal{paypal: p} 
}
func (p *paypal) Pay(amount float64) error {
	return nil
}

type stripe struct {
	stripe *stripe.Stripe
}
func NewStripe(s *stripe.Stripe) PaymentGateway {
	return &stripe{stripe: s}
}
func (p *stripe) Pay(amount float64) error {
	return nil
}