package port

import (
	dbill "github.com/fbriansyah/micro-biller-service/internal/application/domain/biller"
)

type BillerServicePort interface {
	// Inquiry find billing using bill number
	Inquiry(billNumber string) (dbill.Bill, error)
	// Payment update bill to payed and set refference number
	Payment(updateBill dbill.Bill, refferenceNumber string) (dbill.Transaction, error)
	// Advice check bill is already payed or not
	Advice(searchBill dbill.Bill, refferenceNumber string) (dbill.Transaction, error)
}
