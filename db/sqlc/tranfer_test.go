package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ThoPham02/simp_bank/util"
	"github.com/stretchr/testify/require"
)

func CreateTransfer(t *testing.T) Transfer {
	account1 := CreateAccount(t)
	account2 := CreateAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomInt(0, 100),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreateAt)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateTransfer(t)

	transfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.Equal(t, transfer.ID, transfer.ID)
	require.Equal(t, transfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transfer.Amount, transfer.Amount)
	require.WithinDuration(t, transfer.CreateAt, transfer.CreateAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateTransfer(t)
	}

	transfers, err := testQueries.ListTransfers(context.Background())

	require.NoError(t, err)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.NotEmpty(t, transfer.ID)
		require.NotEmpty(t, transfer.CreateAt)
	}
}

func TestUpdateTransfer(t *testing.T) {
	newTransfer := CreateTransfer(t)
	newAccount1 := CreateAccount(t)
	newAccount2 := CreateAccount(t)

	arg := UpdateTransferParams{
		ID:            newTransfer.ID,
		FromAccountID: newAccount1.ID,
		ToAccountID:   newAccount2.ID,
		Amount:        util.RandomInt(0, 100),
	}

	err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
}
func TestDeleteTransfer(t *testing.T) {
	newTransfer := CreateTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), newTransfer.ID)
	require.NoError(t, err)

	transfer, err := testQueries.GetTransfer(context.Background(), newTransfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer)
}
