package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrFundsNotCredited  = errors.New("error occured during processing funds")
	ErrAccountNotFound   = errors.New("user account not found")
	ErrZeroDeposit       = errors.New("amount of funds to deposit must be greater than 0")
	ErrZeroReserveAmount = errors.New("amount to reserve must be greated than 0")
	ErrNotEnoughFunds    = errors.New("not enough funds on account balance")
)

type Account struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance Amount
}
