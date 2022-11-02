package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/ysomad/pgxatomic"

	"github.com/ysomad/avito-internship-task/internal"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/pkg/sort"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type transactionRepo struct {
	log     internal.Logger
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
	table   string
}

func NewTransactionRepo(l internal.Logger, p pgxatomic.Pool, b sq.StatementBuilderType) *transactionRepo {
	return &transactionRepo{
		log:     l,
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
	sb := r.builder.
		Select("*").
		From(r.table).
		Where(sq.And{
			sq.Expr("account_id = (SELECT id FROM account WHERE user_id = ?)", args.UserID),
			sq.GtOrEq{"id": args.PrevID},
		})

	// handle next pages
	// if (args.LastCommitedAt != time.Time{}) && (args.PrevID != uuid.UUID{}) {
	// 	whereClause := fmt.Sprintf("(commited_at, id) %s (?, ?)", paging.SeekSign(args.Sorts.CommitedAt.String()))
	// 	sb = sb.Where(sq.Expr(whereClause, args.LastCommitedAt, args.PrevID))
	// }

	switch {
	case args.Sorts.CommitedAt != sort.OrderEmpty:
		sb = sb.OrderBy(orderByClause("commited_at", args.Sorts.CommitedAt))
	case args.Sorts.Amount != sort.OrderEmpty:
		sb = sb.OrderBy(orderByClause("amount", args.Sorts.Amount))
	}

	sql, sqlArgs, err := sb.Limit(args.PageSize + 1).ToSql()
	if err != nil {
		return nil, err
	}

	r.log.Debug(sql)

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

func (r *transactionRepo) CreateMultiple(ctx context.Context, args []dto.CreateTransactionArgs) ([]domain.Transaction, error) {
	sb := r.builder.
		Insert(r.table).
		Columns("account_id, comment, amount, operation")

	for _, a := range args {
		sb = sb.Values(sq.Expr("(SELECT id FROM account WHERE user_id = ?)", a.UserID), a.Comment, a.Amount, a.Operation)
	}

	sql, sqlArgs, err := sb.Suffix("RETURNING *").ToSql()
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
