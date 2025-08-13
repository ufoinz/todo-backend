package env

import (
	"os"
	"strconv"
)

func GetEnvString(key, defaultValue string) string {
	if value, exisits := os.LookupEnv(key); exisits {
		return value
	}

	return defaultValue
}

func GetEnvInt(key string, defaultValue int) int {
	if value, exisits := os.LookupEnv(key); exisits {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return defaultValue
}
