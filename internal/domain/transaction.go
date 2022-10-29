package domain

import (
	"time"
)

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

// DepositTransaction represents transaction for deposits.
type DepositTransaction struct {
	UserID    string
	Comment   string
	Operation operation

	amount amount
}

func NewDepositTransaction(userID string, a amount) (DepositTransaction, error) {
	return DepositTransaction{
		UserID:    userID,
		Comment:   "billing deposit",
		amount:    a,
		Operation: operationDeposit,
	}, nil
}
