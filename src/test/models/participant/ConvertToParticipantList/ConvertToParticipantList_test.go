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

func TestConvertToParticipantListSuccess(t *testing.T){

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

	pl := []models.ParticipantWithUserName{
		{
			UserName: "user",
			MeetingId: 1,
			IsHost: true,
			HasResponded: true,
		},
	}

	presult, err := models.ConvertToParticipantList(gormDB, &pl)
	pexpected := &[]models.Participant{
		models.Participant{
			UserId: 1,
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

func TestConvertToParticipantListEmpty(t *testing.T){

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

	pl := []models.ParticipantWithUserName{}

	presult, err := models.ConvertToParticipantList(gormDB, &pl)
	pexpected := &[]models.Participant{}

	result := *presult
	expected := *pexpected

	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if !errors.Is(err, config.ErrParticipantListIsEmpty) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantListIsEmpty.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestConvertToParticipantListNil(t *testing.T){

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

	presult, err := models.ConvertToParticipantList(gormDB, nil)
	pexpected := &[]models.Participant{}

	result := *presult
	expected := *pexpected

	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if !errors.Is(err, config.ErrParticipantListIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantListIsNil.Error())
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
