package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
  	"github.com/cuappdev/hustle-backend/models"
  	"github.com/cuappdev/hustle-backend/controllers"
	"github.com/cuappdev/hustle-backend/firebaseadmin"
	"github.com/cuappdev/hustle-backend/middleware"  
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file.")
	}

	log.Println("Starting hustle-backend...")
  	r := gin.Default()
	log.Println("Connecting to database...")
	models.ConnectDatabase()

	ac, err := firebaseadmin.NewAuthClient(context.Background(), os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil { log.Fatalf("firebase init: %v", err) }

	log.Println("Setting up routes...")
	// Public routes
	r.GET("/healthcheck", controllers.HealthCheck)

	// Protected routes
	authd := r.Group("/api")
	authd.Use(middleware.RequireFirebaseUser(ac))
	{
		authd.GET("/users", controllers.FindUsers)
		authd.POST("/users", controllers.CreateUser)
	}
	log.Println("Server starting on :8080")

  	r.Run()
}


// TODO

// Copy over .github folders
// Docker file and docker compose ymlhow