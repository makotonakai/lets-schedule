package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsPasswordEmptyOrNullEmpty(t *testing.T){
	u := models.User{}
	u.Password = ""
	result, err := models.IsPasswordEmptyOrNull(&u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsPasswordEmptyOrNullSpace(t *testing.T){
	u := models.User{}
	u.Password = " "
	result, err := models.IsPasswordEmptyOrNull(&u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsPasswordEmptyOrNullValid(t *testing.T){
	u := models.User{}
	u.Password = "password"
	result, err := models.IsPasswordEmptyOrNull(&u)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsPasswordEmptyOrNullNil(t *testing.T){

	result, err := models.IsPasswordEmptyOrNull(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrUserIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrUserIsNil.Error())
	}
}
