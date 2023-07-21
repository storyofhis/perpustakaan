package gorm

import (
	"context"
	"strings"
	"time"

	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) CreateUser(ctx context.Context, user *models.Users) error {
	user.CreatedAt = time.Now()
	err := repo.db.WithContext(ctx).Create(user).Error
	return err
}

func (repo *userRepo) FindUserByNIM(ctx context.Context, nim string) (*models.Users, error) {
	user := new(models.Users)
	err := repo.db.WithContext(ctx).Where("LOWER(nim) = ?", strings.ToLower((nim))).Take(user).Error
	return user, err
}

func (repo *userRepo) FindUserById(ctx context.Context, id uint) (*models.Users, error) {
	user := new(models.Users)
	err := repo.db.WithContext(ctx).Where("id = ?", id).Take(user).Error
	return user, err
}
