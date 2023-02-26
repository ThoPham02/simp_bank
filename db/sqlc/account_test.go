package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ThoPham02/simp_bank/util"
	"github.com/stretchr/testify/require"
)

func CreateAccount(t *testing.T) Account {
	user := CreateRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBallance(400),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, account.ID)
	require.NotEmpty(t, account.CreateAt)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAccount := CreateAccount(t)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.NoError(t, err)
	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Balance, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)
	require.WithinDuration(t, newAccount.CreateAt, account.CreateAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateAccount(t)
	}

	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  5,
		Offset: 5,
	})
	require.NoError(t, err)
	require.Equal(t, len(accounts), 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	newAccount := CreateAccount(t)
	newUser := CreateRandomUser(t)

	arg := UpdateAccountParams{
		ID:       newAccount.ID,
		Owner:    newUser.Username,
		Balance:  util.RandomBallance(400),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, newAccount.ID, account.ID)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := CreateAccount(t)

	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}
