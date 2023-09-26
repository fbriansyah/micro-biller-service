package application

import (
	"context"
	"errors"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
	dbill "github.com/fbriansyah/micro-biller-service/internal/application/domain/biller"
	"github.com/fbriansyah/micro-biller-service/internal/port"
)

var (
	ErrorBillAlreadyPaid = errors.New("bill already paid")
	ErrorInvalidAmount   = errors.New("invalid amount")
)

type BillerService struct {
	db port.DatabasePort
}

// NewBillerService create new service and need database port as parameter.
func NewBillerService(dbport port.DatabasePort) *BillerService {
	return &BillerService{
		db: dbport,
	}
}

// Inquiry get bill data using billNumber. This function will check whether bill already paid or not
func (s *BillerService) Inquiry(billNumber string) (dbill.Bill, error) {
	bill, err := s.db.GetBillingByNumber(context.Background(), billNumber)
	if err != nil {
		return dbill.Bill{}, err
	}

	if bill.IsPayed {
		return dbill.Bill{}, ErrorBillAlreadyPaid
	}

	return dbill.Bill{
		BillNumber:  billNumber,
		Name:        bill.Name,
		BaseAmount:  bill.BaseAmount,
		FineAmount:  bill.FineAmount,
		TotalAmount: bill.TotalAmount,
	}, nil
}

func (s *BillerService) Payment(updateBill dbill.Bill, refferenceNumber string) (dbill.Transaction, error) {
	bill, err := s.Inquiry(updateBill.BillNumber)
	if err != nil {
		return dbill.Transaction{}, err
	}

	if bill.TotalAmount != updateBill.TotalAmount {
		return dbill.Transaction{}, ErrorInvalidAmount
	}

	arg := db.PayBillParams{
		RefferenceNumber: refferenceNumber,
		TotalAmount:      updateBill.TotalAmount,
		BillNumber:       updateBill.BillNumber,
	}

	paidBill, err := s.db.PayBill(context.Background(), arg)
	if err != nil {
		return dbill.Transaction{}, err
	}

	return dbill.Transaction{
		RefferenceNumber: paidBill.RefferenceNumber,
		Billing:          updateBill,
		CreatedAt:        paidBill.PayTimestampt.Time,
	}, nil
}
