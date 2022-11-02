package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/ysomad/avito-internship-task/internal/pkg/sort"
)

const (
	DefaultPageSize = 10
	MaxPageSize     = 500
)

type Transaction struct {
	ID         uint32
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
	NextPageToken uint32
}

func NewTransactionList(txs []Transaction, pageSize uint64) (TransactionList, error) {
	list := TransactionList{
		Transactions: txs,
	}

	txsLen := uint64(len(txs))
	if txsLen == pageSize+1 {
		log.Debug().Uint32("lastID", txs[txsLen-1].ID)
		list.NextPageToken = txs[txsLen-1].ID
		list.Transactions = txs[:txsLen-1]
	}

	return list, nil
}
