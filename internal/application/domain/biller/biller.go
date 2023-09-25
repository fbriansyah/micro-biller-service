package biller

import "time"

type Bill struct {
	BillNumber  string
	Name        string
	BaseAmount  int64
	FineAmount  int64
	TotalAmount int64
}

type Transaction struct {
	Billing          Bill
	RefferenceNumber string
	CreatedAt        time.Time
}
