package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidMajorAmount = errors.New("invalid major amount format, must be '13.37'")
	ErrZeroAmount         = errors.New("amount must be greated than 0")
)

type Amount uint64

// NewAmount returns amount with validated value from format "54.22"
// and converted into uint64 of minor units.
//
// For example 52.22 amount is given as input but 5222 will be stored.
func NewAmount(major string) (Amount, error) {
	units := strings.Split(major, ".")
	if len(units) != 2 {
		return 0, ErrInvalidMajorAmount
	}

	if len(units[1]) > 2 {
		return 0, fmt.Errorf("invalid small unit format: %w", ErrInvalidMajorAmount)
	}

	largeUnit, err := strconv.ParseUint(units[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid large unit: %s, %w", err.Error(), ErrInvalidMajorAmount)
	}

	smallUnit, err := strconv.ParseUint(units[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid small unit: %s, %w", err.Error(), ErrInvalidMajorAmount)
	}

	return Amount((largeUnit * 100) + smallUnit), nil
}

func (a Amount) String() string {
	return strconv.FormatUint(a.LargeUnit(), 10) + "." + strconv.FormatUint(a.SmallUnit(), 10)
}

func (a Amount) UInt64() uint64    { return uint64(a) }
func (a Amount) LargeUnit() uint64 { return uint64(a) / 100 }
func (a Amount) SmallUnit() uint64 { return uint64(a) % 100 }
func (a Amount) IsZero() bool      { return uint64(a) == 0 }
