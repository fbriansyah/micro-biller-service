--name: CreateBilling :one
INSERT INTO billings 
    (name, bill_number, base_amount, fine_amount, total_amount, createdAt)
VALUES (
    $1,$2,$3,$4,$5,now()
) RETURNING *;

--name: GetBillingByNumber
SELECT * FROM billings
WHERE bill_number
