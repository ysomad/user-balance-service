package dto

import (
	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
)

type CreateReservationArgs struct {
	AccountID uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Amount    domain.Amount
}

type AddToRevenueReportArgs struct {
	UserID          uuid.UUID
	ServiceID       uuid.UUID
	OrderID         uuid.UUID
	Amount          domain.Amount
	RevenueReportID uuid.UUID
}
