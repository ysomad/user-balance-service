package postgres

import "github.com/ysomad/avito-internship-task/internal/pkg/sort"

func orderByClause(column string, order sort.Order) string {
	return column + " " + string(order)
}
