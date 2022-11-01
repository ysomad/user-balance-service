package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ysomad/pgxatomic"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

type reservationRepo struct {
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
	table   string
}

func NewReservationRepo(p pgxatomic.Pool, b sq.StatementBuilderType) *reservationRepo {
	return &reservationRepo{
		pool:    p,
		builder: b,
		table:   "reservation",
	}
}

func (r *reservationRepo) Create(ctx context.Context, args dto.CreateReservationArgs) (*domain.Reservation, error) {
	sql, sqlArgs, err := r.builder.
		Insert(r.table).
		Columns("account_id, service_id, order_id, amount").
		Values(args.AccountID, args.ServiceID, args.OrderID, args.Amount).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Reservation])
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return nil, fmt.Errorf("account not found: %s - %w", err.Error(), domain.ErrAccountNotFound)
		}

		return nil, err
	}

	return &res, nil
}

func (r *reservationRepo) AddToRevenueReport(ctx context.Context, args dto.AddToRevenueReportArgs) (*domain.Reservation, error) {
	sql, sqlArgs, err := r.builder.
		Update(r.table).
		Set("is_declared", true).
		Set("declared_at", time.Now()).
		Set("revenue_report_id", args.RevenueReportID).
		Where(sq.And{
			sq.Expr("account_id = (SELECT id FROM account WHERE user_id = ?)", args.UserID),
			sq.Eq{"service_id": args.ServiceID},
			sq.Eq{"order_id": args.OrderID},
			sq.Eq{"amount": args.Amount},
			sq.Eq{"is_declared": false},
		}).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Reservation])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrReservationNotDeclared
		}

		return nil, err
	}

	return &res, nil
}
