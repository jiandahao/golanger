// Code generated.

package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jiandahao/golanger/pkg/storage/cache"
	"github.com/jiandahao/golanger/pkg/storage/dbutils"
	"gorm.io/gorm"
)

var _ = sql.NullString{}

var (
	cacheTestProjectUserTabIdPrefix       = "cache:testProject:userTab:id:"
	cacheTestProjectUserTabEmailPrefix    = "cache:testProject:userTab:email:"
	cacheTestProjectUserTabUsernamePrefix = "cache:testProject:userTab:username:"
)

type (
	// UserTabModel is an interface that wraps the CURD method.
	UserTabModel interface {
		Insert(ctx context.Context, data *UserTab) error
		FindOne(ctx context.Context, id int64) (*UserTab, error)
		FindOneByEmail(ctx context.Context, email string) (*UserTab, error)
		FindOneByUsername(ctx context.Context, username string) (*UserTab, error)
		Update(ctx context.Context, data *UserTab) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserTabModel struct {
		dbConn     *gorm.DB
		cachedConn cache.CachedConn
	}

	// UserTab describes the table schema structure.
	UserTab struct {
		Id       int64  `gorm:"column:id"`
		Username string `gorm:"column:username"` // the username
		Password string `gorm:"column:password"`
		Nickname string `gorm:"column:nickname"` // nickname
		Email    string `gorm:"column:email"`    // user email address
		Avatar   string `gorm:"column:avatar"`
	}
)

// TableName returns the table name.
func (UserTab) TableName() string {
	return "`user_tab`"
}

// NewUserTabModel creates a defaultUserTabModel.
func NewUserTabModel(conn *gorm.DB, cacheConn cache.CachedConn) UserTabModel {
	return &defaultUserTabModel{
		dbConn:     conn,
		cachedConn: cacheConn,
	}
}

// Insert insert one record into user_tab.
func (m *defaultUserTabModel) Insert(ctx context.Context, data *UserTab) error {
	err := dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(&data).Error
	})

	if err != nil {
		return err
	}

	return nil
}

// FindOne find records by primary key.
func (m *defaultUserTabModel) FindOne(ctx context.Context, id int64) (*UserTab, error) {
	var resp UserTab
	testProjectUserTabIdKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabIdPrefix, id)
	err := m.cachedConn.QueryRow(&resp, func(v interface{}) error {
		return m.dbConn.WithContext(ctx).Where("`id`  = ?", id).Limit(1).Take(v).Error
	}, testProjectUserTabIdKey)

	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

// FindOneByEmail finds records by Email.
func (m *defaultUserTabModel) FindOneByEmail(ctx context.Context, email string) (*UserTab, error) {
	var resp UserTab
	testProjectUserTabEmailKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabEmailPrefix, email)
	err := m.cachedConn.QueryRow(&resp, func(v interface{}) error {
		return m.dbConn.WithContext(ctx).Where("`email` = ?", email).Limit(1).Take(v).Error
	}, testProjectUserTabEmailKey)

	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

// FindOneByUsername finds records by Username.
func (m *defaultUserTabModel) FindOneByUsername(ctx context.Context, username string) (*UserTab, error) {
	var resp UserTab
	testProjectUserTabUsernameKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabUsernamePrefix, username)
	err := m.cachedConn.QueryRow(&resp, func(v interface{}) error {
		return m.dbConn.WithContext(ctx).Where("`username` = ?", username).Limit(1).Take(v).Error
	}, testProjectUserTabUsernameKey)

	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

// Update update a record.
func (m *defaultUserTabModel) Update(ctx context.Context, data *UserTab) error {
	testProjectUserTabIdKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabIdPrefix, data.Id)
	testProjectUserTabEmailKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabEmailPrefix, data.Email)
	testProjectUserTabUsernameKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabUsernamePrefix, data.Username)
	keys := []string{testProjectUserTabIdKey, testProjectUserTabEmailKey, testProjectUserTabUsernameKey}

	return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
		return tx.WithContext(ctx).Updates(data).Error
	}, func() {
		m.cachedConn.DelCache(keys...)
	})
}

// Delete deletes by primary key.
func (m *defaultUserTabModel) Delete(ctx context.Context, id int64) error {

	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	testProjectUserTabIdKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabIdPrefix, id)
	testProjectUserTabEmailKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabEmailPrefix, data.Email)
	testProjectUserTabUsernameKey := fmt.Sprintf("%s%v", cacheTestProjectUserTabUsernamePrefix, data.Username)

	keys := []string{testProjectUserTabIdKey, testProjectUserTabEmailKey, testProjectUserTabUsernameKey}
	return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
		return tx.WithContext(ctx).Exec(fmt.Sprintf("DELETE FROM %s WHERE `id` = ? LIMIT 1", UserTab{}.TableName()), id).Error
	}, func() {
		m.cachedConn.DelCache(keys...)
	})
}
