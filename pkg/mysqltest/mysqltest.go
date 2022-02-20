package mysqltest

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	testmysqld "github.com/lestrrat-go/test-mysqld"
)

// Mysqltester mysqld
type Mysqltester struct {
	mysqld *testmysqld.TestMysqld
	dbConn *sqlx.DB
}

// Explain represents the explain result.
type Explain struct {
	SQL          string
	ID           *string `db:"id"`
	SelectType   *string `db:"select_type"`
	Table        *string `db:"table"`
	Partitions   *string `db:"partitions"`
	Type         *string `db:"type"`
	PossibleKeys *string `db:"possible_keys"`
	Key          *string `db:"key"`
	KeyLen       *string `db:"key_len"`
	Ref          *string `db:"ref"`
	Rows         *string `db:"rows"`
	Filtered     *string `db:"filtered"`
	Extra        *string `db:"Extra"`
}

// NewMysqltester create a mysqld instance.
func NewMysqltester() (*Mysqltester, error) {
	config := testmysqld.NewConfig()
	mysqld, err := testmysqld.NewMysqld(config)
	if err != nil {
		return nil, err
	}

	dbConn, err := sqlx.Connect("mysql", mysqld.DSN())
	if err != nil {
		return nil, err
	}

	return &Mysqltester{
		mysqld: mysqld,
		dbConn: dbConn,
	}, nil
}

// MustExec (panic) runs MustExec using this database. Any placeholder parameters are replaced with supplied args.
func (m *Mysqltester) MustExec(query string, args ...interface{}) sql.Result {
	return m.dbConn.MustExec(query, args...)
}

// Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
func (m *Mysqltester) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.dbConn.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
func (m *Mysqltester) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.dbConn.Query(query, args...)
}

// Explain explain the sql.
func (m *Mysqltester) Explain(sql string, args ...interface{}) (*Explain, error) {
	explianResult := &Explain{SQL: sql}
	err := m.dbConn.Get(explianResult, fmt.Sprintf("explain %s", sql), args...)
	if err != nil {
		return nil, err
	}
	return explianResult, nil
}

// EnableGeneralLog sets global general_log to ON.
// After doing this, all receiving sql will be recorded in table `mysql.general_log`
func (m *Mysqltester) EnableGeneralLog() error {
	_, err := m.dbConn.Exec("SET GLOBAL general_log = 'ON'")
	if err != nil {
		return err
	}

	_, err = m.dbConn.Exec("SET GLOBAL log_output = 'table'")
	if err != nil {
		return err
	}
	return nil
}

// DSN creates a datasource name string that is appropriate for connecting to the database instance
func (m *Mysqltester) DSN() string {
	return m.mysqld.DSN()
}

// Stop explicitly stops the execution of mysqld.
func (m *Mysqltester) Stop() {
	m.mysqld.Stop()
}
