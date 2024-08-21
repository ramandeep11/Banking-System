package db

import (
	"context"
	"database/sql"
	"simplebank/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestCreateAccount(t *testing.T) {
// 	arg := CreateAccountParams{
// 		Owner:    util.RandomOwner(),
// 		Balance:  util.RandomMoney(),
// 		Currency: util.RandomCurrency(),
// 	}

// 	account, err := testQueries.CreateAccount(context.Background(), arg)
// 	require.NoError(t,err)
// 	require.NotEmpty(t,account)

// 	require.Equal(t,arg.Owner,account.Owner)
// 	require.Equal(t,arg.Balance,account.Balance)
// 	require.Equal(t,arg.Currency,account.Currency)
// }

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, _ := testQueries.GetAccount(context.Background(), account1.ID)
	print(account2.ID)
	// require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

func TestUpdateAccount(t* testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID : account1.ID,
		Balance: util.RandomMoney(),
	}

	err := testQueries.UpdateAccount(context.Background(),arg)

	require.NoError(t ,err)
	account2, _ := testQueries.GetAccount(context.Background(), arg.ID)
	require.NotEmpty(t,account2)
	require.Equal(t, arg.Balance,account2.Balance)
}


func TestDeleteAccount(t* testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t ,err)

	account2 , err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListAccount(t* testing.T) {

	var lastaccount Account
	for i:=0 ;i<10; i++ {
		lastaccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner: lastaccount.Owner,
		Limit: 5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t,err)
	require.NotEmpty(t,accounts)

	for _, account := range accounts {
		require.NotEmpty(t,account)
	}
}