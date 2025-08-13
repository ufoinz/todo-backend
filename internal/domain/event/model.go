package event

import "time"

type Event struct {
	ID      int64     `gorm:"primaryKey" json:"id"`
	OwnerId int64     `gorm:"not null" json:"owner_id"`
	Name    string    `json:"name" binding:"required,min=3"`
	Content string    `json:"content" binding:"required,min=10"`
	Time    time.Time `json:"time" binding:"required"`
}
