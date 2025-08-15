package user

// User — сущность пользователя
type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `gorm:"not null" json:"-"`
	Name     string `json:"name" bindeing:"required"`
}

// RegisterRequest — DTO для запроса регистрации нового пользователя
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequest — DTO для запроса авторизации
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
