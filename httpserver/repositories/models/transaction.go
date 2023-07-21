package models

import "time"

type Transaction struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	BookId    uint
	Book      Books `gorm:"foreignKey:BookId"`
	UserId    uint
	User      Users `gorm:"foreignKey:userId"`
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
