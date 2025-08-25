package utils

import (
	"os"
)

func GetEnv(key, fallback string) string {
	str := os.Getenv(key)
	if str == "" {
		return fallback
	}
	return str
}
