package dto

import (
	"time"

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
	UserID         uuid.UUID
	LastID         uuid.UUID
	LastCommitedAt time.Time
	Sorts          domain.TransactionSorts
	PageSize       uint64
}
