package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type accountRepo interface {
	DepositOrCreate(ctx context.Context, userID uuid.UUID, amount domain.Amount) (domain.Account, error)
	Deposit(ctx context.Context, userID uuid.UUID, amount domain.Amount) (domain.Account, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (domain.Account, error)
	Withdraw(ctx context.Context, userID uuid.UUID, a domain.Amount) (domain.Account, error)
}

type reservationRepo interface {
	Create(context.Context, dto.CreateReservationArgs) (*domain.Reservation, error)
	AddToRevenueReport(context.Context, dto.AddToRevenueReportArgs) (*domain.Reservation, error)
	Cancel(context.Context, dto.CancelReservationArgs) (*domain.Reservation, error)
}

type revenueReportRepo interface {
	GetOrCreate(ctx context.Context, userID uuid.UUID) (domain.RevenueReport, error)
}

type transactionRepo interface {
	Create(context.Context, dto.CreateTransactionArgs) (*domain.Transaction, error)
	CreateMultiple(context.Context, []dto.CreateTransactionArgs) ([]domain.Transaction, error)
	FindAllByUserID(context.Context, dto.FindTransactionListArgs) ([]domain.Transaction, error)
}
