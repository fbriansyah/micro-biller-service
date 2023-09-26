package application

import (
	"context"
	"testing"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
	"github.com/fbriansyah/micro-biller-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomBill(t *testing.T) db.Billing {
	arg := db.CreateBillingParams{
		Name:       util.RandomString(8),
		BillNumber: util.RandomBillNumber(),
		BaseAmount: util.RandomMoney(),
		FineAmount: util.RandomInt(100, 999),
	}

	bill, err := testDBMemoryAdapter.CreateBilling(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bill)

	return bill
}

func TestInquiryInMemory(t *testing.T) {

	bill1 := createRandomBill(t)
	require.NotEmpty(t, bill1)

	bill2, err := testServiceWithMemoryDB.Inquiry(bill1.BillNumber)
	require.NoError(t, err)
	require.NotEmpty(t, bill2)

	require.Equal(t, bill1.Name, bill2.Name)
	require.Equal(t, bill1.BillNumber, bill2.BillNumber)
	require.Equal(t, bill1.BaseAmount, bill2.BaseAmount)
	require.Equal(t, bill1.FineAmount, bill2.FineAmount)
	require.Equal(t, bill1.TotalAmount, bill2.TotalAmount)
}

func TestInquiryAlreadyPaidInMemory(t *testing.T) {
	bill1 := createRandomBill(t)
	require.NotEmpty(t, bill1)

	reffnum := uuid.New().String()

	testDBMemoryAdapter.PayBill(context.Background(), db.PayBillParams{
		RefferenceNumber: reffnum,
		BillNumber:       bill1.BillNumber,
		TotalAmount:      bill1.TotalAmount,
	})

	_, err := testServiceWithMemoryDB.Inquiry(bill1.BillNumber)
	require.Error(t, err, ErrorBillAlreadyPaid)
}
