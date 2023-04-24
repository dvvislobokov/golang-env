package env

import (
	"os"
	"strconv"
	"strings"
)

func String(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func Int(name string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		return defaultValue
	}
	return value
}

func Float32(name string, defaultValue float32) float32 {
	value, err := strconv.ParseFloat(os.Getenv(name), 32)
	if err != nil {
		return defaultValue
	}
	return float32(value)
}

func Float(name string, defaultValue float64) float64 {
	value, err := strconv.ParseFloat(os.Getenv(name), 64)
	if err != nil {
		return defaultValue
	}
	return value
}

func Bool(name string) bool {
	return strings.ToLower(os.Getenv(name)) == "true"
}
