package v1

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/ysomad/avito-internship-task/internal/domain"
)

func (h *handler) GetAccount(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	a, err := h.account.GetByUserID(r.Context(), userID)
	if err != nil {
		h.log.Error(err.Error())

		if errors.Is(err, domain.ErrAccountNotFound) {
			writeError(w, http.StatusNotFound, err, nil)
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
