package domain

import "github.com/google/uuid"

type AccountReserve struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Balance   Amount
}
