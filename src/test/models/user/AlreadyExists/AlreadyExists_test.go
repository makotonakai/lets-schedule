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

func TestAlreadyExistsWithExistingUser(t *testing.T){
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
	emailAddress := "user@email.com"

	// Test data
	testUser := models.User{
		UserName:     userName,
		EmailAddress: emailAddress,
	}

	// Set up the expected SQL for email check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.email_address = ?")).
		WithArgs(emailAddress).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email_address"}).
			AddRow(1, "user", "user@email.com")) // Simulate email already exists

	// Set up the expected SQL for username check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.user_name = ?")).
		WithArgs(userName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email_address"}).
		AddRow(1, "user", "user@email.com")) // Simulate username does not exist

	// Call the AlreadyExists function
	result, emailAddressErr, userNameErr := models.AlreadyExists(gormDB, &testUser)
	expected := true

	// Verify the result
	if result != expected {
		t.Errorf("expected user to exist, but AlreadyExists returned false")
	}

	if emailAddressErr != nil {
		t.Errorf("got %t, wanted nil", emailAddressErr)
	}

	if userNameErr != nil {
		t.Errorf("got %t, wanted nil", userNameErr)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestAlreadyExistsEmptyWithExistingUserName(t *testing.T){
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
	emailAddress := "user@email.com"

	// Test data
	testUser := models.User{
		UserName:     userName,
		EmailAddress: emailAddress,
	}

	// Set up the expected SQL for email check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.email_address = ?")).
		WithArgs(emailAddress).
		WillReturnError(fmt.Errorf("Email address not found"))

	// Set up the expected SQL for username check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.user_name = ?")).
		WithArgs(userName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email_address"}).
		AddRow(1, "user", "user@email.com")) // Simulate username does not exist

	// Call the AlreadyExists function
	result, emailAddressErr, userNameErr := models.AlreadyExists(gormDB, &testUser)
	expected := true

	// Verify the result
	if result != expected {
		t.Errorf("expected user to exist, but AlreadyExists returned false")
	}

	if !errors.Is(emailAddressErr, config.ErrEmailAddressNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrEmailAddressNotFound.Error())
	}

	if userNameErr != nil {
		t.Errorf("got %t, wanted nil", userNameErr)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestAlreadyExistsEmptyWithExistingEmailAddress(t *testing.T){
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

	userName := ""
	emailAddress := "user@email.com"

	// Test data
	testUser := models.User{
		UserName:     userName,
		EmailAddress: emailAddress,
	}

	// Set up the expected SQL for email check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.email_address = ?")).
		WithArgs(emailAddress).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email_address"}).
		AddRow(1, "user", "user@email.com")) // Simulate username does not exist

	// Set up the expected SQL for username check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.user_name = ?")).
		WithArgs(userName).
		WillReturnError(fmt.Errorf("User name not found"))


	// Call the AlreadyExists function
	result, emailAddressErr, userNameErr := models.AlreadyExists(gormDB, &testUser)
	expected := true

	// Verify the result
	if result != expected {
		t.Errorf("expected user to exist, but AlreadyExists returned false")
	}

	if emailAddressErr != nil {
		t.Errorf("got %t, wanted nil", emailAddressErr)
	} 

	if !errors.Is(userNameErr, config.ErrUserNameNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrUserNameNotFound.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestAlreadyExistsEmptyWithNonExistingUser(t *testing.T){
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

	userName := ""
	emailAddress := ""

	// Test data
	testUser := models.User{
		UserName:     userName,
		EmailAddress: emailAddress,
	}

	// Set up the expected SQL for email check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.email_address = ?")).
		WithArgs(emailAddress).
		WillReturnError(fmt.Errorf("Email adderss not found"))

	// Set up the expected SQL for username check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.user_name = ?")).
		WithArgs(userName).
		WillReturnError(fmt.Errorf("User name not found"))


	// Call the AlreadyExists function
	result, emailAddressErr, userNameErr := models.AlreadyExists(gormDB, &testUser)
	expected := false

	// Verify the result
	if result != expected {
		t.Errorf("expected user to exist, but AlreadyExists returned false")
	}

	if !errors.Is(emailAddressErr, config.ErrEmailAddressNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrEmailAddressNotFound.Error())
	}

	if !errors.Is(userNameErr, config.ErrUserNameNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrUserNameNotFound.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestAlreadyExistsEmptyNil(t *testing.T){

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

	// Call the AlreadyExists function
	result, emailAddressErr, userNameErr := models.AlreadyExists(gormDB, nil)
	expected := false

	// Verify the result
	if result != expected {
		t.Errorf("expected user to exist, but AlreadyExists returned false")
	}

	if !errors.Is(emailAddressErr, config.ErrUserIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrUserIsNil.Error())
	}

	if userNameErr != nil {
		t.Errorf("got %t, wanted nil", userNameErr)
	}
	
	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}
