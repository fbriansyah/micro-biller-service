package port

import (
	"context"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
)

type DatabasePort interface {
	CheckBill(ctx context.Context, arg db.CheckBillParams) (db.Billing, error)
	CreateBilling(ctx context.Context, arg db.CreateBillingParams) (db.Billing, error)
	GetBillingByNumber(ctx context.Context, billNumber string) (db.Billing, error)
	PayBill(ctx context.Context, arg db.PayBillParams) (db.Billing, error)
}
