package domain

import "time"

type Transaction struct {
	ID     string
	UserID string

	// comment about the operation.
	Comment string

	// from where funds was withdrawn or deposited
	// if empty that means funds came from billing.
	FromUserID string

	// amount of funds withdrawn or deposited.
	Amount uint64

	Operation   operation
	CompletedAt time.Time
}
