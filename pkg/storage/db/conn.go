package dbutils

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config connection config
type Config struct {
	gorm.Config

	DSN          string
	MaxIdleConns int           // the maximum number of connections in the idle connection pool.
	MaxOpenConns int           // the maximum number of open connections to the database.
	MaxLifetime  time.Duration // the maximum amount of time a connection may be reused.
	MaxIdleTime  time.Duration // the maximum amount of time a connection may be idle.
}

// Open initialize db connection
func Open(c Config) (*gorm.DB, error) {
	dbConn, err := gorm.Open(mysql.Open(c.DSN), &c.Config)
	if err != nil {
		return nil, err
	}

	sqlDB, _ := dbConn.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(c.MaxLifetime)
	sqlDB.SetConnMaxIdleTime(c.MaxIdleTime)
	return dbConn, nil
}
