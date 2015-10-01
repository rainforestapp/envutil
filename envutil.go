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
	} else {
		return val
	}
}

// MustGetenv returns the value of an environment variable and panics
// if the variable isn't set or is empty.
func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment value %s must be set", key))
	} else {
		return val
	}
}

// GetenvInt returns an environment variable parsed as an integer,
// returning a default of def if the variable isn't set or is
// empty. If the variable is set but can't be parsed to an integer, it
// panics.
func GetenvInt(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	} else {
		ret, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("Error parsing value for %s: %v", key, err))
		}
		return ret
	}
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
