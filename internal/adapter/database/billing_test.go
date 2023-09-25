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
	require.Equal(t, false, bill.IsPayed.Bool)
	require.Equal(t, "", bill.RefferenceNumber)

	require.Empty(t, bill.PayTimestampt)
	require.NotZero(t, bill.CreatedAt)

	return bill
}

func TestCreateRandomBilling(t *testing.T) {
	CreateRandomBilling(t)
}
