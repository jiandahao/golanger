package model

import (
	"context"
	"fmt"
	"github.com/jiandahao/golanger/pkg/storage/filter"
	"github.com/jiandahao/golanger/pkg/storage/options"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type (
	// UserTabModel is an interface that wraps the CURD method.
	UserTabModel interface {
		Insert(ctx context.Context, data *UserTab) error
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, data *UserTab) error
		FindOne(ctx context.Context, id int64) (*UserTab, error)
		FindOneByEmail(ctx context.Context, email string) (*UserTab, error)
		FindOneByUsername(ctx context.Context, username string) (*UserTab, error)
		Query(ctx context.Context, filters string, opts ...options.QueryOption) ([]*UserTab, int64, error)
		WithDB(db *gorm.DB) UserTabModel
	}

	defaultUserTabModel struct {
		dbConn *gorm.DB
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
func NewUserTabModel(conn *gorm.DB) UserTabModel {
	return &defaultUserTabModel{
		dbConn: conn,
	}
}

func (m *defaultUserTabModel) WithDB(dbConn *gorm.DB) UserTabModel {
	return &defaultUserTabModel{
		dbConn: dbConn,
	}
}

// Insert insert one record into user_tab.
func (m *defaultUserTabModel) Insert(ctx context.Context, data *UserTab) error {
	err := m.dbConn.WithContext(ctx).Create(data).Error
	if err != nil {
		return errors.Wrap(err, "Insert error")
	}

	return nil
}

// Delete delete a record by primary key.
func (m *defaultUserTabModel) Delete(ctx context.Context, id int64) error {
	if err := m.dbConn.WithContext(ctx).Exec(fmt.Sprintf("DELETE FROM %s WHERE `id` = ? LIMIT 1", UserTab{}.TableName()), id).Error; err != nil {
		return errors.Wrap(err, "Delete error")
	}

	return nil
}

// Update update a record.
func (m *defaultUserTabModel) Update(ctx context.Context, data *UserTab) error {
	err := m.dbConn.WithContext(ctx).Where("`id`  = ?", data.Id).Updates(data).Error
	if err != nil {
		return errors.Wrap(err, "Update error")
	}
	return nil
}

// FindOne find records by primary key.
func (m *defaultUserTabModel) FindOne(ctx context.Context, id int64) (*UserTab, error) {
	var resp UserTab
	err := m.dbConn.WithContext(ctx).Where("`id`  = ?", id).Limit(1).Take(&resp).Error
	if err != nil {
		return nil, errors.Wrap(err, "FindOne error")
	}

	return &resp, nil
}

// FindOneByEmail find one record by unique key email.
func (m *defaultUserTabModel) FindOneByEmail(ctx context.Context, email string) (*UserTab, error) {
	var resp UserTab
	err := m.dbConn.WithContext(ctx).Where("`email`  = ?", email).Limit(1).Take(&resp).Error
	if err != nil {
		return nil, errors.Wrap(err, "FindOneByEmail error")
	}

	return &resp, nil
}

// FindOneByUsername find one record by unique key username.
func (m *defaultUserTabModel) FindOneByUsername(ctx context.Context, username string) (*UserTab, error) {
	var resp UserTab
	err := m.dbConn.WithContext(ctx).Where("`username`  = ?", username).Limit(1).Take(&resp).Error
	if err != nil {
		return nil, errors.Wrap(err, "FindOneByUsername error")
	}

	return &resp, nil
}

const (
	UserTabColumn_Id       string = "id"
	UserTabColumn_Username string = "username"
	UserTabColumn_Password string = "password"
	UserTabColumn_Nickname string = "nickname"
	UserTabColumn_Email    string = "email"
	UserTabColumn_Avatar   string = "avatar"
)

// define all permitted query conditions
var userTabQueryFilter = filter.NewParser(map[filter.FieldNameType][]filter.Operator{
	"id":       {filter.Equal, filter.In},
	"email":    {filter.Equal, filter.In},
	"username": {filter.Equal, filter.In},
})

// Query query records by filters.
func (m *defaultUserTabModel) Query(ctx context.Context, filters string, opts ...options.QueryOption) ([]*UserTab, int64, error) {
	dbClient := m.dbConn.WithContext(ctx)
	if filters != "" {
		conds, args, err := userTabQueryFilter.Parse(filters)
		if err != nil {
			return nil, 0, errors.Wrap(ErrInvalidArgument, err.Error())
		}

		dbClient = dbClient.Where(conds, args...)
	}

	var counter int64
	if err := dbClient.Model(&UserTab{}).Count(&counter).Error; err != nil {
		return nil, 0, errors.Wrap(err, "Query:Count error")
	}

	for _, opt := range opts {
		dbClient = opt(dbClient)
	}

	var records []*UserTab
	if err := dbClient.Find(&records).Error; err != nil {
		return nil, 0, errors.Wrap(err, "Query:Find error")
	}

	return records, counter, nil
}
