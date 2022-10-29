package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errInvalidMajorFormat     = errors.New("invalid major format, must be '13.37'")
	errInvalidSmallUnitFormat = errors.New("small unit format cannot be more than 2 characters")
)

type amount uint64

// NewAmount returns amount with validated value from format "54.22"
// and converted into uint64 of minor units.
//
// For example 52.22 amount is given as input but 5222 will be stored.
func NewAmount(major string) (amount, error) {
	units := strings.Split(major, ".")
	if len(units) != 2 {
		return 0, errInvalidMajorFormat
	}

	if len(units[1]) > 2 {
		return 0, errInvalidSmallUnitFormat
	}

	largeUnit, err := strconv.ParseUint(units[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid large unit: %w", err)
	}

	smallUnit, err := strconv.ParseUint(units[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid small unit: %w", err)
	}

	return amount((largeUnit * 100) + smallUnit), nil
}

func (a amount) MajorUnits() string {
	return strconv.FormatUint(a.LargeUnit(), 10) + "." + strconv.FormatUint(a.SmallUnit(), 10)
}

func (a amount) UInt64() uint64    { return uint64(a) }
func (a amount) LargeUnit() uint64 { return uint64(a) / 100 }
func (a amount) SmallUnit() uint64 { return uint64(a) % 100 }
