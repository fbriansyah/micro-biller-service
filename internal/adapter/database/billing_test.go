package db

import (
	"context"
	"testing"

	"github.com/fbriansyah/micro-biller-service/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomBilling(t *testing.T) Billing {
	arg := CreateBillingParams{
		Name:       util.RandomString(8),
		BillNumber: util.RandomBillNumber(),
		BaseAmount: util.RandomMoney(),
		FineAmount: util.RandomInt(100, 999),
	}
	arg.TotalAmount = arg.BaseAmount + arg.FineAmount

	bill, err := testQueries.CreateBilling(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bill)

	require.Equal(t, arg.Name, bill.Name)
	require.Equal(t, arg.BillNumber, bill.BillNumber)
	require.Equal(t, arg.BaseAmount, bill.BaseAmount)
	require.Equal(t, arg.FineAmount, bill.FineAmount)
	require.Equal(t, arg.TotalAmount, bill.TotalAmount)
	require.Equal(t, false, bill.IsPayed)
	require.Equal(t, "", bill.RefferenceNumber)

	require.Empty(t, bill.PayTimestampt)
	require.NotZero(t, bill.CreatedAt)

	return bill
}

func TestCreateRandomBilling(t *testing.T) {
	CreateRandomBilling(t)
}

func TestGetBillingByNumber(t *testing.T) {
	bill1 := CreateRandomBilling(t)

	bill2, err := testQueries.GetBillingByNumber(context.Background(), bill1.BillNumber)
	require.NoError(t, err)
	require.NotEmpty(t, bill2)

	require.Equal(t, bill1.Name, bill2.Name)
	require.Equal(t, bill1.BillNumber, bill2.BillNumber)
	require.Equal(t, bill1.BaseAmount, bill2.BaseAmount)
	require.Equal(t, bill1.FineAmount, bill2.FineAmount)
	require.Equal(t, bill1.TotalAmount, bill2.TotalAmount)
	require.Equal(t, bill1.IsPayed, bill2.IsPayed)
	require.Equal(t, bill1.RefferenceNumber, bill2.RefferenceNumber)
}

func TestPayBill(t *testing.T) {
	bill1 := CreateRandomBilling(t)

	reffnum := util.RandomRefferenceNumber()

	arg := PayBillParams{
		RefferenceNumber: reffnum,
		BillNumber:       bill1.BillNumber,
		TotalAmount:      bill1.TotalAmount,
	}

	bill2, err := testQueries.PayBill(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, true, bill2.IsPayed)
	require.Equal(t, reffnum, bill2.RefferenceNumber)

	require.NotZero(t, bill2.PayTimestampt)
}

func TestCheckBill(t *testing.T) {
	bill1 := CreateRandomBilling(t)

	reffnum := util.RandomRefferenceNumber()

	arg := PayBillParams{
		RefferenceNumber: reffnum,
		BillNumber:       bill1.BillNumber,
		TotalAmount:      bill1.TotalAmount,
	}

	_, err := testQueries.PayBill(context.Background(), arg)
	require.NoError(t, err)

	bill2, err := testQueries.CheckBill(context.Background(), CheckBillParams{
		RefferenceNumber: arg.RefferenceNumber,
		BillNumber:       arg.BillNumber,
		TotalAmount:      arg.TotalAmount,
	})
	require.NoError(t, err)

	require.Equal(t, true, bill2.IsPayed)
	require.Equal(t, reffnum, bill2.RefferenceNumber)

	require.NotZero(t, bill2.PayTimestampt)
}
