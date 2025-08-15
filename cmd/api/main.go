package main

import (
	"log"
	"os"

	eventModels "todo-app/internal/domain/event"
	userModels "todo-app/internal/domain/user"
	application "todo-app/internal/interface/delivery"

	"todo-app/internal/infrastructure/env"
	"todo-app/internal/infrastructure/persistence"
	"todo-app/internal/infrastructure/server"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed upload .env", err)
	}

	// Получаем строку подключения к базе данных из переменных окружения
	dsn := os.Getenv("DB_DSN")
	log.Println("DB_DSN:", dsn)

	// Подключаемся к базе данных
	db, err := persistence.ConnectDB()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	// Миграции для таблиц пользователей и событий
	if err = db.AutoMigrate(&userModels.User{}, &eventModels.Event{}); err != nil {
		log.Fatal("Migration error", err)
	}

	// Инициализируем репозитории для событий и пользователей
	evRepo := persistence.NewPostgresEventRepo(db)
	userRepo := persistence.NewPostgresUserRepo(db)

	// Создаём сервис пользователей
	userSvc := userModels.NewService(userRepo)

	// Сбор приложения
	app := &application.Application{
		Port:        env.GetEnvInt("PORT", 8080),
		JWTSecret:   env.GetEnvString("JWT_SECRET", "defaul_secret"),
		Events:      evRepo,
		UserService: userSvc,
	}

	// Запуск HTTP-сервера
	if err := server.Start(server.Config{
		Port:   app.Port,
		Router: app.Routes(),
	}); err != nil {
		log.Fatal(err)
	}
}
