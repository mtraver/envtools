package envtools

import (
	"fmt"
	"os"
	"testing"
)

func IsSet(key string) bool {
	return os.Getenv(key) != ""
}

func MustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("environment variable must be set: %s\n", key))
	}
	return v
}

func MustGetenvInTest(t testing.TB, key string) string {
	t.Helper()
	val := os.Getenv(key)
	if val == "" {
		t.Fatalf("environment variable must be set: %v", key)
	}
	return val
}
