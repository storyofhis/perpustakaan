package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers/params"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/services"
)

type UserControllers struct {
	svc services.UserSvc
}

func NewUserController(svc services.UserSvc) *UserControllers {
	return &UserControllers{svc: svc}
}

func (control *UserControllers) Register(ctx *gin.Context) {
	var req params.Register
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = validator.New().Struct(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := control.svc.Register(ctx, &req)
	WriteJsonResponse(ctx, response)

}

func (control *UserControllers) Login(ctx *gin.Context) {
	var req params.Login
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = validator.New().Struct(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := control.svc.Login(ctx, &req)
	WriteJsonResponse(ctx, response)

}
