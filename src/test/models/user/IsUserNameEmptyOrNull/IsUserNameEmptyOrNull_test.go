package test 

import (
	"testing"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsUserNameEmptyOrNullEmpty(t *testing.T){
	u := models.User{}
	u.UserName = ""
	result := models.IsUserNameEmptyOrNull(u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsUserNameEmptyOrNullSpace(t *testing.T){
	u := models.User{}
	u.UserName = " "
	result := models.IsUserNameEmptyOrNull(u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsUserNameEmptyOrNullValid(t *testing.T){
	u := models.User{}
	u.UserName = "user"
	result := models.IsUserNameEmptyOrNull(u)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}
