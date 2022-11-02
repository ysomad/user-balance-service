package dto

import (
	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
)

type CreateTransactionArgs struct {
	UserID    uuid.UUID
	Comment   domain.Reason
	Operation domain.Op
	Amount    domain.Amount
}

type FindTransactionListArgs struct {
	UserID   uuid.UUID
	PrevID   uint32
	Sorts    domain.TransactionSorts
	PageSize uint64
}
