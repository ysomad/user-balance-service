CREATE TABLE IF NOT EXISTS account (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    user_id uuid UNIQUE NOT NULL,
    balance bigint DEFAULT 0 NOT NULL
);

CREATE TABLE IF NOT EXISTS reserve_account (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    account_id uuid UNIQUE NOT NULL REFERENCES account (id),
    balance bigint DEFAULT 0 NOT NULL
);

CREATE TYPE transaction_operation AS enum (
    'WITHDRAW',
    'DEPOSIT'
);

CREATE TABLE IF NOT EXISTS transaction (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    comment text NOT NULL,
    from_user_id uuid,
    amount bigint NOT NULL,
    operation transaction_operation NOT NULL,
    completed_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);