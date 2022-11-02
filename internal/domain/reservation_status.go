package domain

import "strings"

type ReservationStatus uint8

const (
	ReservationStatusUndefined = iota
	ReservationStatusActive
	ReservationStatusDeclared
	ReservationStatusCanceled
)

func NewReservationStatus(s string) ReservationStatus {
	switch strings.ToUpper(s) {
	case "ACTIVE":
		return ReservationStatusActive
	case "DECLARED":
		return ReservationStatusDeclared
	case "CANCELED":
		return ReservationStatusCanceled
	}
	return ReservationStatusUndefined
}

func (s ReservationStatus) String() string {
	switch s {
	case ReservationStatusActive:
		return "ACTIVE"
	case ReservationStatusDeclared:
		return "DECLARED"
	case ReservationStatusCanceled:
		return "CANCELED"
	}
	return "UNDEFINED"
}
