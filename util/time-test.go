package util

import "testing"

// TestStringToTime blabla
func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-02-12T10:10:05")

	if convertedTime.Year() != 2019 {
		t.Errorf("Espera que o ano seja 2019")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Espera que o mês seja Fevereiro")
	}

	if convertedTime.Day() != 12 {
		t.Errorf("Espera que o dia seja 12")
	}

}
