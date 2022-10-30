package dto

import "github.com/google/uuid"

type ReserveFundsArgs struct {
	UserID    uuid.UUID
	ServiceID uuid.UUID
	OrderID   uuid.UUID
	Amount    string
}
