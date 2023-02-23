package db

import (
	"context"
	"testing"
	"time"

	"github.com/ThoPham02/simp_bank/util"
	"github.com/stretchr/testify/require"
)

func CreateEntry(t *testing.T) Entry {
	account := CreateAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(0, 200),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.CreateAt)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateAccount(t)
}

func TestGetEntry(t *testing.T) {
	newEntry := CreateEntry(t)

	entry, err := testQueries.GetEntry(context.Background(), newEntry.ID)
	require.NoError(t, err)

	require.Equal(t, entry.ID, newEntry.ID)
	require.Equal(t, entry.AccountID, newEntry.AccountID)
	require.Equal(t, entry.Amount, newEntry.Amount)
	require.WithinDuration(t, entry.CreateAt, newEntry.CreateAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i:= 0; i <10; i++ {
		CreateEntry(t)
	}

	entries, err := testQueries.ListEntries(context.Background())

	require.NoError(t, err)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.NotEmpty(t, entry.ID)
        require.NotEmpty(t, entry.CreateAt)
	}
}

func TestUpdateEntry(t *testing.T) {
	newAccount := CreateAccount(t)
	newEntry := CreateEntry(t)

	arg := UpdateEntryParams{
		ID:        newEntry.ID,
		AccountID: newAccount.ID,
		Amount:    util.RandomInt(0, 200),
	}

	err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
}

func TestDeleteEntry(t *testing.T) {
	newEntry := CreateEntry(t)
	err := testQueries.DeleteEntry(context.Background(), newEntry.ID)

	require.NoError(t, err)
}
