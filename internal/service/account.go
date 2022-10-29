package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type account struct {
	accountRepo     accountRepository
	transactionRepo transactionRepository
}

func NewAccount(ar accountRepository, tr transactionRepository) *account {
	return &account{
		accountRepo:     ar,
		transactionRepo: tr,
	}
}

func (a *account) DepositFunds(ctx context.Context, userID uuid.UUID, amount string) (domain.Account, error) {
	depAmount, err := domain.NewAmount(amount)
	if err != nil {
		return domain.Account{}, err
	}

	if depAmount.IsZero() {
		return domain.Account{}, domain.ErrZeroDeposit
	}

	t, err := domain.NewDepositTransaction(userID, depAmount)
	if err != nil {
		return domain.Account{}, err
	}

	acc, err := a.accountRepo.UpsertAccountBalance(ctx, t)
	if err != nil {
		return domain.Account{}, err
	}

	// _, err = a.transactionRepo.CreateTransaction(txCtx, dto.CreateTransactionArgs{
	// 	UserID:    t.UserID(),
	// 	Comment:   t.Comment(),
	// 	Operation: t.Operation().String(),
	// 	Amount:    t.Amount().UInt64(),
	// })
	// if err != nil {
	// 	return err
	// }

	return acc, nil
}

func (a *account) GetByUserID(ctx context.Context, userID uuid.UUID) (domain.AccountAggregate, error) {
	panic("implement me")
	return a.accountRepo.FindUserAccount(ctx, userID)
}

func (a *account) ReserveFunds(ctx context.Context, args dto.ReserveFundsArgs) (domain.AccountAggregate, error) {
	panic("implement me")
	return domain.AccountAggregate{}, nil
}
