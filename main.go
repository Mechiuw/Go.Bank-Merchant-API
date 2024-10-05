package main

import (
	"bank/app/controller"
	"bank/app/repository"
	"bank/app/service/impl"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepo := &repository.UserRepository{"data/users.json"}
	authService := &impl.AuthService{UserRepo: userRepo}
	authController := &controller.AuthController{AuthService: authService}

	r := gin.Default()

	// basic routes
	r.POST("/login", authController.Login)
	r.POST("/logout/:id", authController.Logout) // Logout with path variable for user ID

	log.Println("Server starting on :8080")
	log.Fatal(r.Run(":8080"))
}
