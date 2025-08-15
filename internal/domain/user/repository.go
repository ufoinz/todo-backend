package user

// Repository — интерфейс абстракции над хранилищем пользователей
// Позволяет работать с данными без привязки к БД
type Repository interface {
	Insert(u *User) error
	GetByEmail(email string) (User, error)
	GetByID(id int64) (User, error)
}
