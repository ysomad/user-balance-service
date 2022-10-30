package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type accountRepo interface {
	Upsert(context.Context, domain.DepositTransaction) (domain.Account, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (domain.AccountAggregate, error)
	Withdraw(ctx context.Context, userID uuid.UUID, a domain.Amount) (domain.Account, error)
}

type reserveAccountRepo interface {
	Upsert(ctx context.Context, accountID uuid.UUID, a domain.Amount) (domain.AccountReserve, error)
}

type transactionRepo interface {
	Create(context.Context, dto.CreateTransactionArgs) (*domain.Transaction, error)
}
