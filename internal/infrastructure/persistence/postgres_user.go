package persistence

import (
	"todo-app/internal/domain/user"

	"gorm.io/gorm"
)

// реализация user.Repository на основе GORM/Postgres
type PostgresUserRepo struct {
	db *gorm.DB
}

// создаёт новый репозиторий пользователей, используя переданное соединение GORM
func NewPostgresUserRepo(db *gorm.DB) user.Repository {
	return &PostgresUserRepo{db: db}
}

// сохраняет нового пользователя u в таблицу users
func (r *PostgresUserRepo) Insert(u *user.User) error {
	return r.db.Create(u).Error
}

// ищет пользователя по email
func (r *PostgresUserRepo) GetByEmail(email string) (user.User, error) {
	var u user.User
	err := r.db.Where("email = ?", email).First(&u).Error
	return u, err
}

// ищет пользователя по его ID
func (r *PostgresUserRepo) GetByID(id int64) (user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	return u, err
}
