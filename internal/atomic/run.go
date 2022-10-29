package atomic

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Runnner interface {
	RunAtomic(ctx context.Context, txFunc func(ctx context.Context) error) error
}

type StarterWithOpts interface {
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

type Starter interface {
	Begin(context.Context) (pgx.Tx, error)
}

// Run runs txFunx within shared transaction via context.
// Must be wrapped by storage implementation of Atomic interface.
//
// Implementation example:
//
//	func (s *storage) RunAtomic(ctx context.Context, txFunc func(ctx context.Context) error) error {
//			return atomic.Run(ctx, s.Pool, txFunc)
//		}
func Run(ctx context.Context, db Starter, txFunc func(ctx context.Context) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("atomic: begin transaction - %w", err)
	}

	return run(ctx, tx, txFunc)
}

func RunWithOptions(ctx context.Context, db StarterWithOpts, opts pgx.TxOptions, txFunc func(ctx context.Context) error) error {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("atomic: begin transaction - %w", err)
	}

	return run(ctx, tx, txFunc)
}

func run(ctx context.Context, tx pgx.Tx, txFunc func(ctx context.Context) error) error {
	if err := txFunc(withTx(ctx, tx)); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("atomic: rollback - : %w", rbErr)
		}

		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("atomic: commit - %w", err)
	}

	return nil
}
