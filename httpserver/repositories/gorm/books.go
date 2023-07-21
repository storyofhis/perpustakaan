package gorm

import (
	"context"
	"time"

	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories/models"
	"gorm.io/gorm"
)

type BooksRepo struct {
	db *gorm.DB
}

func NewBooksRepo(db *gorm.DB) *BooksRepo {
	return &BooksRepo{
		db: db,
	}
}

func (repo *BooksRepo) CreateBooks(ctx context.Context, books *models.Books) error {
	books.CreatedAt = time.Now()
	err := repo.db.WithContext(ctx).Create(books).Error
	return err
}

func (repo *BooksRepo) GetAllBooks(ctx context.Context) ([]models.Books, error) {
	var books []models.Books
	if err := repo.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (repo *BooksRepo) GetBooksById(ctx context.Context, id uint) error {
	books := new(models.Books)
	err := repo.db.WithContext(ctx).Where("id = ?", id).Take(books).Error
	return err
}

func (repo *BooksRepo) UpdateBooksById(ctx context.Context, books *models.Books, id uint) error {
	books.UpdatedAt = time.Now()
	return repo.db.WithContext(ctx).Model(books).Where("id = ?", id).Updates(*books).Error
}

func (repo *BooksRepo) DeleteBooks(ctx context.Context, id uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Books{}, "id = ?", id).Error
}
