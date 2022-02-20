package mysqltest

import (
	"database/sql"
	"testing"
)

func TestMysqltester(t *testing.T) {
	mysqltester, err := NewMysqltester()
	if err != nil {
		t.Errorf("TestMysqltester() error = %v", err)
		return
	}

	mysqltester.MustExec(`CREATE DATABASE if not exists test;`)

	rows, err := mysqltester.Query(`show databases`)
	if err != nil {
		t.Errorf("TestMysqltester() error = %v", err)
		return
	}

	columns, _ := rows.Columns()
	for rows.Next() {
		var result = make([]sql.NullString, len(columns))
		var input = make([]interface{}, len(columns))
		for i := range input {
			input[i] = &result[i]
		}

		if err := rows.Scan(input...); err != nil {
			t.Fatalf("TestMysqltester() failed to scan data: %v", err)
		}
	}
}
