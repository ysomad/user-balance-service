package domain

import (
	"errors"
	"strings"
)

type operation string

const (
	operationDeposit  = "DEPOSIT"
	operationWithdraw = "WITHDRAW"
)

func NewOperation(s string) (operation, error) {
	s = strings.ToUpper(s)
	switch s {
	case operationDeposit, operationWithdraw:
		return operation(s), nil
	}

	return "", errors.New("not support transaction operation")
}

func (op operation) String() string { return string(op) }
