package db

import (
	"context"
	"database/sql"
	"fmt"
)

/*
The first field Queries is a pointer to a type that contains SQL queries
to be executed against the database. This means that the Store struct has

	access to these queries and can use them to perform CRUD (create, read, update, delete)
	operations on the database.

The second field db is a pointer to a database connection (*sql.DB),

	which is used to actually execute the SQL queries. This allows the Store struct to
	establish a connection to the database and execute SQL statements against it.
*/
type Store struct {
	*Queries
	db *sql.DB
}

/*
The NewStore function initializes a new *Store instance by creating a new Queries struct
using the New function (which is presumably defined elsewhere in the codebase),
and setting the db field to the provided *sql.DB. Finally, it returns the *Store instance.
This allows for easy access to all the SQL queries in the Queries struct via the *Store instance,
 while still maintaining a single database connection.
*/

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

/*
This is a method called execTx that belongs to a type SQLStore. It takes in a context object ctx and a function fn that takes in a pointer to Queries and returns an error. The purpose of this method is to execute the provided function within a database transaction.

First, the method begins a transaction on the database using the BeginTx method of the db object that is part of the SQLStore type. If there is an error starting the transaction, it is returned immediately.

Next, a new Queries object is created by calling the New method with the transaction object tx. The provided function fn is executed with this Queries object.

If there is an error executing the function, the transaction is rolled back using the Rollback method. If the rollback fails, an error is returned that combines the original error and the rollback error. Otherwise, the original error is returned.

If the function completes successfully, the transaction is committed using the Commit method of the tx object. If there is an error committing the transaction, it is returned immediately.
*/
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("txt err : %v, rb err: %v", err, rbErr)
		}
		return err

	}
	return tx.Commit()
}
