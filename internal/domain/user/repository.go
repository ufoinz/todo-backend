package user

type Repository interface {
	Insert(u *User) error
	GetByEmail(email string) (User, error)
	GetByID(id int64) (User, error)
}
