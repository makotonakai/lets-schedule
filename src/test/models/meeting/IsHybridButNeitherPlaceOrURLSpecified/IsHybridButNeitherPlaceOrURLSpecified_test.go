package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsHybridButNeitherPlaceOrURLSpecifiedFirst(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = ""
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedSecond(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = ""
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedThird(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = "hoge"
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedForth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Place = "hoge"
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedFifth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = ""
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedSixth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = ""
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedSeventh(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = "hoge"
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedEighth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Place = "hoge"
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedNinth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = ""
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedTenth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = ""
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedEleventh(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = "hoge"
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedTwelf(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Place = "hoge"
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedThirteenth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = ""
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedFourteenth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = ""
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedFifteenth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = "hoge"
	m.Url = ""
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedSixteenth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Place = "hoge"
	m.Url = "zoom.com"
	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNeitherPlaceOrURLSpecifiedNil(t *testing.T){

	result, err := models.IsHybridButNeitherPlaceOrURLSpecified(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingNotFound.Error())
	}
}
