package wigle

import (
	"reflect"
	"testing"
)

func TestParseModesValid(t *testing.T) {
	value := "[WPA2-PSK-CCMP][WPS-AUTH][ESS]"
	expected := []string{
		"WPA2-PSK-CCMP", "WPS-AUTH", "ESS",
	}
	actual, _ := parseModes(value)
	logTempl := "expected:%v got:%v"
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(logTempl, expected, actual)
	}
	t.Logf(logTempl, expected, actual)
}

func TestParseModesErr(t *testing.T) {
	modeTable := map[string]error{
		"THIS MUST FAIL": badMode,
		"[SUBTLE FAIL":   badMode,
		"[ALL GOOD]":     nil,
	}
	for m, e := range modeTable {
		logTempl := "tested %s for error. expected:%v got:%v"
		_, err := parseModes(m)
		if err != e {
			t.Fatalf(logTempl, m, e, err)
		}
		t.Logf(logTempl, m, e, err)
	}
}
