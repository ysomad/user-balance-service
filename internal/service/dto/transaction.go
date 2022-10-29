package dto

type CreateTransactionArgs struct {
	UserID     string
	Comment    string
	FromUserID *string
	Operation  string
	Amount     uint64
}
