package dto

import (
	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
)

type ReserveFundsArgs struct {
	UserID    uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Amount    string
}

type DeclareRevenueArgs struct {
	UserID    uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Amount    string
}

type AccountWithReservation struct {
	Account     domain.Account
	Reservation *domain.Reservation
}

type GetTransactionListArgs struct {
	UserID    uuid.UUID
	PageToken string
	PageSize  uint64
	Sorts     domain.TransactionSorts
}
