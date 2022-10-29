package service

import (
	"context"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type walletRepository interface {
	UpsertWalletBalance(context.Context, domain.DepositTransaction) (domain.Wallet, error)
}

type transactionRepository interface {
	CreateTransaction(context.Context, dto.CreateTransactionArgs) (*domain.Transaction, error)
}
