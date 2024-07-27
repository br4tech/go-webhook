package domain

type Subscription struct {
	Id           int
	PaymentId    int
	SubscriberId int
}

func NewSubscription(id, paymentId, subscriberId int) *Subscription {
	return &Subscription{
		Id:           id,
		PaymentId:    paymentId,
		SubscriberId: subscriberId,
	}
}
