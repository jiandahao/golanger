package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	utils "github.com/jiandahao/golanger/pkg/utils/db"
	"gorm.io/gorm"
)

type (
	// UserTabModel is an interface that wraps the CURD method.
	UserTabModel interface {
		Insert(ctx context.Context, data *UserTab) error
		FindOne(ctx context.Context, id int64) (*UserTab, error)
		FindOneByEmail(ctx context.Context, email sql.NullString) (*UserTab, error)
		FindOneByUsername(ctx context.Context, username string) (*UserTab, error)
		Update(ctx context.Context, data *UserTab) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserTabModel struct {
		dbConn *gorm.DB
	}

	// UserTab describes the table schema structure.
	UserTab struct {
		Id       int64          `gorm:"column:id"`
		Username string         `gorm:"column:username"` // 用户名
		Password string         `gorm:"column:password"`
		Nickname sql.NullString `gorm:"column:nickname"` // 用户昵称
		Email    sql.NullString `gorm:"column:email"`    // 用户邮箱
		Avatar   sql.NullString `gorm:"column:avatar"`
	}
)

// TableName returns the table name.
func (UserTab) TableName() string {
	return "`user_tab`"
}

// NewUserTabModel creates a defaultUserTabModel.
func NewUserTabModel(conn *gorm.DB) UserTabModel {
	return &defaultUserTabModel{
		dbConn: conn,
		// table:      "`user_tab`",
	}
}

// Insert insert one record into user_tab.
func (m *defaultUserTabModel) Insert(ctx context.Context, data *UserTab) error {
	err := utils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
		return tx.Create(&data).Error
	})

	if err != nil {
		return err
	}

	return nil
}

// FindOne find records by primary key.
func (m *defaultUserTabModel) FindOne(ctx context.Context, id int64) (*UserTab, error) {
	var resp UserTab
	err := m.dbConn.Where("`id` = ?", id).Limit(1).Find(&resp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// FindOneBy find records by Email.
func (m *defaultUserTabModel) FindOneByEmail(ctx context.Context, email sql.NullString) (*UserTab, error) {
	var resp UserTab
	err := m.dbConn.Where("`email` = ?", email).Limit(1).Take(&resp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// FindOneBy find records by Username.
func (m *defaultUserTabModel) FindOneByUsername(ctx context.Context, username string) (*UserTab, error) {
	var resp UserTab
	err := m.dbConn.Where("`username` = ?", username).Limit(1).Take(&resp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Update update a record.
func (m *defaultUserTabModel) Update(ctx context.Context, data *UserTab) error {
	return utils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
		return tx.Updates(data).Error
	})
}

// Delete deletes by primary key.
func (m *defaultUserTabModel) Delete(ctx context.Context, id int64) error {
	return utils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
		return tx.Exec(fmt.Sprintf("DELETE FROM %s WHERE `id` = ? LIMIT 1", UserTab{}.TableName()), id).Error
	})
}
