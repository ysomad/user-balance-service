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
	"github.com/ysomad/avito-internship-task/internal/pkg/pgclient"
	"github.com/ysomad/pgxatomic"
)

type revenueReportRepo struct {
	*pgclient.Client
	table string
}

func (r *revenueReportRepo) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return pgxatomic.Query(ctx, r.Pool, sql, args...)
}

func NewRevenueReportRepo(c *pgclient.Client) *revenueReportRepo {
	return &revenueReportRepo{c, "revenue_report"}
}

func (r *revenueReportRepo) GetOrCreate(ctx context.Context, userID uuid.UUID) (domain.RevenueReport, error) {
	sql, args, err := r.Builder.
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

	rows, err := r.query(ctx, sql, args...)
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
