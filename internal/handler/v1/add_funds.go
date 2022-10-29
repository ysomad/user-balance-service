package v1

import (
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
		h.log.Error(err.Error())
		// TODO: handle specific errors

		writeError(w, http.StatusInternalServerError, domain.ErrFundsNotAdded, nil)
		return
	}

	writeJSON(w, http.StatusOK, Wallet{
		ID:       wallet.ID,
		UserID:   wallet.UserID,
		Balance:  wallet.Balance.MajorUnits(),
		Reserved: wallet.Reserved.MajorUnits(),
	})
}
