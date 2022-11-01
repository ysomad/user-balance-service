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
	ErrSelfTransfer      = errors.New("cannot transfer funds from and to the same account")
)

type Account struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance Amount
}
