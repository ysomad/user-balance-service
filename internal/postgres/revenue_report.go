package postgres

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/pgxatomic"
)

type revenueReportRepo struct {
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
	table   string
}

func NewRevenueReportRepo(p pgxatomic.Pool, b sq.StatementBuilderType) *revenueReportRepo {
	return &revenueReportRepo{
		pool:    p,
		builder: b,
		table:   "revenue_report",
	}
}

func (r *revenueReportRepo) GetOrCreate(ctx context.Context, userID uuid.UUID) (domain.RevenueReport, error) {
	sql, args, err := r.builder.
		Insert(r.table + " as r").
		Columns("account_id").
		Values(sq.Expr("(SELECT id FROM account WHERE user_id = ?)", userID)).
		Suffix("ON CONFLICT (account_id) DO UPDATE").
		Suffix("SET account_id = EXCLUDED.account_id").
		Suffix("WHERE r.account_id = EXCLUDED.account_id").
		Suffix("RETURNING id, account_id, created_at").
		ToSql()
	if err != nil {
		return domain.RevenueReport{}, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return domain.RevenueReport{}, err
	}

	rr, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.RevenueReport])
	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.NotNullViolation {
			return domain.RevenueReport{}, fmt.Errorf(
				"account with user id not found: %s - %w", err.Error(), domain.ErrAccountNotFound)
		}

		return domain.RevenueReport{}, err
	}

	return rr, nil
}
