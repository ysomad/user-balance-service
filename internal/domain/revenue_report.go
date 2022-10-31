package domain

import (
	"time"

	"github.com/google/uuid"
)

type RevenueReport struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	CreatedAt time.Time
}
