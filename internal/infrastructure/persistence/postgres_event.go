package persistence

import (
	"todo-app/internal/domain/event"

	"gorm.io/gorm"
)

// Реализация event.Repository на основе GORM/Postgres
type PostgresEventRepo struct{ DB *gorm.DB }

// Создаёт новый репозиторий событий, используя переданное соединение GORM
func NewPostgresEventRepo(db *gorm.DB) event.Repository {
	return &PostgresEventRepo{DB: db}
}

// сохраняет новое событие e в таблицу events
func (r *PostgresEventRepo) Create(e *event.Event) error {
	return r.DB.Create(e).Error
}

// возвращает все события из таблицы events
func (r *PostgresEventRepo) GetAll() ([]event.Event, error) {
	var list []event.Event
	if err := r.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// ищет событие по его ID в таблице
func (r *PostgresEventRepo) GetByID(id int64) (*event.Event, error) {
	var ev event.Event
	if err := r.DB.First(&ev, id).Error; err != nil {
		return nil, err
	}
	return &ev, nil
}

// сохраняет изменения объекта e в базе
func (r *PostgresEventRepo) Update(e *event.Event) error {
	return r.DB.Save(e).Error
}

// удаляет событие по его ID
func (r *PostgresEventRepo) Delete(id int64) error {
	return r.DB.Delete(&event.Event{}, id).Error
}
