package test 

import (
	"testing"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsEmailAddressEmptyOrNullEmpty(t *testing.T){
	u := models.User{}
	u.EmailAddress = ""
	result := models.IsEmailAddressEmptyOrNull(u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressEmptyOrNullSpace(t *testing.T){
	u := models.User{}
	u.EmailAddress = " "
	result := models.IsEmailAddressEmptyOrNull(u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressEmptyOrNullValid(t *testing.T){
	u := models.User{}
	u.EmailAddress = "test@email.com"
	result := models.IsEmailAddressEmptyOrNull(u)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

