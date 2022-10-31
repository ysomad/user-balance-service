package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type accountRepo interface {
	UpdateOrCreate(context.Context, domain.DepositTransaction) (domain.Account, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (domain.AccountAggregate, error)
	Withdraw(ctx context.Context, userID uuid.UUID, a domain.Amount) (domain.Account, error)
}

type reservationRepo interface {
	Create(context.Context, dto.CreateReservationArgs) (*domain.Reservation, error)
	AddToRevenueReport(context.Context, dto.AddToRevenueReportArgs) (*domain.Reservation, error)
}

type revenueReportRepo interface {
	GetOrCreate(ctx context.Context, userID uuid.UUID) (domain.RevenueReport, error)
}
