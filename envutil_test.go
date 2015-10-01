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
