package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/pkg/pagetoken"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type account struct {
	log               internal.Logger
	accountRepo       accountRepo
	revenueReportRepo revenueReportRepo
	reservationRepo   reservationRepo
	transactionRepo   transactionRepo
}

func NewAccount(l internal.Logger, ar accountRepo, rrr revenueReportRepo, rr reservationRepo, tr transactionRepo) *account {
	return &account{
		log:               l,
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

	acc, err := a.accountRepo.DepositOrCreate(ctx, userID, depAmount)
	if err != nil {
		return domain.Account{}, err
	}

	_, err = a.transactionRepo.Create(ctx, dto.CreateTransactionArgs{
		UserID:    userID,
		Comment:   domain.ReasonBillingDeposit,
		Operation: domain.OperationDeposit,
		Amount:    depAmount,
	})
	if err != nil {
		return domain.Account{}, err
	}

	return acc, nil
}

// ReserveFunds withdraws funds from user account and adds it to reserve account.
func (a *account) ReserveFunds(ctx context.Context, args dto.ReserveFundsArgs) (*domain.Reservation, error) {
	resAmount, err := domain.NewAmount(args.Amount)
	if err != nil {
		return nil, err
	}

	if resAmount.IsZero() {
		return nil, domain.ErrZeroReserveAmount
	}

	acc, err := a.accountRepo.Withdraw(ctx, args.UserID, resAmount)
	if err != nil {
		return nil, err
	}

	res, err := a.reservationRepo.Create(ctx, dto.CreateReservationArgs{
		AccountID: acc.ID,
		ServiceID: args.ServiceID,
		OrderID:   args.OrderID,
		Amount:    resAmount,
	})
	if err != nil {
		return nil, err
	}

	_, err = a.transactionRepo.Create(ctx, dto.CreateTransactionArgs{
		UserID:    acc.UserID,
		Comment:   domain.ReasonReservationWithdraw,
		Operation: domain.OperationWithdraw,
		Amount:    resAmount,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
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
	var err error

	d := dto.FindTransactionListArgs{
		UserID:   args.UserID,
		Sorts:    args.Sorts,
		PageSize: args.PageSize,
	}

	if args.PageToken != "" {
		d.LastID, d.LastCommitedAt, err = pagetoken.Decode(args.PageToken)
		if err != nil {
			return domain.TransactionList{}, err
		}
	}

	if args.PageSize > domain.MaxPageSize || args.PageSize == 0 {
		d.PageSize = domain.DefaultPageSize
	}

	txs, err := a.transactionRepo.FindAllByUserID(ctx, d)
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

	if _, err = a.accountRepo.DepositOrCreate(ctx, args.ToUserID, transferAmount); err != nil {
		return domain.Account{}, err
	}

	_, err = a.transactionRepo.CreateMultiple(ctx, []dto.CreateTransactionArgs{
		{
			UserID:    args.FromUserID,
			Comment:   domain.ReasonTransferTo(args.ToUserID),
			Operation: domain.OperationWithdraw,
			Amount:    transferAmount,
		},
		{
			UserID:    args.ToUserID,
			Comment:   domain.ReasonTransferFrom(args.FromUserID),
			Operation: domain.OperationDeposit,
			Amount:    transferAmount,
		},
	})
	if err != nil {
		return domain.Account{}, err
	}

	return fromAcc, nil
}

func (a *account) CancelReservation(ctx context.Context, args dto.RawCancelReservationArgs) (*domain.Reservation, error) {
	amount, err := domain.NewAmount(args.Amount)
	if err != nil {
		return nil, err
	}

	res, err := a.reservationRepo.Cancel(ctx, dto.CancelReservationArgs{
		Amount:    amount,
		UserID:    args.UserID,
		ServiceID: args.ServiceID,
		OrderID:   args.OrderID,
	})
	if err != nil {
		return nil, err
	}

	if _, err = a.accountRepo.Deposit(ctx, args.UserID, amount); err != nil {
		return nil, err
	}

	_, err = a.transactionRepo.Create(ctx, dto.CreateTransactionArgs{
		UserID:    args.UserID,
		Comment:   domain.ReasonReservationCancel(res.ServiceID, res.OrderID),
		Operation: domain.OperationDeposit,
		Amount:    amount,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
