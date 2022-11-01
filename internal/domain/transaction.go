package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Comment    Reason
	Amount     Amount
	Operation  Op
	CommitedAt time.Time
}
