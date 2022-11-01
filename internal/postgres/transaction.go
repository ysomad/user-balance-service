package postgres

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ysomad/pgxatomic"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type transactionRepo struct {
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
	table   string
}

func NewTransactionRepo(p pgxatomic.Pool, b sq.StatementBuilderType) *transactionRepo {
	return &transactionRepo{
		pool:    p,
		builder: b,
		table:   "transaction",
	}
}

func (r *transactionRepo) Create(ctx context.Context, args dto.CreateTransactionArgs) (*domain.Transaction, error) {
	sql, sqlArgs, err := r.builder.
		Insert(r.table).
		Columns("account_id, comment, amount, operation").
		Values(sq.Expr("(SELECT id FROM account WHERE user_id = ?)", args.UserID), args.Comment, args.Amount, args.Operation).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Transaction])
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *transactionRepo) FindAllByUserID(ctx context.Context, args dto.FindTransactionListArgs) ([]domain.Transaction, error) {
	whereClause := sq.And{
		sq.Expr("account_id = (SELECT id FROM account WHERE user_id = ?)", args.UserID),
	}

	if (args.LastID != uuid.UUID{}) && (args.LastCommitedAt != time.Time{}) {
		whereClause = append(whereClause, sq.Expr("(id, commited_at) > (?, ?)", args.LastID, args.LastCommitedAt))
	}

	sb := r.builder.
		Select("*").
		From(r.table).
		Where(whereClause)

	switch {
	case args.Sorts.Amount != "":
		sb = sb.OrderBy(orderByClause("amount", args.Sorts.Amount))
	case args.Sorts.CommitedAt != "":
		sb = sb.OrderBy(orderByClause("commited_at", args.Sorts.CommitedAt))
	}

	sql, sqlArgs, err := sb.Limit(args.PageSize + 1).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	txs, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Transaction])
	if err != nil {
		return nil, err
	}

	return txs, nil
}
