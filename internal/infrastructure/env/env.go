package env

import (
	"os"
	"strconv"
)

// Возвращает строковое значение переменной окружения по ключу.
// Если переменная не найдена — возвращает значение по умолчанию.
func GetEnvString(key, defaultValue string) string {
	if value, exisits := os.LookupEnv(key); exisits {
		return value
	}

	return defaultValue
}

// Возвращает целочисленное значение переменной окружения по ключу.
// Если переменная не найдена или не может быть преобразована в int — возвращает значение по умолчанию.
func GetEnvInt(key string, defaultValue int) int {
	if value, exisits := os.LookupEnv(key); exisits {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return defaultValue
}
