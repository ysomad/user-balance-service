package sort

import "strings"

type Order string

const (
	OrderASC  Order = "ASC"
	OrderDESC Order = "DESC"

	OrderEmpty   Order = ""
	OrderDefault Order = OrderDESC
)

// Default returns default order if Order is empty string,
// or returns order as it is.
func (o Order) Default() Order {
	if o == OrderEmpty {
		return OrderDefault
	}
	return o
}

func (o Order) String() string { return string(o) }

// NewOrder returns Order - ASC, DESC or empty string.
func NewOrder(s string) Order {
	s = strings.ToUpper(s)
	if s != "ASC" && s != "DESC" {
		s = ""
	}
	return Order(s)
}
