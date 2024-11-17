package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestErrorsExistTrue(t *testing.T){
	errors := []string{""}
	result, err := models.ErrorsExist(&errors)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestErrorsExistFalse(t *testing.T){
	errors := []string{}
	result, err := models.ErrorsExist(&errors)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestErrorsExistNil(t *testing.T){
	result, err := models.ErrorsExist(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrListOfErrorsNotFound) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrListOfErrorsNotFound.Error())
	}
}
