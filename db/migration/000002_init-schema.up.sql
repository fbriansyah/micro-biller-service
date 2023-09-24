CREATE TABLE "billings" (
    "id" uuid PRIMARY KEY,
    "bill_number" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "base_amount" bigint NOT NULL DEFAULT 1000,
    "fine_amount" bigint NOT NULL DEFAULT 1000,
    "total_amount" bigint NOT NULL DEFAULT 1000,
    "createdAt" timestamptz NOT NULL DEFAULT 'now()',
    "pay_timestampt" timestamptz,
    "refference_number" varchar NOT NULL DEFAULT '',
    "isPayed" bool
);
