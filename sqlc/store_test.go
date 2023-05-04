package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	//run  n concurrent transfer transaction
	for i:=0; i<n;i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransfreTxParams{
				FromAccountID:		account1.ID,
				ToAccountID:		account2.ID,
				Amount:				amount,
			})
			errs <- err
			results <- result
		}()
	}

}