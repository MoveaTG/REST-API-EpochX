package models

import "time"

type Year struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Year      uint `json:"year"`
}
