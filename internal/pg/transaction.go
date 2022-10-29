package pg

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

func (db *database) CreateTransaction(ctx context.Context, args dto.CreateTransactionArgs) (*domain.Transaction, error) {
	sql, sqlArgs, err := db.Builder.
		Insert("transaction").
		Columns("user_id, comment, from_user_id, amount, operation").
		Values(args.UserID, args.Comment, args.FromUserID, args.Amount, args.Operation).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.query(ctx, sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Transaction])
	if err != nil {
		return nil, err
	}

	return &t, nil
}
