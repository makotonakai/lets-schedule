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

func TestConvertToParticipantSuccess(t *testing.T){

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

	pw := models.ParticipantWithUserName{
		UserName: "user",
		MeetingId: 1,
		IsHost: true,
		HasResponded: true,
	}

	result, err := models.ConvertToParticipant(gormDB, &pw)
	expected := &models.Participant{
		UserId: 1,
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

func TestConvertToParticipantNil(t *testing.T){

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

	result, err := models.ConvertToParticipant(gormDB, nil)
	expected := &models.Participant{}

	if *result != *expected {
			t.Errorf("got %v, wanted %v", result, expected)
	}

	if !errors.Is(err, config.ErrParticipantWithUserNameIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantWithUserNameIsNil.Error())
	}

}
