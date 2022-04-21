package utils

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

const (
	// transactionKey is the key for transaction values in Contexts.
	transactionKey key = iota + 1
)

type transactionContext struct {
	tx             *gorm.DB
	afterCommitFns []func() // callback functions called after committing the transaction.
	invalid        bool
}

// TransactionHandler transaction handler
type TransactionHandler func(ctx context.Context, tx *gorm.DB) error

// Transaction handles a task under transaction. It's NOT concurrency-safe.
//
// Transaction creates a transaction automatically if there is no available transaction in the context, otherwise reuse the existing one.
// The existing transaction instance is passed to the context.
//
// Example usage:
/*
	type Storage struct {
		conn *gorm.DB
	}

	func (s *Storage) CreateUser(ctx context.Context, user *User) error {
		return Transaction(ctx, s.conn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.Create(user).Error
		})
	}
*/
func Transaction(ctx context.Context, conn *gorm.DB, handler TransactionHandler, afterCommitFns ...func()) (err error) {
	var isCreator bool

	tc, exists := ctx.Value(transactionKey).(*transactionContext)
	if !exists {
		// there is no available transaction yet, initialize one
		tc = &transactionContext{
			tx:      conn.Begin(&sql.TxOptions{}),
			invalid: false,
		}

		isCreator = true
		ctx = context.WithValue(ctx, transactionKey, tc)
	}

	// transaction committed or rollbacked already.
	if tc.invalid {
		return fmt.Errorf("transactionContext is already invalid")
	}

	var panicked bool = true

	defer func() {
		tc, exists := ctx.Value(transactionKey).(*transactionContext)
		if !exists || tc.invalid {
			// transacton is committed or rollbacked already.
			fmt.Printf("transactionContext is already invalid\n")
			return
		}

		if panicked || err != nil {
			// Do rollback immediately once anything wrong happens.
			// Due to the user may forget to handle errors inside a transaction,
			// it's unsafe to leave this transaction to the creator to do rollback.
			if err := tc.tx.Rollback().Error; err != nil {
				fmt.Printf("failed to rollback the transaction: %v\n", err)
			}

			tc.invalid = true
			return
		}

		// it's the creator's duty to commit this transaction.
		if !isCreator {
			return
		}

		if err := tc.tx.Commit().Error; err != nil {
			fmt.Printf("failed to commit the transaction: %v\n", err)
		}

		for _, callback := range tc.afterCommitFns {
			callback()
		}

		tc.invalid = true
	}()

	if len(afterCommitFns) > 0 {
		for _, callback := range afterCommitFns {
			if callback == nil {
				return fmt.Errorf("invalid callback functions: nil")
			}
		}
		tc.afterCommitFns = append(tc.afterCommitFns, afterCommitFns...)
	}

	err = handler(ctx, tc.tx)

	panicked = false

	return err
}
