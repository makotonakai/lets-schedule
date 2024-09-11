package test 

import (
	"testing"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsPasswordEmptyOrNullEmpty(t *testing.T){
	u := models.User{}
	u.Password = ""
	result := models.IsPasswordEmptyOrNull(u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsPasswordEmptyOrNullSpace(t *testing.T){
	u := models.User{}
	u.Password = " "
	result := models.IsPasswordEmptyOrNull(u)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsPasswordEmptyOrNullValid(t *testing.T){
	u := models.User{}
	u.Password = "password"
	result := models.IsPasswordEmptyOrNull(u)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}
