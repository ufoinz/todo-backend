package persistence

import (
	"todo-app/internal/domain/user"

	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) user.Repository {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Insert(u *user.User) error {
	return r.db.Create(u).Error
}

func (r *PostgresUserRepo) GetByEmail(email string) (user.User, error) {
	var u user.User
	err := r.db.Where("email = ?", email).First(&u).Error
	return u, err
}

func (r *PostgresUserRepo) GetByID(id int64) (user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	return u, err
}
