package domain

import (
	"time"

	"github.com/ysomad/avito-internship-task/internal/pkg/pagetoken"
	"github.com/ysomad/avito-internship-task/internal/pkg/sort"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID `json:"id"`
	AccountID  uuid.UUID `json:"account_id"`
	Comment    Reason    `json:"comment"`
	Amount     Amount    `json:"amount"`
	Operation  Op        `json:"operation"`
	CommitedAt time.Time `json:"commited_at,omitempty"`
}

type TransactionSorts struct {
	Amount     sort.Order
	CommitedAt sort.Order
}

type TransactionList struct {
	Transactions  []Transaction
	NextPageToken string
}

func NewTransactionList(txs []Transaction, pageSize uint64) (TransactionList, error) {
	list := TransactionList{
		Transactions: txs,
	}

	txsLen := uint64(len(txs))
	if txsLen == pageSize+1 {
		list.NextPageToken = pagetoken.Encode(txs[txsLen-1].ID, txs[txsLen-1].CommitedAt)
		list.Transactions = txs[:txsLen-1]
	}

	return list, nil
}
