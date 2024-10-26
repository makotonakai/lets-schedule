package test 

import (
	"time"
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestPastCandidateTimeExistsWithStartTime(t *testing.T){
	ctlist := []models.CandidateTime{
		models.CandidateTime{
			StartTime: time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
			EndTime: time.Time{},
			},
		}
	result, err := models.PastCandidateTimeExists(&ctlist)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestPastCandidateTimeExistsWithEndTime(t *testing.T){
	ctlist := []models.CandidateTime{
		models.CandidateTime{
			StartTime: time.Time{},
			EndTime: time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
			},
		}
	result, err := models.PastCandidateTimeExists(&ctlist)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestPastCandidateTimeExistsWithFuture(t *testing.T){
	ctlist := []models.CandidateTime{
		models.CandidateTime{
			StartTime: time.Date(2025, 11, 17, 20, 34, 58, 651387237, time.UTC),
			EndTime: time.Date(2025, 11, 17, 20, 34, 58, 651387237, time.UTC),
		},
	}
	result, err := models.PastCandidateTimeExists(&ctlist)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestPastCandidateTimeExistsWithEmptyList(t *testing.T){
	
	ctlist := []models.CandidateTime{}
	result, err := models.PastCandidateTimeExists(&ctlist)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrArrayIsEmpty) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrArrayIsEmpty.Error())
	}
}

func TestPastCandidateTimeExistsNil(t *testing.T){
	
	result, err := models.PastCandidateTimeExists(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrArrayIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrArrayIsNil.Error())
	}
}
