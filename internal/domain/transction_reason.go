package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Reason string

const (
	ReasonReservationWithdraw Reason = "service reservation withdraw"
	ReasonBillingDeposit      Reason = "deposit from billing"
)

func ReasonTransfer(to uuid.UUID) Reason {
	return Reason(fmt.Sprintf("transfer to user with id: %s", to.String()))
}

func ReasonTransferFrom(from uuid.UUID) Reason {
	return Reason(fmt.Sprintf("transfer from user with id: %s", from.String()))
}
