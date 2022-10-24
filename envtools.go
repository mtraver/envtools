package envtools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	knownStrings = map[string]bool{
		"true":     true,
		"t":        true,
		"yes":      true,
		"y":        true,
		"on":       true,
		"enable":   true,
		"enabled":  true,
		"false":    false,
		"f":        false,
		"no":       false,
		"n":        false,
		"off":      false,
		"disable":  false,
		"disabled": false,
	}
)

func MustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("environment variable must be set: %s\n", key))
	}
	return v
}

func IsTruthy(key string) bool {
	val := os.Getenv(key)
	if val == "" {
		return false
	}

	// Check for known string values.
	if b, ok := knownStrings[strings.ToLower(val)]; ok {
		return b
	}

	n, err := strconv.Atoi(val)
	if err != nil {
		// It's some non-empty string value, so we'll call that true.
		return true
	}

	// Take anything greater than 0 to be true.
	return n > 0
}
