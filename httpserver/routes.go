package httpserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/perpustakaan-backend-go/common"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers"
)

type Routes struct {
	router *gin.Engine
	user   *controllers.UserControllers
}

func NewRoutes(router *gin.Engine, user *controllers.UserControllers) *Routes {
	return &Routes{
		router: router,
		user:   user,
	}
}

func (r *Routes) Start(port string) {
	r.router.POST("/register", r.user.Register)
	r.router.POST("/login", r.user.Login)

	// run
	r.router.Run(port)
}

func (r *Routes) verifyToken(ctx *gin.Context) {
	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
	if len(bearerToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid bearer token",
		})
		return
	}
	claims, err := common.ValidateToken(bearerToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("userData", claims)
}
