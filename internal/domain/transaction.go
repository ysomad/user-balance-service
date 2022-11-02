package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/ysomad/avito-internship-task/internal/pkg/pagetoken"
	"github.com/ysomad/avito-internship-task/internal/pkg/sort"
)

const (
	DefaultPageSize = 10
	MaxPageSize     = 500
)

type Transaction struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Comment    Reason
	Amount     Amount
	Operation  Op
	CommitedAt time.Time
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
		log.Debug().Msg(txs[txsLen-1].ID.String())
		list.NextPageToken = pagetoken.Encode(txs[txsLen-1].ID, txs[txsLen-1].CommitedAt)
		list.Transactions = txs[:txsLen-1]
	}

	return list, nil
}
