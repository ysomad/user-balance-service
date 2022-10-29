package service

import "context"

type transactor interface {
	WithinTransaction(context.Context, func(ctx context.Context) error) error
}
