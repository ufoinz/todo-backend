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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed upload .env", err)
	}

	dsn := os.Getenv("DB_DSN")
	log.Println("DB_DSN:", dsn)

	db, err := persistence.ConnectDB()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	if err = db.AutoMigrate(&userModels.User{}, &eventModels.Event{}); err != nil {
		log.Fatal("Migration error", err)
	}

	evRepo := persistence.NewPostgresEventRepo(db)
	userRepo := persistence.NewPostgresUserRepo(db)

	userSvc := userModels.NewService(userRepo)

	app := &application.Application{
		Port:        env.GetEnvInt("PORT", 8080),
		JWTSecret:   env.GetEnvString("JWT_SECRET", "defaul_secret"),
		Events:      evRepo,
		UserService: userSvc,
	}

	if err := server.Start(server.Config{
		Port:   app.Port,
		Router: app.Routes(),
	}); err != nil {
		log.Fatal(err)
	}
}
