package envtools

import (
	"os"
	"strconv"
	"strings"
)

var (
	knownBoolStrings = map[string]bool{
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

func IsTruthy(key string) bool {
	val := os.Getenv(key)
	if val == "" {
		return false
	}

	// Check for known string values.
	if b, ok := knownBoolStrings[strings.ToLower(val)]; ok {
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
