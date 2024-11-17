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

func TestGetUserNameFromUserIdSuccess(t *testing.T){
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
	gormDB, e := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to initialize GORM DB: %v", err)
	}

	id := 1

	mock.ExpectQuery(regexp.QuoteMeta("SELECT users.user_name FROM `users` WHERE users.id = ?")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"user_name"}).
		AddRow("user")) 

	result, err := models.GetUserNameFromUserId(gormDB, id)
	expected := "user"

	if result != expected {
		t.Errorf("expected to get username %s, but you actually got %s", result, expected)
	}

	if e != nil {
		t.Errorf("got %t, wanted nil", e)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestGetUserNameFromUserIdFail(t *testing.T){
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

	id := -1

	mock.ExpectQuery(regexp.QuoteMeta("SELECT users.user_name FROM `users` WHERE users.id = ?")).
		WithArgs(id).
		WillReturnError(fmt.Errorf("user with id %d not found", id))

	result, e := models.GetUserNameFromUserId(gormDB, id)
	expected := ""

	if result != expected {
		t.Errorf("expected to get username %s, but you actually got %s", result, expected)
	}

	if !errors.Is(e, config.ErrUserWithUserIdNotFound) {
		t.Errorf("got %s, wanted %s", e.Error(), config.ErrUserWithUserIdNotFound.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

