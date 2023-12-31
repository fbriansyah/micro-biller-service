package memorydb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
	"github.com/google/uuid"
)

func (*MemoryDatabase) CheckBill(ctx context.Context, arg db.CheckBillParams) (db.Billing, error) {

	bill, ok := (*data)[arg.BillNumber]

	if !ok {
		return db.Billing{}, errors.New("bill not found")
	}

	if bill.RefferenceNumber != arg.RefferenceNumber {
		return db.Billing{}, errors.New("reff not found")
	}

	return bill, nil
}

func (*MemoryDatabase) CreateBilling(ctx context.Context, arg db.CreateBillingParams) (db.Billing, error) {
	if _, ok := (*data)[arg.BillNumber]; ok {
		return db.Billing{}, errors.New("duplicate bill number")
	}

	bill := db.Billing{
		ID:               uuid.New(),
		BillNumber:       arg.BillNumber,
		Name:             arg.Name,
		BaseAmount:       arg.BaseAmount,
		FineAmount:       arg.FineAmount,
		TotalAmount:      arg.TotalAmount,
		CreatedAt:        time.Now(),
		PayTimestampt:    sql.NullTime{},
		RefferenceNumber: "",
		IsPayed:          false,
	}

	(*data)[bill.BillNumber] = bill

	return bill, nil
}

func (*MemoryDatabase) GetBillingByNumber(ctx context.Context, billNumber string) (db.Billing, error) {
	if bill, ok := (*data)[billNumber]; ok {
		return bill, nil
	}

	return db.Billing{}, errors.New("bill not found")
}

func (*MemoryDatabase) PayBill(ctx context.Context, arg db.PayBillParams) (db.Billing, error) {
	bill, ok := (*data)[arg.BillNumber]
	if !ok {
		return db.Billing{}, errors.New("bill not found")
	}

	bill.RefferenceNumber = arg.RefferenceNumber
	bill.IsPayed = true
	bill.PayTimestampt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	(*data)[arg.BillNumber] = bill

	return bill, nil
}
