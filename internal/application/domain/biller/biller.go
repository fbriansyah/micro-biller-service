package biller

import "time"

type Bill struct {
	BillNumber  string `json:"bill_number"`
	Name        string `json:"name"`
	BaseAmount  int64  `json:"base_amount"`
	FineAmount  int64  `json:"fine_amount"`
	TotalAmount int64  `json:"total_amount"`
}

type Transaction struct {
	Billing          Bill      `json:"bill"`
	RefferenceNumber string    `json:"refference_number"`
	CreatedAt        time.Time `json:"created_at"`
}
