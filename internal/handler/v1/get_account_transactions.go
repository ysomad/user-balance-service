package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/pkg/pagetoken"
	"github.com/ysomad/avito-internship-task/internal/pkg/sort"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

func (h *handler) GetAccountTransactions(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	var req TransactionListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody)
		return
	}

	l, err := h.account.GetTransactionList(r.Context(), dto.GetTransactionListArgs{
		UserID:    userID,
		PageToken: req.PageToken,
		PageSize:  req.PageSize,
		Sorts: domain.TransactionSorts{
			Amount:     sort.NewOrder(string(req.Sort.Amount)),
			CommitedAt: sort.NewOrder(string(req.Sort.CommitTime)),
		},
	})
	if err != nil {
		h.log.Error(err.Error())

		if errors.Is(err, pagetoken.ErrInvalid) {
			writeError(w, http.StatusBadRequest, err)
			return
		}

		writeError(w, http.StatusInternalServerError, err)
		return
	}

	txs := make([]Transaction, len(l.Transactions))
	for i, d := range l.Transactions {
		txs[i] = Transaction{
			ID:         d.ID,
			AccountID:  d.AccountID,
			Amount:     d.Amount.String(),
			Comment:    string(d.Comment),
			Operation:  TransactionOperation(d.Operation.String()),
			CommitedAt: d.CommitedAt,
		}
	}

	writeOK(w, TransactionList{
		Transactions:  txs,
		NextPageToken: l.NextPageToken,
	})
}
