package service

import (
	"context"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type wallet struct {
	tx              transactor
	walletRepo      walletRepository
	transactionRepo transactionRepository
}

func NewWallet(t transactor, wr walletRepository, tr transactionRepository) *wallet {
	return &wallet{
		tx:              t,
		walletRepo:      wr,
		transactionRepo: tr,
	}
}

func (w *wallet) AddFunds(ctx context.Context, userID, amount string) (domain.Wallet, error) {
	a, err := domain.NewAmount(amount)
	if err != nil {
		return domain.Wallet{}, err
	}

	t, err := domain.NewDepositTransaction(userID, a)
	if err != nil {
		return domain.Wallet{}, err
	}

	var wt domain.Wallet

	err = w.tx.WithinTransaction(ctx, func(txCtx context.Context) error {
		wt, err = w.walletRepo.UpsertWalletBalance(txCtx, t)
		if err != nil {
			return err
		}

		_, err = w.transactionRepo.CreateTransaction(txCtx, dto.CreateTransactionArgs{
			UserID:     t.UserID(),
			Comment:    t.Comment(),
			FromUserID: nil,
			Operation:  t.Operation().String(),
			Amount:     t.Amount().UInt64(),
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return domain.Wallet{}, err
	}

	return wt, nil
}
