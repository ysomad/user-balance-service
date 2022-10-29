package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ysomad/avito-internship-task/pgclient"
)

type database struct {
	*pgclient.Client
}

func NewDB(c *pgclient.Client) *database {
	return &database{c}
}

func (db *database) WithinTransaction(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("db.Pool.Begin: %w", err)
	}

	if err = txFunc(withTx(ctx, tx)); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return rbErr
		}

		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (db *database) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx := txFromContext(ctx); tx != nil {
		return tx.Query(ctx, sql, args...)
	}
	return db.Pool.Query(ctx, sql, args...)
}

func (db *database) exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if tx := txFromContext(ctx); tx != nil {
		return tx.Exec(ctx, sql, args...)
	}
	return db.Pool.Exec(ctx, sql, args...)
}

func (db *database) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx := txFromContext(ctx); tx != nil {
		return tx.QueryRow(ctx, sql, args...)
	}
	return db.Pool.QueryRow(ctx, sql, args...)
}
