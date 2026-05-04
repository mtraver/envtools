package envtools

import (
	"fmt"
	"os"
	"strconv"
)

func GetInt(key string) (int, error) {
	return strconv.Atoi(os.Getenv(key))
}

func GetIntWithDefault(key string, defaultVal int) int {
	i, err := GetInt(key)
	if err != nil {
		return defaultVal
	}
	return i
}

func MustGetInt(key string) int {
	i, err := GetInt(key)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int from environment variable: %s\n", key))
	}
	return i
}

func GetInt64(key string) (int64, error) {
	return strconv.ParseInt(os.Getenv(key), 10, 64)
}

func GetInt64WithDefault(key string, defaultVal int64) int64 {
	i, err := GetInt64(key)
	if err != nil {
		return defaultVal
	}
	return i
}

func MustGetInt64(key string) int64 {
	i, err := GetInt64(key)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int64 from environment variable: %s\n", key))
	}
	return i
}
