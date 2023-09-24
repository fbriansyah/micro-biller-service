package application

import (
	"errors"

	dbill "github.com/fbriansyah/micro-biller-service/internal/application/domain/biller"
	"github.com/fbriansyah/micro-biller-service/internal/port"
)

type BillerService struct {
	db port.DatabasePort
}

func NewBillerService(dbport port.DatabasePort) *BillerService {
	return &BillerService{
		db: dbport,
	}
}

func (s *BillerService) Inquiry(billNumber string) (dbill.Bill, error) {
	bill, err := s.db.GetBillingByNumber(billNumber)
	if err != nil {
		return dbill.Bill{}, err
	}

	if bill.IsPayed.Valid && bill.IsPayed.Bool {
		return dbill.Bill{}, errors.New("bill already payed")
	}

	return dbill.Bill{
		BillNumber:  billNumber,
		Name:        bill.Name,
		BaseAmount:  bill.BaseAmount,
		FineAmount:  bill.FineAmount,
		Totalamount: bill.TotalAmount,
	}, nil
}
