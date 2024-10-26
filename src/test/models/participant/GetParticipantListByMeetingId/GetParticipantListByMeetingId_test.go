package test 

import (
	"fmt"
	"time"
	"errors"
	"regexp"
	"testing"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestGetParticipantListByMeetingIdSuccess(t *testing.T){

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

	id := 1
	meetingId := 1
	userId := 1
	isHost := true
	hasResponded := true
	createdAt := time.Date(2024, time.September, 19, 19, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.September, 19, 20, 0, 0, 0, time.UTC)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `participants` WHERE participants.meeting_id = ?")).
		WithArgs(meetingId).
		WillReturnRows(sqlmock.NewRows([]string{"id",  "meeting_id", "user_id", "is_host", "has_responded", "created_at", "updated_at"}).
			AddRow(id, meetingId, userId, isHost, hasResponded, createdAt,  updatedAt))

	result, err := models.GetParticipantListByMeetingId(gormDB, meetingId)
	expected := []models.Participant{
		{
			Id: id,
			MeetingId: meetingId,
			UserId: userId,
			IsHost: isHost,
			HasResponded: hasResponded,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}
		
	// Verify the result

	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if result[0] != expected[0] {
		t.Errorf("got %v, wanted %v", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}

func TestGetParticipantListByMeetingIdFail(t *testing.T){

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

	meetingId := 100

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `participants` WHERE participants.meeting_id = ?")).
		WithArgs(meetingId).
		WillReturnError(fmt.Errorf("record not found"))

	result, err := models.GetParticipantListByMeetingId(gormDB, meetingId)
	expected := []models.Participant{}

	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if !errors.Is(err, config.ErrRecordNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrRecordNotFound.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}

}
