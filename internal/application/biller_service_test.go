package application

import (
	"context"
	"testing"
	"time"

	db "github.com/fbriansyah/micro-biller-service/internal/adapter/database"
	dmbill "github.com/fbriansyah/micro-biller-service/internal/application/domain/biller"
	"github.com/fbriansyah/micro-biller-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomBillInMemory(t *testing.T) db.Billing {
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

func createRandomBill(t *testing.T) db.Billing {
	arg := db.CreateBillingParams{
		Name:       util.RandomString(8),
		BillNumber: util.RandomBillNumber(),
		BaseAmount: util.RandomMoney(),
		FineAmount: util.RandomInt(100, 999),
	}

	bill, err := testDBAdapter.CreateBilling(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bill)

	return bill
}

func TestInquiryInMemory(t *testing.T) {

	bill1 := createRandomBillInMemory(t)
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

func TestInquiry(t *testing.T) {

	bill1 := createRandomBill(t)
	require.NotEmpty(t, bill1)

	bill2, err := testServiceWithPostgresDB.Inquiry(bill1.BillNumber)
	require.NoError(t, err)
	require.NotEmpty(t, bill2)

	require.Equal(t, bill1.Name, bill2.Name)
	require.Equal(t, bill1.BillNumber, bill2.BillNumber)
	require.Equal(t, bill1.BaseAmount, bill2.BaseAmount)
	require.Equal(t, bill1.FineAmount, bill2.FineAmount)
	require.Equal(t, bill1.TotalAmount, bill2.TotalAmount)
}

func TestInquiryAlreadyPaidInMemory(t *testing.T) {
	bill1 := createRandomBillInMemory(t)
	require.NotEmpty(t, bill1)

	reffnum := util.RandomRefferenceNumber()

	testDBMemoryAdapter.PayBill(context.Background(), db.PayBillParams{
		RefferenceNumber: reffnum,
		BillNumber:       bill1.BillNumber,
		TotalAmount:      bill1.TotalAmount,
	})

	_, err := testServiceWithMemoryDB.Inquiry(bill1.BillNumber)
	require.Error(t, err, ErrorBillAlreadyPaid)
}

func TestInquiryAlreadyPaid(t *testing.T) {
	bill1 := createRandomBill(t)
	require.NotEmpty(t, bill1)

	reffnum := uuid.New().String()

	testDBAdapter.PayBill(context.Background(), db.PayBillParams{
		RefferenceNumber: reffnum,
		BillNumber:       bill1.BillNumber,
		TotalAmount:      bill1.TotalAmount,
	})

	_, err := testServiceWithPostgresDB.Inquiry(bill1.BillNumber)
	require.Error(t, err, ErrorBillAlreadyPaid)
}

func TestPaymentInMemory(t *testing.T) {
	bill1 := createRandomBillInMemory(t)
	require.NotEmpty(t, bill1)

	reff := util.RandomRefferenceNumber()

	arg := dmbill.Bill{
		BillNumber:  bill1.BillNumber,
		TotalAmount: bill1.TotalAmount,
	}

	transaction, err := testServiceWithMemoryDB.Payment(arg, reff)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, reff, transaction.RefferenceNumber)
	require.WithinDuration(t, time.Now(), transaction.CreatedAt, 24*time.Hour)
}

func TestPaymentInvalidAmountInMemory(t *testing.T) {
	bill1 := createRandomBillInMemory(t)
	require.NotEmpty(t, bill1)

	reff := util.RandomRefferenceNumber()

	arg := dmbill.Bill{
		BillNumber:  bill1.BillNumber,
		TotalAmount: bill1.TotalAmount + 2000,
	}

	transaction, err := testServiceWithMemoryDB.Payment(arg, reff)
	require.Error(t, err, ErrorInvalidAmount)
	require.Empty(t, transaction)
}

func TestAdviceInMemory(t *testing.T) {
	bill1 := createRandomBillInMemory(t)
	require.NotEmpty(t, bill1)

	reff := util.RandomRefferenceNumber()
	searchedBill := dmbill.Bill{
		BillNumber:  bill1.BillNumber,
		TotalAmount: bill1.TotalAmount,
	}
	testServiceWithMemoryDB.Payment(searchedBill, reff)

	transaction, err := testServiceWithMemoryDB.Advice(searchedBill, reff)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, reff, transaction.RefferenceNumber)
	require.Equal(t, bill1.TotalAmount, transaction.Billing.TotalAmount)
}
