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

func (h *handler) CancelReservation(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	var (
		req CancelReservationRequest
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody)
		return
	}

	var res *domain.Reservation

	err = h.atomic.Run(r.Context(), func(txCtx context.Context) error {
		res, err = h.account.CancelReservation(txCtx, dto.RawCancelReservationArgs{
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
		case errors.Is(err, domain.ErrReservationNotFound):
			writeError(w, http.StatusNotFound, err)
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
