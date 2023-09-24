package port

import (
	dbill "github.com/fbriansyah/micro-biller-service/internal/application/domain/biller"
)

type BillerServicePort interface {
	Inquiry(billNumber string) (dbill.Bill, error)
}
