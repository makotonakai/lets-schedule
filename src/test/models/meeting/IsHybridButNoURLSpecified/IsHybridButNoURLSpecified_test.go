package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsHybridButNoURLSpecifiedFirst(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Url = ""
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedSecond(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Url = "hoge"
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedThird(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Url = ""
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedForth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Url = "hoge"
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedFifth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Url = ""
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedSixth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Url = "hoge"
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedSeventh(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Url = ""
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedEighth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Url = "hoge"
	result, err := models.IsHybridButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsHybridButNoURLSpecifiedNil(t *testing.T){

	result, err := models.IsHybridButNoURLSpecified(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingNotFound.Error())
	}
}

