package example

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB inits a global db instance.
func InitDB(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Connect to mysql failed with error: %s", err.Error()))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// UserProfile user profile table
type UserProfile struct {
	ID       uint64 `gorm:"column:id;primaryKey:autoIncrement"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
}

// TableName returns the table name.
func (UserProfile) TableName() string {
	return "user_profile_tab"
}

// InsertUserProfile insert a new record.
func InsertUserProfile(profile *UserProfile) error {
	return db.Create(profile).Error
}

// GetUserProfiles get all user profiles.
func GetUserProfiles() ([]*UserProfile, error) {
	var profiles []*UserProfile

	err := db.Where("id > ?", 0).Order("id asc").Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil
}
