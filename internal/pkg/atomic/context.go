package atomic

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type txKey struct{}

// WithTx injects tx to context.
func withTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// TxFromContext extracts transaction from context.
func txFromContext(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return tx
	}
	return nil
}
