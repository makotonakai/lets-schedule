package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsHourEmptySuccess(t *testing.T){

	m := models.Meeting{}
	m.Hour = 0
	result, err := models.IsHourEmpty(&m)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHourEmptyFail(t *testing.T){

	m := models.Meeting{}
	m.Hour = 1
	result, err := models.IsHourEmpty(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHourEmptyNegative(t *testing.T){

	m := models.Meeting{}
	m.Hour = -1
	result, err := models.IsHourEmpty(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingHourIsNegative) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingHourIsNegative.Error())
	}
}

func TestIsHourEmptyNil(t *testing.T){

	result, err := models.IsHourEmpty(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingDoesNotExist) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingDoesNotExist.Error())
	}
}


