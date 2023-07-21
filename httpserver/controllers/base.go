package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers/views"
)

func WriteJsonResponse(ctx *gin.Context, res *views.Response) {
	ctx.JSON(res.Status, res)
}
