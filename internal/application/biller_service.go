package application

import (
	"context"
	"errors"

	dbill "github.com/fbriansyah/micro-biller-service/internal/application/domain/biller"
	"github.com/fbriansyah/micro-biller-service/internal/port"
)

var (
	ErrorBillAlreadyPaid = errors.New("bill already paid")
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
