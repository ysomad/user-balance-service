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

func (h *handler) DeclareRevenue(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	var (
		req DeclareRevenueRequest
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errInvalidRequestBody, nil)
		return
	}

	var res *domain.Reservation

	err = h.tx.Run(r.Context(), func(txCtx context.Context) error {
		res, err = h.account.DeclareRevenue(txCtx, dto.DeclareRevenueArgs{
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
		case errors.Is(err, domain.ErrAccountNotFound):
			writeError(w, http.StatusBadRequest, domain.ErrAccountNotFound, nil)
			return
		case errors.Is(err, domain.ErrZeroAmount):
			writeError(w, http.StatusBadRequest, err, nil)
			return
		case errors.Is(err, domain.ErrReservationNotDeclared):
			writeError(w, http.StatusNotFound, err, nil)
			return
		}

		writeError(w, http.StatusInternalServerError, err, nil)
		return
	}

	writeOK(w, &DeclareRevenueResponse{
		DeclaredAmount:  res.Amount.String(),
		DeclaredAt:      res.DeclaredAt,
		Declared:        res.Declared,
		OrderID:         res.OrderID,
		RevenueReportID: res.RevenueReportID,
		ServiceID:       res.ServiceID,
	})
}
