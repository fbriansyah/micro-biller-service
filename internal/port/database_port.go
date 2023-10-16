package port

import (
	"context"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
)

type DatabasePort interface {
	// CheckBill check payment status to given bill number
	CheckBill(ctx context.Context, arg db.CheckBillParams) (db.Billing, error)
	// CreateBilling create new billing in table billings
	CreateBilling(ctx context.Context, arg db.CreateBillingParams) (db.Billing, error)
	// Find billing for given bill number
	GetBillingByNumber(ctx context.Context, billNumber string) (db.Billing, error)
	// PayBill set billing status to payed
	PayBill(ctx context.Context, arg db.PayBillParams) (db.Billing, error)
}
