CREATE TABLE "billings" (
    "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "bill_number" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "base_amount" bigint NOT NULL DEFAULT 1000,
    "fine_amount" bigint NOT NULL DEFAULT 1000,
    "total_amount" bigint NOT NULL DEFAULT 1000,
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    "pay_timestampt" timestamptz,
    "refference_number" varchar NOT NULL DEFAULT '',
    "is_payed" bool NOT NULL DEFAULT false
);
