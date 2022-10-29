package service

import (
	"context"

	"github.com/ysomad/avito-internship-task/internal/domain"
)

type wallet struct{}

func NewWallet() *wallet {
	return &wallet{}
}

func (w *wallet) AddFunds(ctx context.Context, userID, amount string) (domain.Wallet, error) {

	return domain.Wallet{}, nil
}
