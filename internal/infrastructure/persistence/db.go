package persistence

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB открывает соединение с PostgreSQL по строке подключения из переменной окружения DB_DSN.
// Возвращает экземпляр *gorm.DB или ошибку, если подключение не удалось.
func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
