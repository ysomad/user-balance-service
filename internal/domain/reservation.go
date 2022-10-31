package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrReservationNotDeclared = errors.New("service and order with amount of funds to declare not found")
)

type Reservation struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	ServiceID       uuid.UUID
	OrderID         uuid.UUID
	Amount          Amount
	Declared        bool
	CreatedAt       time.Time
	DeclaredAt      *time.Time
	RevenueReportID uuid.UUID
}
