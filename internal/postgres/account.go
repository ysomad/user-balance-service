package postgres

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ysomad/pgxatomic"

	"github.com/ysomad/avito-internship-task/internal/domain"

	"github.com/ysomad/avito-internship-task/internal/pkg/pgclient"
)

type accountRepo struct {
	*pgclient.Client
	table string
}

func NewAccountRepo(c *pgclient.Client) *accountRepo {
	return &accountRepo{c, "account"}
}

func (r *accountRepo) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return pgxatomic.Query(ctx, r.Pool, sql, args...)
}

func (r *accountRepo) UpdateOrCreate(ctx context.Context, t domain.DepositTransaction) (domain.Account, error) {
	sql, args, err := r.Builder.
		Insert(r.table+" as a").
		Columns("user_id, balance").
		Values(t.UserID(), t.Amount()).
		Suffix("ON CONFLICT (user_id) DO UPDATE").
		Suffix("SET balance = a.balance + EXCLUDED.balance").
		Suffix("WHERE a.user_id = EXCLUDED.user_id").
		Suffix("RETURNING id, user_id, balance").
		ToSql()
	if err != nil {
		return domain.Account{}, err
	}

	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return domain.Account{}, err
	}

	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Account])
	if err != nil {
		return domain.Account{}, err
	}

	return a, nil
}

func (r *accountRepo) Withdraw(ctx context.Context, userID uuid.UUID, amount domain.Amount) (domain.Account, error) {
	sql, args, err := r.Builder.
		Update(r.table).
		Set("balance", sq.Expr("balance - ?", amount.UInt64())).
		Where(sq.And{
			sq.Eq{"user_id": userID},
			sq.GtOrEq{"balance": amount.UInt64()},
		}).
		Suffix("RETURNING id, user_id, balance").
		ToSql()
	if err != nil {
		return domain.Account{}, err
	}

	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return domain.Account{}, err
	}

	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Account])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Account{}, domain.ErrNotEnoughFunds
		}

		return domain.Account{}, err
	}

	return a, nil
}

func (r *accountRepo) FindByUserID(ctx context.Context, userID uuid.UUID) (domain.AccountAggregate, error) {
	sql, args, err := r.Builder.
		Select("a.id, a.user_id, a.balance, sum(r.amount), count(*)").
		From("account a").
		LeftJoin("reservation r ON a.id = r.account_id").
		Where(sq.And{
			sq.Eq{"a.user_id": userID},
			sq.Eq{"r.is_declared": false},
		}).
		GroupBy("a.id").
		ToSql()
	if err != nil {
		return domain.AccountAggregate{}, err
	}

	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return domain.AccountAggregate{}, err
	}

	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.AccountAggregate])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.AccountAggregate{}, domain.ErrAccountNotFound
		}

		return domain.AccountAggregate{}, err
	}

	return a, nil
}
