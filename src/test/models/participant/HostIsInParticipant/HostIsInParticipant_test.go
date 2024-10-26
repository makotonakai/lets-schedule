package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)


func TestHostIsParticipantTrue(t *testing.T){

	pl := []models.Participant{
		models.Participant{
			UserId: 1,
			MeetingId: 1,
			IsHost: true,
			HasResponded: true,
		},
	}

	result, err := models.HostIsInParticipant(&pl)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestHostIsParticipantFalse(t *testing.T){

	pl := []models.Participant{
		models.Participant{
			UserId: 1,
			MeetingId: 1,
			IsHost: false,
			HasResponded: true,
		},
	}

	result, err := models.HostIsInParticipant(&pl)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("got %t, wanted nil", err)
	}
}

func TestHostIsParticipantEmpty(t *testing.T){

	pl := []models.Participant{}

	result, err := models.HostIsInParticipant(&pl)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrParticipantListIsEmpty) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantListIsEmpty.Error())
	}
}

func TestHostIsParticipantNil(t *testing.T){

	result, err := models.HostIsInParticipant(nil)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrParticipantListIsNil) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrParticipantListIsNil.Error())
	}
}
