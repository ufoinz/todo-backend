package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GenerateToken создает JWT с полем "user_id" и временем жизни 24 часа
func GenerateToken(userID int64, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// JWTMiddleware возвращает Gin middleware
func JWTMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем значение заголовка "Authorization"
		header := c.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)

		// Ожидаем формат "Bearer <token>"
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		// Парсим и проверяем подписанный токен
		tok, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		})
		if err != nil || !tok.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Извлекаем набор полей (claims) и приводим user_id к int64
		claims := tok.Claims.(jwt.MapClaims)
		id := int64(claims["user_id"].(float64))

		// Сохраняем user_id в контексте для последующих хендлеров
		c.Set("user_id", id)

		// Продолжаем выполнение цепочки
		c.Next()
	}
}
