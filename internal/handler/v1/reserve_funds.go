package v1

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

func (h *handler) ReserveFunds(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	var req ReserveFundsRequest
	if err := h.validate.Into(r.Body, &req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody, h.validate.Translate(err))
		return
	}

	a, err := h.account.ReserveFunds(r.Context(), dto.ReserveFundsArgs{
		UserID:    userID.String(),
		ServiceID: req.ServiceID.String(),
		OrderID:   req.OrderID.String(),
		Amount:    req.Amount,
	})
	if err != nil {
		h.log.Error(err.Error())

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
