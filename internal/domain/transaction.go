package domain

import "time"

type operation uint8

const (
	operationDeposit = iota
	operationWithdraw
)

type Transaction struct {
	ID     string
	UserID string

	// comment about the operation.
	Comment string

	// from where funds was withdrawn or deposited
	// if empty that means funds came from billing.
	From string

	// amount of funds withdrawn or deposited.
	Amount uint64

	Operation   operation
	CompletedAt time.Time
}
