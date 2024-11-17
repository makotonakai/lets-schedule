package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsHybridButNoPlaceSpecifiedFirst(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = ""
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedSecond(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = "hoge"
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedThird(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = ""
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedForth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = "hoge"
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedFifth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = ""
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedSixth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = "hoge"
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedSeventh(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = ""
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedEighth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = "hoge"
	result, err := models.IsHybridButNoPlaceSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoPlaceSpecifiedNil(t *testing.T){

	result, err := models.IsHybridButNoPlaceSpecified(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingNotFound.Error())
	}
}


