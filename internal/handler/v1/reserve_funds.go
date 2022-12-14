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
		writeError(w, http.StatusBadRequest, errInvalidRequestBody)
		return
	}

	var res *domain.Reservation

	err = h.atomic.Run(r.Context(), func(txCtx context.Context) error {
		res, err = h.account.ReserveFunds(txCtx, dto.ReserveFundsArgs{
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
			writeError(w, http.StatusBadRequest, err)
			return
		case errors.Is(err, domain.ErrZeroReserveAmount):
			writeError(w, http.StatusBadRequest, err)
			return
		}

		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeOK(w, &Reservation{
		ID:              res.ID,
		Amount:          res.Amount.String(),
		CreatedAt:       res.CreatedAt,
		DeclaredAt:      res.DeclaredAt,
		ServiceID:       res.ServiceID,
		OrderID:         res.OrderID,
		RevenueReportID: res.RevenueReportID,
		Status:          ReservationStatus(res.Status.String()),
	})
}
