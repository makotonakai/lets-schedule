package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestIsMinSuccess(t *testing.T){

	a := 1
	b := 2

	presult, err := models.Min(&a, &b)
	expected := 1

	result := *presult
	if result != expected {
			t.Errorf("got %v, wanted %v", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsMinSame(t *testing.T){

	a := 1
	b := 1

	presult, err := models.Min(&a, &b)
	expected := 1

	result := *presult
	if result != expected {
			t.Errorf("got %v, wanted %v", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestIsMinSameNil(t *testing.T){

	a := 1

	presult, err := models.Min(&a, nil)
	expected := 0

	result := *presult
	if result != expected {
			t.Errorf("got %v, wanted nil", result)
	}

	if !errors.Is(err, config.ErrIntegerIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrIntegerIsNil.Error())
	}
}


