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

func TestConvertToParticipantWithUserNameListSuccess(t *testing.T){

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

	pl := []models.Participant{
		{
			UserId: 1,
			MeetingId: 1,
			IsHost: true,
			HasResponded: true,
		},
	}

	presult, err := models.ConvertToParticipantWithUserNameList(gormDB, &pl)
	pexpected := &[]models.ParticipantWithUserName{
		models.ParticipantWithUserName{
			UserName: "user",
			MeetingId: 1,
			IsHost: true,
			HasResponded: true,
		},
	}

	result := *presult
	expected := *pexpected

	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if result[0] != expected[0] {
		t.Errorf("got %v, wanted %v", result, expected)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestConvertToParticipantWithUserNameListEmpty(t *testing.T){

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

	pl := []models.Participant{}

	presult, err := models.ConvertToParticipantWithUserNameList(gormDB, &pl)
	pexpected := &[]models.ParticipantWithUserName{}

	result := *presult
	expected := *pexpected

	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if !errors.Is(err, config.ErrParticipantWithUserNameListIsEmpty) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantWithUserNameListIsEmpty.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestConvertToParticipantWithUserNameListNil(t *testing.T){

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

	presult, err := models.ConvertToParticipantWithUserNameList(gormDB, nil)
	pexpected := &[]models.ParticipantWithUserName{}

	result := *presult
	expected := *pexpected

	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if !errors.Is(err, config.ErrParticipantWithUserNameListIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantWithUserNameListIsNil.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
