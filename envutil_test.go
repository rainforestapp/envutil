package envutil

import (
	"os"
	"testing"
)

func TestGetenvStr(t *testing.T) {
	// without the environment variable set
	want := "foo"
	got := GetenvStr("BOGUS", want)
	if want != got {
		t.Errorf("GetenvStr should have returned %s but returned %s", want, got)
	}

	// with the environment variable set
	os.Setenv("FOO", want)
	got = GetenvStr("FOO", "bogus")
	if want != got {
		t.Errorf("GetenvStr should have returned %s but returned %s", want, got)
	}
}

func TestGetenvInt(t *testing.T) {
	// without being set
	want := 42
	got := GetenvInt("BOGUS", want)
	if want != got {
		t.Errorf("GetenvInt should have returned %d but returned %d", want, got)
	}

	// with a valid value
	os.Setenv("FOO", "42")
	got = GetenvInt("FOO", 23)
	if want != got {
		t.Errorf("GetenvInt should have returned %d but returned %d", want, got)
	}

	// with an invalid value
	os.Setenv("FOO", "The meaning of life")
	defer func() {
		if r := recover(); r == nil {
			t.Error("GetenvInt should have panicked but didn't")
		}
	}()
	GetenvInt("FOO", 42)
}

func TestMustGetenvInt(t *testing.T) {
	// with a valid value
	os.Setenv("FOO", "42")
	want := 42
	got := MustGetenvInt("FOO")
	if want != got {
		t.Errorf("MustGetenvInt should have returned %d but returned %d", want, got)
	}

	// when the variable isn't set
	os.Setenv("FOO", "")
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGetenvInt should have panicked on empty value but didn't")
		}
	}()
	MustGetenvInt("FOO")
}

func TestMustGetenv(t *testing.T) {
	// with a valid value
	want := "foo"
	os.Setenv("FOO", "foo")
	got := MustGetenv("FOO")
	if want != got {
		t.Errorf("MustGetenv should have returned %s but returned %s", want, got)
	}

	// with an empty value
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGetenv should have panicked on empty value but didn't")
		}
	}()
	MustGetenv("BOGUS")
}

var boolTests = []struct {
	val      string
	expected bool
}{
	{"true", true},
	{"1", true},
	{"false", false},
	{"0", false},
}

func TestGetenvBool(t *testing.T) {
	for _, tt := range boolTests {
		os.Setenv("FOO", tt.val)
		if GetenvBool("FOO", !tt.expected) != tt.expected {
			t.Errorf("Parse error: %s should have been parsed as %v", tt.val, tt.expected)
		}
	}

	os.Setenv("FOO", "")
	if !GetenvBool("FOO", true) {
		t.Error("GetenvBool should have returned the default value of true")
	}
	if GetenvBool("FOO", false) {
		t.Error("GetenvBool should have returned the default value of false")
	}

	os.Setenv("FOO", "fdsafdsaf")
	defer func() {
		if r := recover(); r == nil {
			t.Error("GetenvBool should have panicked on invalid value")
		}
	}()
	GetenvBool("FOO", true)
}

func TestMustGetenvBool(t *testing.T) {
	for _, tt := range boolTests {
		os.Setenv("FOO", tt.val)
		if MustGetenvBool("FOO") != tt.expected {
			t.Errorf("Parse error: %s should have been parsed as %v", tt.val, tt.expected)
		}
	}

	os.Setenv("FOO", "")
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGetenvBool should have panicked on empty value")
		}
	}()
	MustGetenvBool("FOO")
}
