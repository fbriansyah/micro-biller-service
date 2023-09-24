// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: billing.sql

package db

import (
	"context"
)

const checkBill = `-- name: CheckBill :one
SELECT id, bill_number, name, base_amount, fine_amount, total_amount, "createdAt", pay_timestampt, refference_number, "isPayed" FROM billings
WHERE 
    bill_number = $1 
AND refference_number = $2
AND total_amount = $3
LIMIT 1
`

type CheckBillParams struct {
	BillNumber       string `json:"bill_number"`
	RefferenceNumber string `json:"refference_number"`
	TotalAmount      int64  `json:"total_amount"`
}

func (q *Queries) CheckBill(ctx context.Context, arg CheckBillParams) (Billing, error) {
	row := q.db.QueryRowContext(ctx, checkBill, arg.BillNumber, arg.RefferenceNumber, arg.TotalAmount)
	var i Billing
	err := row.Scan(
		&i.ID,
		&i.BillNumber,
		&i.Name,
		&i.BaseAmount,
		&i.FineAmount,
		&i.TotalAmount,
		&i.CreatedAt,
		&i.PayTimestampt,
		&i.RefferenceNumber,
		&i.IsPayed,
	)
	return i, err
}

const createBilling = `-- name: CreateBilling :one
INSERT INTO billings 
    (name, bill_number, base_amount, fine_amount, total_amount, createdAt)
VALUES (
    $1,$2,$3,$4,$5,now()
) RETURNING id, bill_number, name, base_amount, fine_amount, total_amount, "createdAt", pay_timestampt, refference_number, "isPayed"
`

type CreateBillingParams struct {
	Name        string `json:"name"`
	BillNumber  string `json:"bill_number"`
	BaseAmount  int64  `json:"base_amount"`
	FineAmount  int64  `json:"fine_amount"`
	TotalAmount int64  `json:"total_amount"`
}

func (q *Queries) CreateBilling(ctx context.Context, arg CreateBillingParams) (Billing, error) {
	row := q.db.QueryRowContext(ctx, createBilling,
		arg.Name,
		arg.BillNumber,
		arg.BaseAmount,
		arg.FineAmount,
		arg.TotalAmount,
	)
	var i Billing
	err := row.Scan(
		&i.ID,
		&i.BillNumber,
		&i.Name,
		&i.BaseAmount,
		&i.FineAmount,
		&i.TotalAmount,
		&i.CreatedAt,
		&i.PayTimestampt,
		&i.RefferenceNumber,
		&i.IsPayed,
	)
	return i, err
}

const getBillingByNumber = `-- name: GetBillingByNumber :one
SELECT id, bill_number, name, base_amount, fine_amount, total_amount, "createdAt", pay_timestampt, refference_number, "isPayed" FROM billings
WHERE bill_number = $1 LIMIT 1
`

func (q *Queries) GetBillingByNumber(ctx context.Context, billNumber string) (Billing, error) {
	row := q.db.QueryRowContext(ctx, getBillingByNumber, billNumber)
	var i Billing
	err := row.Scan(
		&i.ID,
		&i.BillNumber,
		&i.Name,
		&i.BaseAmount,
		&i.FineAmount,
		&i.TotalAmount,
		&i.CreatedAt,
		&i.PayTimestampt,
		&i.RefferenceNumber,
		&i.IsPayed,
	)
	return i, err
}

const payBill = `-- name: PayBill :one
UPDATE billings
SET
    isPayed = true,
    pay_timestampt = now(),
    refference_number = $1
WHERE
    bill_number = $2
AND total_amount = $3
RETURNING id, bill_number, name, base_amount, fine_amount, total_amount, "createdAt", pay_timestampt, refference_number, "isPayed"
`

type PayBillParams struct {
	RefferenceNumber string `json:"refference_number"`
	BillNumber       string `json:"bill_number"`
	TotalAmount      int64  `json:"total_amount"`
}

func (q *Queries) PayBill(ctx context.Context, arg PayBillParams) (Billing, error) {
	row := q.db.QueryRowContext(ctx, payBill, arg.RefferenceNumber, arg.BillNumber, arg.TotalAmount)
	var i Billing
	err := row.Scan(
		&i.ID,
		&i.BillNumber,
		&i.Name,
		&i.BaseAmount,
		&i.FineAmount,
		&i.TotalAmount,
		&i.CreatedAt,
		&i.PayTimestampt,
		&i.RefferenceNumber,
		&i.IsPayed,
	)
	return i, err
}