package test 

import (
	"errors"
	"testing"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsEmailAddressValidEmpty(t *testing.T){
	emailAddress := ""
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if !errors.Is(err, config.ErrEmailAddressIsEmpty) {
		t.Errorf("got %s, wanted %s", err.Error(), config.ErrEmailAddressIsEmpty.Error())
	}
}

func TestIsEmailAddressValidTest(t *testing.T){
	emailAddress := "test"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidAt(t *testing.T){
	emailAddress := "@"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidEmail(t *testing.T){
	emailAddress := "email"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidDot(t *testing.T){
	emailAddress := "."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidCom(t *testing.T){
	emailAddress := "com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestAt(t *testing.T){
	emailAddress := "test@"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestEmail(t *testing.T){
	emailAddress := "testemail"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestDot(t *testing.T){
	emailAddress := "test."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestCom(t *testing.T){
	emailAddress := "testcom"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidAtEmail(t *testing.T){
	emailAddress := "@email"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidAtDot(t *testing.T){
	emailAddress := "@."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidAtCom(t *testing.T){
	emailAddress := "@com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidEmailDot(t *testing.T){
	emailAddress := "email."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidEmailCom(t *testing.T){
	emailAddress := "emailcom"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidDotCom(t *testing.T){
	emailAddress := ".com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestAtEmail(t *testing.T){
	emailAddress := "test@email"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestAtDot(t *testing.T){
	emailAddress := "test@."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestAtCom(t *testing.T){
	emailAddress := "test@com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestEmailDot(t *testing.T){
	emailAddress := "testemail."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestEmailCom(t *testing.T){
	emailAddress := "testemailcom"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestDotCom(t *testing.T){
	emailAddress := "test.com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidAtEmailDot(t *testing.T){
	emailAddress := "@email."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidAtEmailCom(t *testing.T){
	emailAddress := "@emailcom"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidEmailDotCom(t *testing.T){
	emailAddress := "email.com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestAtEmailDot(t *testing.T){
	emailAddress := "test@email."
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestAtEmailCom(t *testing.T){
	emailAddress := "test@emailcom"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestAtDotCom(t *testing.T){
	emailAddress := "test@.com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidTestEmailDotCom(t *testing.T){
	emailAddress := "testemail.com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValidAtEmailDotCom(t *testing.T){
	emailAddress := "@email.com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}

func TestIsEmailAddressValid(t *testing.T){
	emailAddress := "test@email.com"
	result, err := models.IsEmailAddressValid(emailAddress)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}

	if err != nil {
		t.Errorf("expected to get no error")
	}
}






