package test 

import (
	"testing"
	"github.com/MakotoNakai/lets-schedule/models"
)

func TestIsEmailAddressValidEmpty(t *testing.T){
	emailAddress := ""
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTest(t *testing.T){
	emailAddress := "test"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidAt(t *testing.T){
	emailAddress := "@"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidEmail(t *testing.T){
	emailAddress := "email"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidDot(t *testing.T){
	emailAddress := "."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidCom(t *testing.T){
	emailAddress := "com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestAt(t *testing.T){
	emailAddress := "test@"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestEmail(t *testing.T){
	emailAddress := "testemail"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestDot(t *testing.T){
	emailAddress := "test."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestCom(t *testing.T){
	emailAddress := "testcom"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidAtEmail(t *testing.T){
	emailAddress := "@email"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidAtDot(t *testing.T){
	emailAddress := "@."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidAtCom(t *testing.T){
	emailAddress := "@com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidEmailDot(t *testing.T){
	emailAddress := "email."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidEmailCom(t *testing.T){
	emailAddress := "emailcom"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidDotCom(t *testing.T){
	emailAddress := ".com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestAtEmail(t *testing.T){
	emailAddress := "test@email"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestAtDot(t *testing.T){
	emailAddress := "test@."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestAtCom(t *testing.T){
	emailAddress := "test@com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestEmailDot(t *testing.T){
	emailAddress := "testemail."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestEmailCom(t *testing.T){
	emailAddress := "testemailcom"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestDotCom(t *testing.T){
	emailAddress := "test.com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidAtEmailDot(t *testing.T){
	emailAddress := "@email."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidAtEmailCom(t *testing.T){
	emailAddress := "@emailcom"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidEmailDotCom(t *testing.T){
	emailAddress := "email.com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestAtEmailDot(t *testing.T){
	emailAddress := "test@email."
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestAtEmailCom(t *testing.T){
	emailAddress := "test@emailcom"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestAtDotCom(t *testing.T){
	emailAddress := "test@.com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidTestEmailDotCom(t *testing.T){
	emailAddress := "testemail.com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValidAtEmailDotCom(t *testing.T){
	emailAddress := "@email.com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := false

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}

func TestIsEmailAddressValid(t *testing.T){
	emailAddress := "test@email.com"
	result := models.IsEmailAddressValid(emailAddress)
	expected := true

	if result != expected {
			t.Errorf("got %t, wanted %t", result, expected)
	}
}






