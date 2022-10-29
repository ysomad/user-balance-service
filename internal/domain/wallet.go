package domain

import "errors"

var (
	ErrFundsNotAdded = errors.New("error occured during processing funds")
)

type Wallet struct {
	ID      string
	UserID  string
	Balance Amount
}
