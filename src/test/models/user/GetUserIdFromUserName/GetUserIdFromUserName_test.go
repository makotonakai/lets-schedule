package test 

import (
	"fmt"
	"errors"
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestGetUserIdFromUserNameSuccess(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	// Set up GORM to use the mock DB with the MySQL driver
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true, // Skip version check
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to initialize GORM DB: %v", err)
	}

	userName := "user"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT users.id FROM `users` WHERE users.user_name = ?")).
		WithArgs(userName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow(1)) 

	result, e := models.GetUserIdFromUserName(gormDB, userName)
	expected := 1

	if result != expected {
		t.Errorf("expected to get user id %d, but you actually got %d", result, expected)
	}

	if e != nil {
		t.Errorf("got %t, wanted nil", e)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestGetUserIdFromUserNameFail(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	// Set up GORM to use the mock DB with the MySQL driver
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true, // Skip version check
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to initialize GORM DB: %v", err)
	}

	userName := "hoge"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT users.id FROM `users` WHERE users.user_name = ?")).
		WithArgs(userName).
		WillReturnError(fmt.Errorf("user with username %s not found", userName))

	result, e := models.GetUserIdFromUserName(gormDB, userName)
	expected := -1

	if result != expected {
		t.Errorf("expected to get user id %d, but you actually got %d", result, expected)
	}

	if !errors.Is(e, config.ErrUserWithUserNameNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrUserWithUserNameNotFound.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

