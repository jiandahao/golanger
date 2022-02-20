package example

import (
	"fmt"
	"testing"

	"github.com/jiandahao/golanger/pkg/mysqltest"
)

var mysqltester *mysqltest.Mysqltester

func init() {
	var err error
	// create and lanuch a new mysql instance
	mysqltester, err = mysqltest.NewMysqltester()
	if err != nil {
		panic(err)
	}

	// create schema
	mysqltester.MustExec(`CREATE DATABASE if not exists test;`)

	// create a new table
	mysqltester.MustExec(`CREATE TABLE user_profile_tab (
		id int(11) NOT NULL AUTO_INCREMENT,
		username varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
		email varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
		PRIMARY KEY (id),
		UNIQUE KEY username_UNIQUE (username)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`)

	InitDB(mysqltester.DSN())
}

func TestGetUserProfiles(t *testing.T) {
	// insert records
	for i := 0; i < 10; i++ {
		mysqltester.MustExec(`insert into test.user_profile_tab (username, email) VALUES (?,?)`, fmt.Sprintf("user%v", i), fmt.Sprintf("user%v@example.com", i))
	}

	profiles, err := GetUserProfiles()
	if err != nil {
		t.Error(err)
		return
	}

	if len(profiles) != 10 {
		t.Errorf("TestGetUserProfiles() error = incorrect records number, want %v but got %v", 10, len(profiles))
		return
	}

	for index, profile := range profiles {
		if profile.Username != fmt.Sprintf("user%v", index) {
			t.Errorf("TestGetUserProfiles() error = result mismatched, want user %s but got %s", fmt.Sprintf("user%v", index), profile.Username)
		}
	}
}
