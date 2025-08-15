package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Ошибка при неверной авторизации
var ErrInvalidCredentials = errors.New("invalid credentials")

// Интерфейс бизнес-логики пользователей (Регистрация)
type Service interface {
	Register(req RegisterRequest) (User, error)
	Authenticate(email, password string) (User, error)
}

type service struct {
	repo Repository
}

// Конструктор сервиса пользователей
func NewService(r Repository) Service {
	return &service{repo: r}
}

// Хэшируем пароль и сохраняем нового пользователя
func (s *service) Register(req RegisterRequest) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	u := User{
		Email:    req.Email,
		Password: string(hash),
		Name:     req.Name,
	}
	if err := s.repo.Insert(&u); err != nil {
		return User{}, err
	}
	return u, nil
}

// Проверяем email и пароль, сравнивая с хэшем
func (s *service) Authenticate(email, password string) (User, error) {
	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return User{}, ErrInvalidCredentials
	}
	return u, nil
}
