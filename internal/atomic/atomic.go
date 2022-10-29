// atomic package is set of pgx
package atomic

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// DB must be implemented by repository to be able to use atomic operations
// of all repository methods by passing tx through the context.
//
// Implementation example:
//
//	func (s *storage) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
//			return atomic.Exec(ctx, s.Pool, sql, args...)
//	}
//
// Or it's possible only implement needed interfaces, for example Querier if other methods not needed.
type DB interface {
	Querier
	QueryRower
	Executor
}

type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type QueryRower interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Executor interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func Query(ctx context.Context, q Querier, sql string, args ...any) (pgx.Rows, error) {
	if tx := txFromContext(ctx); tx != nil {
		return tx.Query(ctx, sql, args...)
	}
	return q.Query(ctx, sql, args...)
}

func Exec(ctx context.Context, e Executor, sql string, args ...any) (pgconn.CommandTag, error) {
	if tx := txFromContext(ctx); tx != nil {
		return tx.Exec(ctx, sql, args...)
	}
	return e.Exec(ctx, sql, args...)
}

func QueryRow(ctx context.Context, q QueryRower, sql string, args ...any) pgx.Row {
	if tx := txFromContext(ctx); tx != nil {
		return tx.QueryRow(ctx, sql, args...)
	}
	return q.QueryRow(ctx, sql, args...)
}
