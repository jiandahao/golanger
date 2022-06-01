package dbutils

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"gorm.io/gorm"
)

func TestTransaction(t *testing.T) {
	tests := []struct {
		name   string
		result error
	}{
		{
			name:   "without error",
			result: nil,
		},
		{
			name:   "with error",
			result: fmt.Errorf("something wrong"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tester := newTransactionTester(10, tt.result)

			ctx := context.Background()
			tester.runTest(ctx)

			// make sure, the transaction will be initialized only once.
			if tester.beginCounter != 1 {
				t.Fatalf("should begin a transaction once, but got %v times", tester.beginCounter)
			}

			// make sure, transaction commit or rollback once
			if !((tester.commitCounter == 1 && tester.rollbackCounter == 0) ||
				(tester.commitCounter == 0 && tester.rollbackCounter == 1)) {
				t.Fatalf("should commit or rollback only once,but got commit %v times, rollback %v times", tester.commitCounter, tester.rollbackCounter)
			}
		})
	}
}

type transactionTester struct {
	conn            *gorm.DB
	beginCounter    int
	commitCounter   int
	rollbackCounter int

	curDepth int
	maxDepth int

	result error
}

func newTransactionTester(maxDepth int, result error) *transactionTester {
	dbConn := &gorm.DB{}

	tester := &transactionTester{conn: dbConn, maxDepth: maxDepth}

	monkey.PatchInstanceMethod(reflect.TypeOf(dbConn), "Begin", func(_ *gorm.DB, opts ...*sql.TxOptions) *gorm.DB {
		tester.beginCounter = tester.beginCounter + 1
		return dbConn
	})

	monkey.PatchInstanceMethod(reflect.TypeOf(dbConn), "Commit", func(_ *gorm.DB) *gorm.DB {
		tester.commitCounter = tester.commitCounter + 1
		return dbConn
	})

	monkey.PatchInstanceMethod(reflect.TypeOf(dbConn), "Rollback", func(_ *gorm.DB) *gorm.DB {
		tester.rollbackCounter = tester.rollbackCounter + 1
		return dbConn
	})

	return tester
}

func (t *transactionTester) runTest(ctx context.Context) error {
	return Transaction(ctx, t.conn, func(ctx context.Context, tx *gorm.DB) error {
		if t.curDepth >= t.maxDepth {
			return t.result
		}

		t.curDepth++
		return t.runTest(ctx)
	}, func(curDepth int) func() {
		return func() {
			fmt.Println("callback from depth:", curDepth)
		}
	}(t.curDepth))
}
