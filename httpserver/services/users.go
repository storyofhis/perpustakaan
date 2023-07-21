package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/storyofhis/perpustakaan-backend-go/common"
	"github.com/storyofhis/perpustakaan-backend-go/config"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers/params"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers/views"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSvc struct {
	repo repositories.UserRepo
}

func NewUserSvc(repo repositories.UserRepo) *UserSvc {
	return &UserSvc{repo: repo}
}

func (svc *UserSvc) Register(ctx context.Context, user *params.Register) *views.Response {
	_, err := svc.repo.FindUserByNIM(ctx, user.NIM)
	if err == nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_NIM_ALREADY_USED, errors.New("nim already used"))
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, errors.New("Internal Server Error"))
	}

	// generate password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, errors.New("Internal Server Error"))
	}

	// input from user
	input := models.Users{
		FullName: user.FullName,
		NIM:      user.NIM,
		Jurusan:  user.Jurusan,
		Password: string(hashedPassword),
	}

	if err = svc.repo.CreateUser(ctx, &input); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// response
	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.Register{
		Id:        input.Id,
		FullName:  input.FullName,
		NIM:       input.NIM,
		Jurusan:   input.Jurusan,
		Role:      input.Role,
		CreatedAt: input.CreatedAt,
	})
}

func (svc *UserSvc) Login(ctx context.Context, user *params.Login) *views.Response {
	models, err := svc.repo.FindUserByNIM(ctx, user.NIM)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(models.Password), []byte(user.Password))
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
	}

	role := string(models.Role)
	claims := &common.CustomClaims{
		Id:   int(models.Id),
		Role: role,
	}

	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(config.GetJwtExpiredTime())).Unix()
	claims.Subject = models.NIM

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSignature())

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Login{
		Token: ss,
	})
}
