package sort

import "strings"

type Order string

func NewOrder(s string) Order {
	s = strings.ToUpper(s)
	if s != "ASC" && s != "DESC" {
		s = ""
	}
	return Order(s)
}
