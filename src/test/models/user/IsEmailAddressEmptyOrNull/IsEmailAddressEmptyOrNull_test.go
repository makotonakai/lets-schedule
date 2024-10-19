package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)



func TestIsEmailAddressEmptyOrNullEmpty(t *testing.T){
	u := models.User{}
	u.EmailAddress = ""
	result, err := models.IsEmailAddressEmptyOrNull(&u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsEmailAddressEmptyOrNullSpace(t *testing.T){
	u := models.User{}
	u.EmailAddress = " "
	result, err := models.IsEmailAddressEmptyOrNull(&u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsEmailAddressEmptyOrNullValid(t *testing.T){
	u := models.User{}
	u.EmailAddress = "hoge"
	result, err := models.IsEmailAddressEmptyOrNull(&u)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsEmailAddressEmptyOrNullNil(t *testing.T){
	result, err := models.IsEmailAddressEmptyOrNull(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrUserIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrUserIsNil.Error())
	}
}
