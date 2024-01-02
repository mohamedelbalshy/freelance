package main

import (
	"fmt"
	"freelance/controller"
	"freelance/database"
	"freelance/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()
	loadDatabase()
	serveApplication()

}

func loadDatabase() {
	database.Connect()
	// database.Database.AutoMigrate(&model.User{}, &model.Project{})
	// database.Database.AutoMigrate(&model.Project{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func serveApplication() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"data": "Hello World!"}) })
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.GET("/projects", controller.GetAllProjects)
	protectedRoutes.POST("/projects", controller.CreateProject)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
