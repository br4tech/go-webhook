package domain

import "time"

type PaymentStatusChangedEvent struct {
	PaymentId string
	NewStatus PaymentStatus
	Timestamp time.Time
}
