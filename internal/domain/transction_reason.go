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

func ReasonTransferTo(to uuid.UUID) Reason {
	return Reason(fmt.Sprintf("transfer to user %s", to.String()))
}

func ReasonTransferFrom(from uuid.UUID) Reason {
	return Reason(fmt.Sprintf("transfer from user %s", from.String()))
}

func ReasonReservationCancel(serviceID, orderID uuid.UUID) Reason {
	return Reason(fmt.Sprintf("reservation cancel service %s of order %s", serviceID.String(), orderID.String()))
}
