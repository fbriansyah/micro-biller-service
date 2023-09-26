package port

import (
	dbill "github.com/fbriansyah/micro-biller-service/internal/application/domain/biller"
)

type BillerServicePort interface {
	Inquiry(billNumber string) (dbill.Bill, error)
	Payment(updateBill dbill.Bill, refferenceNumber string) (dbill.Transaction, error)
	Advice(searchBill dbill.Bill, refferenceNumber string) (dbill.Transaction, error)
}
