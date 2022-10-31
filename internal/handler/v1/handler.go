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

type atomicWrapper interface {
	Wrap(context.Context, func(context.Context) error) error
}

type accountService interface {
	DepositFunds(ctx context.Context, userID uuid.UUID, amount string) (domain.Account, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (domain.Account, error)
	ReserveFunds(context.Context, dto.ReserveFundsArgs) (*dto.AccountWithReservation, error)
	DeclareRevenue(context.Context, dto.DeclareRevenueArgs) (*domain.Reservation, error)
}

type revenueReportService interface {
}

type handler struct {
	log      internal.Logger
	validate validate
	atomic   atomicWrapper
	account  accountService
}

func NewHandler(l internal.Logger, v validate, ar atomicWrapper, as accountService) *handler {
	return &handler{
		log:      l,
		validate: v,
		atomic:   ar,
		account:  as,
	}
}
