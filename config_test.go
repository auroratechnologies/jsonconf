package jsonconf

import "testing"

func TestFileNotExist(t *testing.T) {
	err := LoadConfig("fake.json")
	if err == nil {
		t.Error("Should return back an error for incorrect file")
	}
}

func TestSingleVal(t *testing.T) {
	err := LoadConfig("config_test.json")
	if err != nil {
		t.Error("Should load the file given")
	}

	v, err := GetVar("FOO")
	if err != nil {
		t.Error("Should return the value of FOO")
	}

	if v != "Hi" {
		t.Error("Wrong value retried for FOO")
	}

	i, err := GetVar("BAR")
	if err != nil {
		t.Error("Should return the value of BAR")
	}
	if i != "1" {
		t.Error("String conversion didn't work")
	}

	_, err = GetVar("FAKE")
	if err == nil {
		t.Error("Should return an error for an entry that doesn't exist")
	}

}

func TestDefaultValue(t *testing.T) {
	err := LoadConfig("config_test.json")
	if err != nil {
		t.Error("Should load the file given")
	}

	v, err := GetVar("FOO", "Bye!")
	if err != nil {
		t.Error("Should return the value of FOO")
	}

	if v != "Hi" {
		t.Error("Wrong value retried for FOO")
	}

	v, err = GetVar("FAKE", "BAR")
	if err != nil || v != "BAR" {
		t.Error("Did not correctly load from default with fake value")
	}

	v, err = GetVar("FAKE", "BAR", "BAZ")
	if err == nil {
		t.Error("Multiple defaults should have caused error")
	}
}

func TestSetVal(t *testing.T) {
	err := LoadConfig("config_test.json")
	if err != nil {
		t.Error("Should load the file given")
	}

	v, err := GetVar("FOO")
	if err != nil {
		t.Error("Should return the value of FOO")
	}

	if v != "Hi" {
		t.Error("Wrong value retried for FOO")
	}

	SetVar("FOO", "BAR")
	v, err = GetVar("FOO")
	if err != nil {
		t.Error("Should return the value of FOO")
	}

	if v != "BAR" {
		t.Error("Wrong value retried for FOO")
	}
}
