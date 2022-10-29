package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID     uuid.UUID
	UserID uuid.UUID

	// comment about the operation.
	Comment string

	// amount of funds withdrawn or deposited.
	Amount uint64

	Operation   operation
	CompletedAt time.Time
}

// DepositTransaction represents transaction for deposits.
type DepositTransaction struct {
	userID    uuid.UUID
	comment   string
	operation operation
	amount    Amount
}

func NewDepositTransaction(userID uuid.UUID, a Amount) (DepositTransaction, error) {
	return DepositTransaction{
		userID:    userID,
		comment:   "billing deposit",
		amount:    a,
		operation: operationDeposit,
	}, nil
}

func (t DepositTransaction) UserID() uuid.UUID    { return t.userID }
func (t DepositTransaction) Amount() Amount       { return t.amount }
func (t DepositTransaction) Comment() string      { return t.comment }
func (t DepositTransaction) Operation() operation { return t.operation }
