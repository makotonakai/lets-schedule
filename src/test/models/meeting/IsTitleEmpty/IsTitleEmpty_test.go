package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsTitleEmptySuccess(t *testing.T){
	m := models.Meeting{}
	m.Title = ""
	result, err := models.IsTitleEmpty(&m)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsTitleEmptyFail(t *testing.T){
	m := models.Meeting{}
	m.Title = "hoge"
	result, err := models.IsTitleEmpty(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsTitleEmptyNil(t *testing.T){
	result, err := models.IsTitleEmpty(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingDoesNotExist) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingDoesNotExist.Error())
	}

}
