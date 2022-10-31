package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type account struct {
	accountRepo       accountRepo
	revenueReportRepo revenueReportRepo
	reservationRepo   reservationRepo
}

func NewAccount(ar accountRepo, rrr revenueReportRepo, rr reservationRepo) *account {
	return &account{
		accountRepo:       ar,
		revenueReportRepo: rrr,
		reservationRepo:   rr,
	}
}

func (a *account) DepositFunds(ctx context.Context, userID uuid.UUID, amount string) (domain.Account, error) {
	depAmount, err := domain.NewAmount(amount)
	if err != nil {
		return domain.Account{}, err
	}

	if depAmount.IsZero() {
		return domain.Account{}, domain.ErrZeroDeposit
	}

	t, err := domain.NewDepositTransaction(userID, depAmount)
	if err != nil {
		return domain.Account{}, err
	}

	acc, err := a.accountRepo.UpdateOrCreate(ctx, t)
	if err != nil {
		return domain.Account{}, err
	}

	return acc, nil
}

// ReserveFunds withdraws funds from user account and adds it to reserve account.
func (a *account) ReserveFunds(ctx context.Context, args dto.ReserveFundsArgs) (*dto.AccountWithReservation, error) {
	amount, err := domain.NewAmount(args.Amount)
	if err != nil {
		return nil, err
	}

	if amount.IsZero() {
		return nil, domain.ErrZeroReserveAmount
	}

	acc, err := a.accountRepo.Withdraw(ctx, args.UserID, amount)
	if err != nil {
		return nil, err
	}

	res, err := a.reservationRepo.Create(ctx, dto.CreateReservationArgs{
		AccountID: acc.ID,
		ServiceID: args.ServiceID,
		OrderID:   args.OrderID,
		Amount:    amount,
	})
	if err != nil {
		return nil, err
	}

	return &dto.AccountWithReservation{
		Account:     acc,
		Reservation: res,
	}, nil
}

func (a *account) DeclareRevenue(ctx context.Context, args dto.DeclareRevenueArgs) (*domain.Reservation, error) {
	amount, err := domain.NewAmount(args.Amount)
	if err != nil {
		return nil, err
	}

	report, err := a.revenueReportRepo.GetOrCreate(ctx, args.UserID)
	if err != nil {
		return nil, err
	}

	res, err := a.reservationRepo.AddToRevenueReport(ctx, dto.AddToRevenueReportArgs{
		UserID:          args.UserID,
		ServiceID:       args.ServiceID,
		OrderID:         args.OrderID,
		Amount:          amount,
		RevenueReportID: report.ID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *account) GetByUserID(ctx context.Context, userID uuid.UUID) (domain.Account, error) {
	panic("implement me")
	return a.accountRepo.FindByUserID(ctx, userID)
}
