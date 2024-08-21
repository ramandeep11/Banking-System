package db

import (
	"context"
	"database/sql"
	"simplebank/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(0, 27),

		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, entry.AccountID, arg.AccountID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {

	entry1 := createRandomEntry(t)

	entry2, _ := testQueries.GetEntry(context.Background(), entry1.ID)

	require.Equal(t, entry1.AccountID, entry2.AccountID)

}

func TestUpdateEntry(t *testing.T) {

	entry1 := createRandomEntry(t)

	arg := UpdateEntryParams{
		ID:     entry1.AccountID,
		Amount: util.RandomMoney(),
	}

	err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)

}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.AccountID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.AccountID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	// createRandomEntry(t)
	for i:=0;i<1;i++ {
		ent1 := createRandomEntry(t)
		print(ent1.AccountID)
	}

	arg1:= ListEntriesParams {
		Limit: 5,
		Offset: 5,
	}

	entries ,err := testQueries.ListEntries(context.Background(),arg1)

	require.NoError(t, err)
	require.Len(t, entries,5)

	for _, entry := range entries {
		require.NotEmpty(t,entry)
	}
}