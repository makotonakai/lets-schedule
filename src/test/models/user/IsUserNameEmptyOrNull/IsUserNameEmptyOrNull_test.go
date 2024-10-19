package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsUserNameEmptyOrNullEmpty(t *testing.T){
	u := models.User{}
	u.UserName = ""
	result, err := models.IsUserNameEmptyOrNull(&u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsUserNameEmptyOrNullSpace(t *testing.T){
	u := models.User{}
	u.UserName = " "
	result, err := models.IsUserNameEmptyOrNull(&u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsUserNameEmptyOrNullValid(t *testing.T){
	u := models.User{}
	u.UserName = "user"
	result, err := models.IsUserNameEmptyOrNull(&u)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsUserNameEmptyOrNullNil(t *testing.T){
	result, err := models.IsUserNameEmptyOrNull(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrUserIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrUserIsNil.Error())
	}
}


