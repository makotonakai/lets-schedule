package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsOnsiteButNoPlaceSpecifiedFirst(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = ""
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedSecond(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = "hoge"
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedThird(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = ""
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedForth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = "hoge"
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedFifth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = ""
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedSixth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = "hoge"
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedSeventh(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = ""
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedEighth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = "hoge"
	result, err := models.IsOnsiteButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnsiteButNoPlaceSpecifiedNil(t *testing.T){

	result, err := models.IsOnsiteButNoPlaceSpecified(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingNotFound.Error())
	}
}

