package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ysomad/avito-internship-task/internal/atomic"
)

type transactor struct {
	pool *pgxpool.Pool
}

func NewTransactor(p *pgxpool.Pool) *transactor {
	return &transactor{
		pool: p,
	}
}

func (t *transactor) RunAtomic(ctx context.Context, txFunc func(ctx context.Context) error) error {
	return atomic.Run(ctx, t.pool, txFunc)
}
