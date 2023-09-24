package biller

import "time"

type Bill struct {
	BillNumber  string
	Name        string
	BaseAmount  int64
	FineAmount  int64
	Totalamount int64
}

type Transaction struct {
	Billing   Bill
	CreatedAt time.Time
}
