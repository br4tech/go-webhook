package domain

type Subscribe struct {
	Id           int
	PaymentId    int
	SubscriberId int
}

func NewSubscribe(id, paymentId, subscriberId int) *Subscribe {
	return &Subscribe{
		Id:           id,
		PaymentId:    paymentId,
		SubscriberId: subscriberId,
	}
}
