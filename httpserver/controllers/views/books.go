package views

import "time"

type CreateBookBooks struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Quantity  int       `json:"quantity"`
	Place     string    `json:"place"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type GetBooks struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Quantity  int       `json:"quantity"`
	Place     string    `json:"place"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UpdateBook struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Quantity  uint      `json:"quantity"`
	Place     string    `json:"place"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
