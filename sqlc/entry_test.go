package db

import (
	"context"
	"db/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

/*
func TestGetEntry(t *testing.T) {

	entry, err := testQueries.GetEntry(context.Background(), util.RandomInt)

	require.NoError(t, err)
	require.NotEmpty(t, err)

} */
