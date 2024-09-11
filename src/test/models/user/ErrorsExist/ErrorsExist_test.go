package test 

import (
	"testing"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestErrorsExistTrue(t *testing.T){
	errors := []string{""}
	result := models.ErrorsExist(errors)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestErrorsExistFalse(t *testing.T){
	errors := []string{}
	result := models.ErrorsExist(errors)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}
