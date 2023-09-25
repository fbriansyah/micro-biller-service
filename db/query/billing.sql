-- name: CreateBilling :one
INSERT INTO billings 
    (name, bill_number, base_amount, fine_amount, total_amount)
VALUES (
    $1,$2,$3,$4,$5
) RETURNING *;

-- name: GetBillingByNumber :one
SELECT * FROM billings
WHERE bill_number = $1 LIMIT 1;

-- name: PayBill :one
UPDATE billings
SET
    isPayed = true,
    pay_timestampt = now(),
    refference_number = $1
WHERE
    bill_number = $2
AND total_amount = $3
RETURNING *;

-- name: CheckBill :one
SELECT * FROM billings
WHERE 
    bill_number = $1 
AND refference_number = $2
AND total_amount = $3
LIMIT 1;