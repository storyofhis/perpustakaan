package views

import (
	"time"

	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories/models"
)

type Register struct {
	Id        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	NIM       string    `json:"nim"`
	Jurusan   string    `json:"jurusan"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Role      models.Role
}

type Login struct {
	Token string `json:"token"`
}
