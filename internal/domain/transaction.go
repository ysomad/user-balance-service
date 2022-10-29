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
	FromUserID *string

	// amount of funds withdrawn or deposited.
	Amount uint64

	Operation   operation
	CompletedAt time.Time
}

// DepositTransaction represents transaction for deposits.
type DepositTransaction struct {
	userID    string
	comment   string
	operation operation
	amount    Amount
}

func NewDepositTransaction(userID string, a Amount) (DepositTransaction, error) {
	return DepositTransaction{
		userID:    userID,
		comment:   "billing deposit",
		amount:    a,
		operation: operationDeposit,
	}, nil
}

func (t DepositTransaction) UserID() string       { return t.userID }
func (t DepositTransaction) Amount() Amount       { return t.amount }
func (t DepositTransaction) Comment() string      { return t.comment }
func (t DepositTransaction) Operation() operation { return t.operation }
