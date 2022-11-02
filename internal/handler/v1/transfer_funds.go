package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	google_uuid "github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

func (h *handler) TransferFunds(w http.ResponseWriter, r *http.Request, userID google_uuid.UUID) {
	var (
		req TransferFundsRequest
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody)
		return
	}

	var a domain.Account

	err = h.atomic.Run(r.Context(), func(txCtx context.Context) error {
		a, err = h.account.TransferFunds(txCtx, dto.TransferFundsArgs{
			FromUserID: userID,
			ToUserID:   req.ToUserID,
			Amount:     req.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		h.log.Error(err.Error())

		switch {
		case errors.Is(err, domain.ErrSelfTransfer):
			writeError(w, http.StatusBadRequest, err)
			return
		case errors.Is(err, domain.ErrInvalidMajorAmount):
			writeError(w, http.StatusBadRequest, err)
			return
		case errors.Is(err, domain.ErrZeroAmount):
			writeError(w, http.StatusBadRequest, err)
			return
		case errors.Is(err, domain.ErrNotEnoughFunds):
			writeError(w, http.StatusBadRequest, err)
			return
		}

		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeOK(w, Account{
		ID:      a.ID,
		UserID:  a.UserID,
		Balance: a.Balance.String(),
	})
}
