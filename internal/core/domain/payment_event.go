package domain

import "time"

type PaymentStatusChangedEvent struct {
	PaymentID string
	NewStatus PaymentStatus
	Timestamp time.Time
}
