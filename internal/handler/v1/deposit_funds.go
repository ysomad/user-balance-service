package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
)

func (h *handler) DepositFunds(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	var (
		req DepositFundsRequest
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody, nil)
		return
	}

	var a domain.Account

	err = h.tx.RunAtomic(r.Context(), func(txCtx context.Context) error {
		a, err = h.account.DepositFunds(txCtx, userID, req.Amount)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		h.log.Error(err.Error())

		switch {
		case errors.Is(err, domain.ErrInvalidMajorAmount):
			writeError(w, http.StatusBadRequest, err, nil)
			return
		case errors.Is(err, domain.ErrZeroDeposit):
			writeError(w, http.StatusBadRequest, err, nil)
			return
		}

		writeError(w, http.StatusInternalServerError, domain.ErrFundsNotCredited, nil)
		return
	}

	writeOK(w, Account{
		ID:      a.ID,
		UserID:  a.UserID,
		Balance: a.Balance.String(),
	})
}
