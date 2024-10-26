package test 

import (
	"errors"
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestConvertToParticipantWithUserNameSuccess(t *testing.T){

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

	userId := 1

	mock.ExpectQuery(regexp.QuoteMeta("SELECT users.user_name FROM `users` WHERE users.id = ?")).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"user_name"}).
		AddRow("user"))

	p := models.Participant{
		UserId: 1,
		MeetingId: 1,
		IsHost: true,
		HasResponded: true,
	}

	result, err := models.ConvertToParticipantWithUserName(gormDB, &p)
	expected := &models.ParticipantWithUserName{
		UserName: "user",
		MeetingId: 1,
		IsHost: true,
		HasResponded: true,
	}

	if *result != *expected {
			t.Errorf("got %v, wanted %v", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestConvertToParticipantWithUserNameNil(t *testing.T){

	db, _, err := sqlmock.New()
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

	result, err := models.ConvertToParticipantWithUserName(gormDB, nil)
	expected := &models.ParticipantWithUserName{}

	if *result != *expected {
			t.Errorf("got %v, wanted %v", result, expected)
	}

	if !errors.Is(err, config.ErrParticipantIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantIsNil.Error())
	}

}
