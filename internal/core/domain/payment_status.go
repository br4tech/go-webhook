package domain

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "PENDENTE"
	PaymentStatusApproved PaymentStatus = "APROVADO"
	PaymentStatusRejected PaymentStatus = "REJEITADO"
)
