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


func TestGetCandidateTimeByMeetingIdAndUserIdSuccess(t *testing.T){

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
	startTime := time.Date(2022, time.September, 19, 19, 0, 0, 0, time.UTC)
	endTime := time.Date(2022, time.September, 19, 20, 0, 0, 0, time.UTC)
	createdAt := time.Date(2024, time.September, 19, 19, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.September, 19, 20, 0, 0, 0, time.UTC)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `candidate_times` WHERE candidate_times.meeting_id = ? AND candidate_times.user_id = ?")).
		WithArgs(meetingId, userId).
		WillReturnRows(sqlmock.NewRows([]string{"id",  "meeting_id", "user_id", "start_time", "end_time", "created_at", "updated_at"}).
			AddRow(id, meetingId, userId, startTime,  endTime, createdAt,  updatedAt))

	result, err := models.GetCandidateTimeByMeetingIdAndUserId(gormDB, meetingId, userId)
	expected := []models.CandidateTime{
		{
			Id: id,
			MeetingId: meetingId,
			UserId: userId,
			StartTime: startTime,
			EndTime: endTime,
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

func TestGetCandidateTimeByMeetingIdAndUserIdWithNonExistingMeetingId(t *testing.T){

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
	userId := 1

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `candidate_times` WHERE candidate_times.meeting_id = ? AND candidate_times.user_id = ?")).
		WithArgs(meetingId, userId).
		WillReturnError(fmt.Errorf("record not found"))

	result, err := models.GetCandidateTimeByMeetingIdAndUserId(gormDB, meetingId, userId)
	expected := []models.CandidateTime{}
		
	// Verify the result
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

func TestGetCandidateTimeByMeetingIdAndUserIdWithNonExistingUserId(t *testing.T){

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

	meetingId := 1
	userId := 100

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `candidate_times` WHERE candidate_times.meeting_id = ? AND candidate_times.user_id = ?")).
		WithArgs(meetingId, userId).
		WillReturnError(fmt.Errorf("record not found"))

	result, err := models.GetCandidateTimeByMeetingIdAndUserId(gormDB, meetingId, userId)
	expected := []models.CandidateTime{}
		
	// Verify the result
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
