package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsCandidateTimeEmptySuccess(t *testing.T){
	ctlist := []models.CandidateTime{}
	result, err := models.IsCandidateTimeEmpty(&ctlist)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsCandidateTimeEmptyFail(t *testing.T){
	ctlist := []models.CandidateTime{
    models.CandidateTime{},
  }
	result, err := models.IsCandidateTimeEmpty(&ctlist)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsCandidateTimeEmptyNil(t *testing.T){

  result, err := models.IsCandidateTimeEmpty(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrArrayIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrArrayIsNil.Error())
	}
}
