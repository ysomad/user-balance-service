package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

func (h *handler) ReserveFunds(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	var (
		req ReserveFundsRequest
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody, nil)
		return
	}

	var a domain.AccountAggregate

	err = h.tx.RunAtomic(r.Context(), func(txCtx context.Context) error {
		a, err = h.account.ReserveFunds(txCtx, dto.ReserveFundsArgs{
			UserID:    userID,
			ServiceID: req.ServiceID,
			OrderID:   req.OrderID,
			Amount:    req.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		h.log.Error(err.Error())

		switch {
		case errors.Is(err, domain.ErrNotEnoughFunds):
			writeError(w, http.StatusConflict, err, nil)
			return
		case errors.Is(err, domain.ErrZeroReserveAmount):
			writeError(w, http.StatusBadRequest, err, nil)
			return
		}

		writeError(w, http.StatusInternalServerError, err, nil)
		return
	}

	writeOK(w, AccountAggregate{
		ID:       a.ID,
		UserID:   a.UserID,
		Balance:  a.Balance.String(),
		Reserved: a.Reserved.String(),
	})
}
