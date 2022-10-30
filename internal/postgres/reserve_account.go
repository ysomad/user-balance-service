package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/ysomad/avito-internship-task/internal/domain"

	"github.com/ysomad/avito-internship-task/internal/pkg/atomic"
	"github.com/ysomad/avito-internship-task/internal/pkg/pgclient"
)

type reserveAccountRepo struct {
	*pgclient.Client
	table string
}

func NewReserveAccountRepo(c *pgclient.Client) *reserveAccountRepo {
	return &reserveAccountRepo{c, "reserve_account"}
}

func (r *reserveAccountRepo) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return atomic.Query(ctx, r.Pool, sql, args...)
}

func (r *reserveAccountRepo) Upsert(ctx context.Context, accountID uuid.UUID, amount domain.Amount) (domain.AccountReserve, error) {
	sql, args, err := r.Builder.
		Insert(r.table+" as a").
		Columns("account_id, balance").
		Values(accountID, amount).
		Suffix("ON CONFLICT (account_id) DO UPDATE").
		Suffix("SET balance = a.balance + EXCLUDED.balance").
		Suffix("WHERE a.account_id = EXCLUDED.account_id").
		Suffix("RETURNING id, account_id, balance").
		ToSql()
	if err != nil {
		return domain.AccountReserve{}, err
	}

	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return domain.AccountReserve{}, err
	}

	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.AccountReserve])
	if err != nil {
		return domain.AccountReserve{}, err
	}

	return a, nil
}
