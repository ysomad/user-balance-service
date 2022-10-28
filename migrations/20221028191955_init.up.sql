CREATE TABLE IF NOT EXISTS wallet (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    user_id uuid UNIQUE NOT NULL,
    balance bigint DEFAULT 0 NOT NULL
);

CREATE TABLE IF NOT EXISTS reserve_wallet (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    wallet_id uuid UNIQUE NOT NULL REFERENCES wallet (id),
    balance bigint DEFAULT 0 NOT NULL
);

CREATE TYPE transaction_operation AS enum (
    'WITHDRAW',
    'DEPOSIT'
);

CREATE TABLE IF NOT EXISTS transaction (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    user_id uuid NOT NULL,
    comment text NOT NULL,
    from_user_id uuid NOT NULL,
    amount bigint NOT NULL,
    operation transaction_operation NOT NULL,
    completed_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);