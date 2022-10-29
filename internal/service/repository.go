package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type accountRepository interface {
	UpsertAccountBalance(context.Context, domain.DepositTransaction) (domain.Account, error)
	FindUserAccount(ctx context.Context, userID uuid.UUID) (domain.AccountAggregate, error)
	WithdrawFromUserAccount(ctx context.Context, userID uuid.UUID, a domain.Amount) (domain.Account, error)
}

type transactionRepository interface {
	Create(context.Context, dto.CreateTransactionArgs) (*domain.Transaction, error)
}
