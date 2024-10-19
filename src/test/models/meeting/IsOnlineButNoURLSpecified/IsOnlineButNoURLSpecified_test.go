package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsOnlineButNoURLSpecifiedFirst(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Url = ""
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedSecond(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = true
	m.Url = "hoge"
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedThird(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Url = ""
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedForth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = true
	m.IsOnline = false
	m.Url = "hoge"
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedFifth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Url = ""
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedSixth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = true
	m.Url = "hoge"
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedSeventh(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Url = ""
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedEighth(t *testing.T){

	m := models.Meeting{}
	m.IsOnsite = false
	m.IsOnline = false
	m.Url = "hoge"
	result, err := models.IsOnlineButNoURLSpecified(&m)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsOnlineButNoURLSpecifiedNil(t *testing.T){

	result, err := models.IsOnlineButNoURLSpecified(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrMeetingDoesNotExist) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrMeetingDoesNotExist.Error())
	}
}

