package models

import "time"

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "user"
)

type Users struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	FullName  string
	NIM       string
	Jurusan   string
	Password  string
	Role      Role `gorm:"type:role;default:'user'"`
	CreatedAt time.Time
	Update    time.Time
}
