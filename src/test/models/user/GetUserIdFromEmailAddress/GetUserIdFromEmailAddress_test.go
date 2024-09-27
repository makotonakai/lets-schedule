package test 

import (
	"fmt"
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestGetUserIdFromEmailAddressSuccess(t *testing.T){
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

	emailAddress := "user@email.com"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT users.id FROM `users` WHERE users.email_address = ?")).
		WithArgs(emailAddress).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow(1)) 

	result, err := models.GetUserIdFromEmailAddress(gormDB, emailAddress)
	expected := 1

	if result != expected {
		t.Errorf("expected to get user id %d, but you actually got %d", result, expected)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestGetUserIdFromEmailAddressFail(t *testing.T){
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

	emailAddress := "user@email.com"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT users.id FROM `users` WHERE users.email_address = ?")).
	WillReturnError(fmt.Errorf("user with email address %s not found", emailAddress))


	result, err := models.GetUserIdFromEmailAddress(gormDB, emailAddress)
	expected := -1

	if result != expected {
		t.Errorf("expected to get user id %d, but you actually got %d", result, expected)
	}

	if err == nil {
		t.Errorf("expected to get an error")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

