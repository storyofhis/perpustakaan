package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/storyofhis/perpustakaan-backend-go/config"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/controllers"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/repositories/gorm"
	"github.com/storyofhis/perpustakaan-backend-go/httpserver/services"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Connection Load is Error : %s\n", err)
		return
	}
}
func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Printf("Connection Database is Error : %s\n", err)
	}

	router := gin.Default()
	config.GenerateJwtSignature()

	// user
	userRepo := gorm.NewUserRepo(db)
	userSvc := services.NewUserSvc(userRepo)
	userControllers := controllers.NewUserController(*userSvc)

	app := httpserver.NewRoutes(router, userControllers)
	app.Start(":" + os.Getenv("PORT"))

}
