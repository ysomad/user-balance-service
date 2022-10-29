package v1

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/domain"
)

func (h *handler) AddFunds(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	var req AddFundsRequest
	if err := h.validate.Into(r.Body, &req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody, h.validate.Translate(err))
		return
	}

	wallet, err := h.wallet.AddFunds(r.Context(), userID.String(), req.Amount)
	if err != nil {
		h.log.Errorf("v1 - AddFunds: %s", err.Error())

		switch {
		case errors.Is(err, domain.ErrInvalidMajorAmount):
			writeError(w, http.StatusBadRequest, errInvalidRequestBody, map[string]string{
				"amount": err.Error(),
			})
			return
		}

		writeError(w, http.StatusInternalServerError, domain.ErrFundsNotAdded, nil)
		return
	}

	writeOK(w, Wallet{
		ID:      wallet.ID,
		UserID:  wallet.UserID,
		Balance: wallet.Balance.MajorUnits(),
	})
}
