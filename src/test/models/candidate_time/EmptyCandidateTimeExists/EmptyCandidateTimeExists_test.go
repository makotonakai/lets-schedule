package test 

import (
	"time"
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestEmptyCandidateTimeExistsWithStartTime(t *testing.T){
	ctlist := []models.CandidateTime{
		models.CandidateTime{
			StartTime: time.Now(),
			EndTime: time.Time{},
			},
		}
	result, err := models.EmptyCandidateTimeExists(&ctlist)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestEmptyCandidateTimeExistsWithEndTime(t *testing.T){
	ctlist := []models.CandidateTime{
		models.CandidateTime{
			StartTime: time.Time{},
			EndTime: time.Now(),
			},
		}
	result, err := models.EmptyCandidateTimeExists(&ctlist)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestEmptyCandidateTimeExistsWithStartTimeAndEndTime(t *testing.T){
	ctlist := []models.CandidateTime{
		models.CandidateTime{
			StartTime: time.Now(),
			EndTime: time.Now(),
			},
		}
	result, err := models.EmptyCandidateTimeExists(&ctlist)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestEmptyCandidateTimeExistsWithEmptyList(t *testing.T){
	ctlist := []models.CandidateTime{}
	result, err := models.EmptyCandidateTimeExists(&ctlist)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrArrayIsEmpty) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrArrayIsEmpty.Error())
	}
}

func TestEmptyCandidateTimeExistsNil(t *testing.T){

	result, err := models.EmptyCandidateTimeExists(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrArrayIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrArrayIsNil.Error())
	}
}
