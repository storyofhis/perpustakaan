package models

import "time"

type Books struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	Title     string
	Author    string
	Quantity  int
	Place     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
