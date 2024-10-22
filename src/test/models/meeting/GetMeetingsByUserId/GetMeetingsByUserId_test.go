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


func TestGetMeetingsByUserIdFirst(t *testing.T){

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
	startTime := time.Date(2022, time.September, 19, 19, 0, 0, 0, time.UTC)
	endTime := time.Date(2022, time.September, 19, 20, 0, 0, 0, time.UTC)
	actualStartTime := time.Date(2022, time.September, 19, 19, 0, 0, 0, time.UTC)
	actualEndTime := time.Date(2022, time.September, 19, 20, 0, 0, 0, time.UTC)
	createdAt := time.Date(2024, time.September, 19, 19, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, time.September, 19, 20, 0, 0, 0, time.UTC)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `meetings` inner join participants on participants.meeting_id = meetings.id inner join users on users.id = participants.user_id WHERE users.id = ?")).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "is_onsite", "is_online", "place", "url", "all_participants_responded", "is_confirmed", "start_time", "end_time", "actual_start_time", "actual_end_time", "hour", "created_at", "updated_at"}).
			AddRow(1, "hoge", "hoge", true, true, "hoge", "zoom.com", true, true, startTime,  endTime, actualStartTime,  actualEndTime, 1, createdAt,  updatedAt))

	result, err := models.GetMeetingsByUserId(gormDB, userId)
	expected := []models.Meeting{
		{
			Id: 1,
			Title: "hoge",
			Description: "hoge",
			IsOnsite: true,
			IsOnline: true,
			Place: "hoge",
			Url: "zoom.com",
			AllParticipantsResponded: true,
			IsConfirmed: true,
			StartTime: startTime,
			EndTime: endTime,
			ActualStartTime: actualStartTime,
			ActualEndTime: actualEndTime,
			Hour: 1, 
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

}

func TestGetMeetingByIdNil(t *testing.T){

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

	userId := 100

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `meetings` inner join participants on participants.meeting_id = meetings.id inner join users on users.id = participants.user_id WHERE users.id = ?")).
		WithArgs(userId).
		WillReturnError(fmt.Errorf("record not found"))

	result, err := models.GetMeetingsByUserId(gormDB, userId)
	expected := []models.Meeting{}
	
	// Verify the result
	if len(result) != len(expected) {
		t.Errorf("The length of the actual array and the one for the expected array is different")
	}

	if !errors.Is(err, config.ErrRecordNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrRecordNotFound.Error())
	}

}
