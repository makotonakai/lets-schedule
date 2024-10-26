package test 

import (
	"time"
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsAvailableTimeEmptySuccess(t *testing.T){
	at := models.AvailableTime{
		ActualStartTime: time.Time{},
		ActualEndTime: time.Time{},
	}
	result, err := models.IsAvailableTimeEmpty(&at)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsAvailableTimeEmptyWithActualStartTime(t *testing.T){
	at := models.AvailableTime{
		ActualStartTime: time.Now(),
		ActualEndTime: time.Time{},
	}
	result, err := models.IsAvailableTimeEmpty(&at)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsAvailableTimeEmptyNil(t *testing.T){
	
	result, err := models.IsAvailableTimeEmpty(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrAvailableTimeIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrAvailableTimeIsNil.Error())
	}
}
