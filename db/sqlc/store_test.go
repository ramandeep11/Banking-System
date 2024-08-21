package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/require"
)

func TestTansferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">>>before ", account1.Balance,account2.Balance )

	// run n concurrent tranfer transactions
	n := 5

	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	// result, err := store.TransferTx(context.Background(), TransferTxParams{
	// 	FromAccountID: account1.ID,
	// 	ToaccountID:   account2.ID,
	// 	amount:        amount,
	// })
	// if(err != nil){
	// 	fmt.Print(err)
	// }

	// fmt.Printf("%d" , result.FromAccountID.ID)
	for i := 0; i < n; i++ {
		txName:= fmt.Sprintf("tx %d ", i+1)

		if(i%2==0){
			go func() {
				ctx:= context.WithValue(context.Background(), txKey, txName)
				result, err := store.TransferTx(ctx, TransferTxParams{
					FromAccountID: account2.ID,
					ToAccountID:   account1.ID,
					Amount:        amount,
				})
	
				errs <- err
				results <- result
			}()
		}else{
			go func() {
				ctx:= context.WithValue(context.Background(), txKey, txName)
				result, err := store.TransferTx(ctx, TransferTxParams{
					FromAccountID: account1.ID,
					ToAccountID:   account2.ID,
					Amount:        amount,
				})
	
				errs <- err
				results <- result
			}()
		}
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
	}

	// check the new balance
}
