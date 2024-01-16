package models

import "time"

type Item struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	YaerRefer  int    `json:"year_id"`
	Year       Year   `gorm:"foreignKey:YaerRefer"`
	Date       string `json:"date"`
	Name       string `json:"name"`
	Text       string `json:"text" gorm:"text"`
	SourceLink string `json:"source_link"`
	ImageReal  string `json:"imageReal"` // changed from [3]string
	ImageAi    string `json:"imageAi"`   // changed from [3]string
}
