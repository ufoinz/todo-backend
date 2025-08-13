package delivery

import (
	nethttp "net/http"
	"todo-app/internal/domain/event"
	"todo-app/internal/domain/user"
	"todo-app/internal/infrastructure/security"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Port        int
	JWTSecret   string
	Events      event.Repository
	UserService user.Service
}

func (app *Application) Routes() nethttp.Handler {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(nethttp.StatusOK, gin.H{"message": "pong"})
	})

	v1 := r.Group("/api/v1")

	users := v1.Group("/users")
	NewUserHandler(users, app.UserService, app.JWTSecret)

	events := v1.Group("/events")
	events.Use(security.JWTMiddleware(app.JWTSecret))
	{
		events.POST("/", app.CreateEvent)
		events.GET("/", app.GetEvents)
		events.GET("/:id", app.GetEvent)
		events.PUT("/:id", app.UpdateEvent)
		events.DELETE("/:id", app.DeleteEvent)
	}

	return r
}
