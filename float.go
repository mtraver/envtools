package envtools

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloat64(key string) (float64, error) {
	return strconv.ParseFloat(os.Getenv(key), 64)
}

func GetFloat64WithDefault(key string, defaultVal float64) float64 {
	f, err := GetFloat64(key)
	if err != nil {
		return defaultVal
	}
	return f
}

func MustGetFloat64(key string) float64 {
	f, err := GetFloat64(key)
	if err != nil {
		panic(fmt.Sprintf("failed to parse float64 from environment variable: %s\n", key))
	}
	return f
}
