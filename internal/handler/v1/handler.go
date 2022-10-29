package v1

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

var _ ServerInterface = &handler{}

type validate interface {
	Into(r io.Reader, dest any) error
	Translate(error) map[string]string
}

type atomicRunner interface {
	RunAtomic(context.Context, func(context.Context) error) error
}

type accountService interface {
	DepositFunds(ctx context.Context, userID uuid.UUID, amount string) (domain.Account, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (domain.AccountAggregate, error)
	ReserveFunds(context.Context, dto.ReserveFundsArgs) (domain.AccountAggregate, error)
}

type handler struct {
	log      internal.Logger
	validate validate
	tx       atomicRunner
	account  accountService
}

func NewHandler(l internal.Logger, v validate, ar atomicRunner, as accountService) *handler {
	return &handler{
		log:      l,
		validate: v,
		tx:       ar,
		account:  as,
	}
}
