package pg

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/ysomad/avito-internship-task/internal/domain"
)

func (db *database) UpsertWalletBalance(ctx context.Context, t domain.DepositTransaction) (domain.Wallet, error) {
	sql := `
INSERT INTO wallet as w (user_id, balance)
VALUES ($1, $2)
ON CONFLICT (user_id) DO UPDATE
SET balance = w.balance + EXCLUDED.balance
WHERE w.user_id = EXCLUDED.user_id
RETURNING id, user_id, balance
`

	rows, err := db.query(ctx, sql, t.UserID(), t.Amount())
	if err != nil {
		return domain.Wallet{}, err
	}

	wallet, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Wallet])
	if err != nil {
		return domain.Wallet{}, err
	}

	return wallet, nil
}
