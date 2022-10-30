package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type account struct {
	accountRepo        accountRepo
	reserveAccountRepo reserveAccountRepo
	transactionRepo    transactionRepo
}

func NewAccount(ar accountRepo, arr reserveAccountRepo, tr transactionRepo) *account {
	return &account{
		accountRepo:        ar,
		reserveAccountRepo: arr,
		transactionRepo:    tr,
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

	acc, err := a.accountRepo.Upsert(ctx, t)
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

// ReserveFunds withdraws funds from user account and adds it to reserve account.
func (a *account) ReserveFunds(ctx context.Context, args dto.ReserveFundsArgs) (domain.AccountAggregate, error) {
	amount, err := domain.NewAmount(args.Amount)
	if err != nil {
		return domain.AccountAggregate{}, err
	}

	if amount.IsZero() {
		return domain.AccountAggregate{}, domain.ErrZeroReserveAmount
	}

	acc, err := a.accountRepo.Withdraw(ctx, args.UserID, amount)
	if err != nil {
		return domain.AccountAggregate{}, err
	}

	accReserve, err := a.reserveAccountRepo.Upsert(ctx, acc.ID, amount)
	if err != nil {
		return domain.AccountAggregate{}, err
	}

	return domain.AccountAggregate{
		ID:       acc.ID,
		UserID:   acc.UserID,
		Balance:  acc.Balance,
		Reserved: accReserve.Balance,
	}, nil
}

func (a *account) GetByUserID(ctx context.Context, userID uuid.UUID) (domain.AccountAggregate, error) {
	panic("implement me")
	return a.accountRepo.FindByUserID(ctx, userID)
}
