package delivery

import (
	"errors"
	"net/http"
	"todo-app/internal/domain/user"
	"todo-app/internal/infrastructure/security"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc       user.Service
	jwtSecret string
}

func NewUserHandler(r *gin.RouterGroup, svc user.Service, jwtSecret string) {
	h := &UserHandler{svc: svc, jwtSecret: jwtSecret}
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req user.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.svc.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.svc.Authenticate(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, user.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	token, err := security.GenerateToken(u.ID, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
