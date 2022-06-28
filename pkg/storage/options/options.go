package options

import (
	"gorm.io/gorm"
)

type QueryOption func(db *gorm.DB) *gorm.DB

// Paging pagination query
func Paging(pageNum int, pageSize int) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum < 0 || pageSize < 0 {
			return db
		}

		return db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
}

// OrderBy specify order when retrieve records from database
func OrderBy(value interface{}) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(value)
	}
}

// Omit specify fields that you want to ignore when querying
func Omit(columns ...string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(columns...)
	}
}

// Where add where conditions
func Where(query string, args ...interface{}) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}
