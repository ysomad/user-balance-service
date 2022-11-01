package domain

type Reason string

const (
	ReasonReservationWithdraw = "service reservation withdraw"
	ReasonBillingDeposit      = "deposit from billing"
)
