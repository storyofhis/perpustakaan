package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers/params"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers/views"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories/models"
	"gorm.io/gorm"
)

type BooksSvc struct {
	repo repositories.BooksRepo
}

func NewBookSvc(repo repositories.BooksRepo) *BooksSvc {
	return &BooksSvc{
		repo: repo,
	}
}

func (svc *BooksSvc) CreateBooks(ctx context.Context, books *params.CreateBook) *views.Response {
	model := models.Books{
		Title:    books.Title,
		Author:   books.Author,
		Quantity: books.Quantity,
		Place:    books.Place,
	}
	if err := svc.repo.CreateBooks(ctx, &model); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateBookBooks{
		Id:       model.Id,
		Title:    model.Title,
		Author:   model.Author,
		Quantity: model.Quantity,
		Place:    model.Place,
	})
}

func (svc *BooksSvc) GetBooks(ctx context.Context) *views.Response {
	books, err := svc.repo.GetAllBooks(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	response := make([]views.GetBooks, 0)
	for i := 0; i < len(books); i++ {
		response = append(response, views.GetBooks{
			Id:        books[i].Id,
			Title:     books[i].Title,
			Author:    books[i].Author,
			Quantity:  books[i].Quantity,
			Place:     books[i].Place,
			CreatedAt: books[i].CreatedAt,
			DeletedAt: books[i].DeletedAt,
		})
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, response)
}

func (svc *BooksSvc) UpdateBooks(ctx context.Context, params *params.UpdateBook, BookId uint) *views.Response {
	model, err := svc.repo.GetBooksById(ctx, BookId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("Books with this id is not exists"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	model.Quantity = params.Quantity
	err = svc.repo.UpdateBooksById(ctx, model, BookId)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateBook{
		Id:       model.Id,
		Title:    model.Title,
		Author:   model.Author,
		Quantity: model.Quantity,
		Place:    model.Place,
	})
}
