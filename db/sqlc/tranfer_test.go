package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ThoPham02/simp_bank/util"
	"github.com/stretchr/testify/require"
)

func CreateTranfer(t *testing.T) Tranfer {
	account1 := CreateAccount(t)
	account2 := CreateAccount(t)

	arg := CreateTranferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
        Amount:        util.RandomInt(0, 100),
	}

	tranfer, err := testQueries.CreateTranfer(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, tranfer.ID)
	require.NotEmpty(t, tranfer.CreateAt)

	require.Equal(t, tranfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, tranfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, tranfer.Amount, arg.Amount)

	return tranfer
}

func TestCreateTranfer(t *testing.T) {
	CreateTranfer(t)
}

func TestGetTranfer(t *testing.T) {
	tranfer := CreateTranfer(t)

    tranfer, err := testQueries.GetTranfer(context.Background(), tranfer.ID)

    require.NoError(t, err)
    require.Equal(t, tranfer.ID, tranfer.ID)
    require.Equal(t, tranfer.FromAccountID, tranfer.FromAccountID)
    require.Equal(t, tranfer.ToAccountID, tranfer.ToAccountID)
    require.Equal(t, tranfer.Amount, tranfer.Amount)
    require.WithinDuration(t, tranfer.CreateAt, tranfer.CreateAt, time.Second)
}

func TestListTranfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateTranfer(t)
	}

	tranfers, err := testQueries.ListTranfers(context.Background())

	require.NoError(t, err)
	for _, tranfer := range tranfers {
		require.NotEmpty(t, tranfer);
		require.NotEmpty(t, tranfer.ID)
        require.NotEmpty(t, tranfer.CreateAt)
	}
}

func TestUpdateTranfer(t *testing.T) { 
	newTranfer := CreateTranfer(t)
	newAccount1 := CreateAccount(t)
	newAccount2 := CreateAccount(t)

	arg := UpdateTranferParams{
        ID: newTranfer.ID,
		FromAccountID: newAccount1.ID,
		ToAccountID: newAccount2.ID,
		Amount: util.RandomInt(0, 100),
    }

    err := testQueries.UpdateTranfer(context.Background(), arg)

    require.NoError(t, err)
}
func TestDeleteTranfer(t *testing.T) {
	newTranfer := CreateTranfer(t)

	err := testQueries.DeleteTranfer(context.Background(), newTranfer.ID)
	require.NoError(t, err)

	tranfer, err := testQueries.GetTranfer(context.Background(), newTranfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, tranfer)
}
