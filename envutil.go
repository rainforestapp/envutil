// envutil contains utility functions for dealing with environment
// variables.
package envutil

import (
	"fmt"
	"os"
	"strconv"
)

// GetenvStr returns the value of an environment varaiable, using a
// default of def if the variable is unset or emtpy.
func GetenvStr(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

// MustGetenv returns the value of an environment variable and panics
// if the variable isn't set or is empty.
func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment value %s must be set", key))
	}
	return val
}

// MustGetenvStr is an alias for the MustGetenv which fits the rest
// of function naming.
func MustGetenvStr(key string) string {
	return MustGetenv(key)
}

// GetenvInt returns an environment variable parsed as an integer,
// returning a default of def if the variable isn't set or is
// empty. If the variable is set but can't be parsed to an integer, it
// panics.
func GetenvInt(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("Error parsing value for %s: %v", key, err))
	}
	return ret
}

// MustGetenvInt returns an environment variable parsed as an integer,
// panicking if the integer is not set, is empty, or can't be parsed
// as an integer.
func MustGetenvInt(key string) int {
	if os.Getenv(key) == "" {
		panic(fmt.Sprintf("MustGetenvInt: %s must be set", key))
	}
	return GetenvInt(key, 0)
}

// GetenvBool returns an environment variable parsed as a boolean,
// returning a default value of def if the variable isn't set or is
// empty. If the variable is set but can't be parsed to a boolean, it
// panics.
func GetenvBool(key string, def bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	trueVals := []string{"true", "1"}
	falseVals := []string{"false", "0"}

	for _, candidate := range trueVals {
		if val == candidate {
			return true
		}
	}

	for _, candidate := range falseVals {
		if val == candidate {
			return false
		}
	}

	panic(fmt.Sprintf("GetenvBool: %s could not be parsed as an integer", val))
}

// MustGetenvBool returns an environment variable parsed as a boolean,
// panicking if the integer is not set, is empty, or can't be parsed
// as a boolean.
func MustGetenvBool(key string) bool {
	if os.Getenv(key) == "" {
		panic(fmt.Sprintf("MustGetenvBool: %s must be set", key))
	}
	return GetenvBool(key, false)
}
