package domain

import (
	"errors"
	"strings"
)

type Op string

const (
	OpDeposit  = "DEPOSIT"
	OpWithdraw = "WITHDRAW"
)

func NewOp(s string) (Op, error) {
	s = strings.ToUpper(s)
	switch s {
	case OpDeposit, OpWithdraw:
		return Op(s), nil
	}

	return "", errors.New("not supported transaction operation")
}

func (o Op) String() string { return string(o) }
