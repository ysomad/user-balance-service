CREATE TABLE IF NOT EXISTS account (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    user_id uuid UNIQUE NOT NULL,
    balance bigint DEFAULT 0 NOT NULL
);

CREATE TABLE IF NOT EXISTS revenue_report (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    account_id uuid UNIQUE NOT NULL REFERENCES account (id),
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS reservation (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    account_id uuid NOT NULL REFERENCES account (id),
    service_id uuid NOT NULL,
    order_id uuid NOT NULL,
    amount bigint DEFAULT 0 NOT NULL,
    is_declared boolean DEFAULT FALSE NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    declared_at timestamptz,
    revenue_report_id uuid REFERENCES revenue_report (id)
);

CREATE TYPE transaction_operation AS enum (
    'WITHDRAW',
    'DEPOSIT'
);

CREATE TABLE IF NOT EXISTS transaction (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
    account_id uuid NOT NULL REFERENCES account (id),
    comment text NOT NULL,
    amount bigint NOT NULL,
    operation transaction_operation NOT NULL,
    commited_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);