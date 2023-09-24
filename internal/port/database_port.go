package port

import (
	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
)

type DatabasePort interface {
	CheckBill(arg db.CheckBillParams) (db.Billing, error)
	CreateBilling(arg db.CreateBillingParams) (db.Billing, error)
	GetBillingByNumber(billNumber string) (db.Billing, error)
	PayBill(arg db.PayBillParams) (db.Billing, error)
}
