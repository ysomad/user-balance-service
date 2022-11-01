package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/pkg/pagetoken"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type account struct {
	accountRepo       accountRepo
	revenueReportRepo revenueReportRepo
	reservationRepo   reservationRepo
	transactionRepo   transactionRepo
}

func NewAccount(ar accountRepo, rrr revenueReportRepo, rr reservationRepo, tr transactionRepo) *account {
	return &account{
		accountRepo:       ar,
		revenueReportRepo: rrr,
		reservationRepo:   rr,
		transactionRepo:   tr,
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

	acc, err := a.accountRepo.UpdateOrCreate(ctx, userID, depAmount)
	if err != nil {
		return domain.Account{}, err
	}

	_, err = a.transactionRepo.Create(ctx, dto.CreateTransactionArgs{
		UserID:    userID,
		Comment:   domain.ReasonBillingDeposit,
		Operation: domain.OpDeposit,
		Amount:    depAmount,
	})
	if err != nil {
		return domain.Account{}, err
	}

	return acc, nil
}

// ReserveFunds withdraws funds from user account and adds it to reserve account.
func (a *account) ReserveFunds(ctx context.Context, args dto.ReserveFundsArgs) (dto.AccountWithReservation, error) {
	resAmount, err := domain.NewAmount(args.Amount)
	if err != nil {
		return dto.AccountWithReservation{}, err
	}

	if resAmount.IsZero() {
		return dto.AccountWithReservation{}, domain.ErrZeroReserveAmount
	}

	acc, err := a.accountRepo.Withdraw(ctx, args.UserID, resAmount)
	if err != nil {
		return dto.AccountWithReservation{}, err
	}

	res, err := a.reservationRepo.Create(ctx, dto.CreateReservationArgs{
		AccountID: acc.ID,
		ServiceID: args.ServiceID,
		OrderID:   args.OrderID,
		Amount:    resAmount,
	})
	if err != nil {
		return dto.AccountWithReservation{}, err
	}

	_, err = a.transactionRepo.Create(ctx, dto.CreateTransactionArgs{
		UserID:    acc.UserID,
		Comment:   domain.ReasonReservationWithdraw,
		Operation: domain.OpWithdraw,
		Amount:    resAmount,
	})
	if err != nil {
		return dto.AccountWithReservation{}, err
	}

	return dto.AccountWithReservation{
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
	return a.accountRepo.FindByUserID(ctx, userID)
}

func (a *account) GetTransactionList(ctx context.Context, args dto.GetTransactionListArgs) (domain.TransactionList, error) {
	var (
		lastID         uuid.UUID
		lastCommitedAt time.Time
		err            error
	)

	if args.PageToken != "" {
		lastID, lastCommitedAt, err = pagetoken.Decode(args.PageToken)
		if err != nil {
			return domain.TransactionList{}, err
		}
	}

	if args.PageSize > domain.MaxPageSize || args.PageSize == 0 {
		args.PageSize = domain.DefaultPageSize
	}

	txs, err := a.transactionRepo.FindAllByUserID(ctx, dto.FindTransactionListArgs{
		UserID:         args.UserID,
		LastID:         lastID,
		LastCommitedAt: lastCommitedAt,
		Sorts:          args.Sorts,
		PageSize:       args.PageSize,
	})
	if err != nil {
		return domain.TransactionList{}, err
	}

	return domain.NewTransactionList(txs, args.PageSize)
}

func (a *account) TransferFunds(ctx context.Context, args dto.TransferFundsArgs) (domain.Account, error) {
	if args.FromUserID == args.ToUserID {
		return domain.Account{}, domain.ErrSelfTransfer
	}

	transferAmount, err := domain.NewAmount(args.Amount)
	if err != nil {
		return domain.Account{}, err
	}

	if transferAmount.IsZero() {
		return domain.Account{}, domain.ErrZeroAmount
	}

	fromAcc, err := a.accountRepo.Withdraw(ctx, args.FromUserID, transferAmount)
	if err != nil {
		return domain.Account{}, err
	}

	if _, err = a.accountRepo.UpdateOrCreate(ctx, args.ToUserID, transferAmount); err != nil {
		return domain.Account{}, err
	}

	_, err = a.transactionRepo.CreateMultiple(ctx, []dto.CreateTransactionArgs{
		{
			UserID:    args.FromUserID,
			Comment:   domain.ReasonTransfer(args.ToUserID),
			Operation: domain.OpWithdraw,
			Amount:    transferAmount,
		},
		{
			UserID:    args.ToUserID,
			Comment:   domain.ReasonTransferFrom(args.FromUserID),
			Operation: domain.OpDeposit,
			Amount:    transferAmount,
		},
	})
	if err != nil {
		return domain.Account{}, err
	}

	return fromAcc, nil
}
