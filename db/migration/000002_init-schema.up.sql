CREATE TABLE "billing" (
    "id" uuid PRIMARY KEY,
    "name" varchar NOT NULL,
    "base_amount" bigint DEFAULT 1000,
    "fine_amount" bigint DEFAULT 1000,
    "total_amount" bigint DEFAULT 1000,
    "createdAt" timestamptz DEFAULT 'now()',
    "pay_timestampt" timestamptz,
    "refference_number" varchar,
    "isPayed" bool
);
