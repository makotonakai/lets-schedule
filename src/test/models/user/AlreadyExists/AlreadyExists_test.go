package test 

import (
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestErrorsExistValid(t *testing.T){
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

	// Test data
	testUser := models.User{
		UserName:     "user",
		EmailAddress: "user@email.com",
	}

	// Set up the expected SQL for email check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.email_address = ?")).
		WithArgs(testUser.EmailAddress).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email_address"}).
			AddRow(1, "user", "user@email.com")) // Simulate email already exists

	// Set up the expected SQL for username check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE users.user_name = ?")).
		WithArgs(testUser.UserName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_name", "email_address"}).
		AddRow(1, "user", "user@email.com")) // Simulate username does not exist

	// Call the AlreadyExists function
	result := models.AlreadyExists(gormDB, testUser)
	expected := true

	// Verify the result
	if result != expected {
		t.Errorf("expected user to exist, but AlreadyExists returned false")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}
