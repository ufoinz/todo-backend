package persistence

import (
	"todo-app/internal/domain/event"

	"gorm.io/gorm"
)

type PostgresEventRepo struct{ DB *gorm.DB }

func NewPostgresEventRepo(db *gorm.DB) event.Repository {
	return &PostgresEventRepo{DB: db}
}

func (r *PostgresEventRepo) Create(e *event.Event) error {
	return r.DB.Create(e).Error
}

func (r *PostgresEventRepo) GetAll() ([]event.Event, error) {
	var list []event.Event
	if err := r.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *PostgresEventRepo) GetByID(id int64) (*event.Event, error) {
	var ev event.Event
	if err := r.DB.First(&ev, id).Error; err != nil {
		return nil, err
	}
	return &ev, nil
}

func (r *PostgresEventRepo) Update(e *event.Event) error {
	return r.DB.Save(e).Error
}

func (r *PostgresEventRepo) Delete(id int64) error {
	return r.DB.Delete(&event.Event{}, id).Error
}
