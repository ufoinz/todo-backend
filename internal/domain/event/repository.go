package event

// Интерфейс абстракции над хранилищем событий
// Позволяет работать с событиями
type Repository interface {
	Create(e *Event) error
	GetAll() ([]Event, error)
	GetByID(id int64) (*Event, error)
	Update(e *Event) error
	Delete(id int64) error
}
