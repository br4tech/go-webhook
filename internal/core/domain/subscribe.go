package domain

type Subscribe struct {
	Id        int    `validate:"required"`
	PaymentId int    `validate:"required"`
	Webhook   string `validate:"required"`
	Token     string `validate:"required"`
}

func NewSubscribe(id, paymentId int, webhook, token string) *Subscribe {
	return &Subscribe{
		Id:        id,
		PaymentId: paymentId,
		Webhook:   webhook,
		Token:     token,
	}
}
