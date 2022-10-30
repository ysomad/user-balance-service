package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"

	"github.com/ysomad/avito-internship-task/internal/pkg/atomic"
	"github.com/ysomad/avito-internship-task/internal/pkg/pgclient"
)

type transactionRepo struct {
	*pgclient.Client
	table string
}

func (r *transactionRepo) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return atomic.Query(ctx, r.Pool, sql, args...)
}

func NewTransactionRepo(c *pgclient.Client) *transactionRepo {
	return &transactionRepo{c, "transaction"}
}

func (r *transactionRepo) Create(ctx context.Context, args dto.CreateTransactionArgs) (*domain.Transaction, error) {
	sql, sqlArgs, err := r.Builder.
		Insert(r.table).
		Columns("user_id, comment, from_user_id, amount, operation").
		Values(args.UserID, args.Comment, args.FromUserID, args.Amount, args.Operation).
		Suffix("RETURNING id, user_id, comment, from_user_id, amount, operation, completed_at").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.query(ctx, sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Transaction])
	if err != nil {
		return nil, err
	}

	return &t, nil
}
