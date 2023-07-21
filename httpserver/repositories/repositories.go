package repositories

import (
	"context"

	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.Users) error
	FindUserByNIM(ctx context.Context, nim string) (*models.Users, error)
	FindUserById(ctx context.Context, id uint) (*models.Users, error)
}

type BooksRepo interface {
	CreateBooks(ctx context.Context, books *models.Books) error
	GetAllBooks(ctx context.Context) ([]models.Books, error)
	GetBooksById(ctx context.Context, id uint) (*models.Books, error)
	UpdateBooksById(ctx context.Context, books *models.Books, id uint) error
	DeleteBooks(ctx context.Context, id uint) error
	// FindBooksByTransaction(ctx context.Context, id uint) ([]models.Books, error)
}
