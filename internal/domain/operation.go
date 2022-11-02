package domain

type Operation uint8

const (
	OperationUndefined = iota
	OperationDeposit
	OperationWithdraw
)

func (o Operation) String() string {
	switch o {
	case OperationDeposit:
		return "DEPOSIT"
	case OperationWithdraw:
		return "WITHDRAW"
	}
	return "UNDEFINED"
}
