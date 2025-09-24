package main

import (
	"log"
	"github.com/gin-gonic/gin"

  	"github.com/cuappdev/hustle-backend/models"
  	"github.com/cuappdev/hustle-backend/controllers"
)

func main() {
	log.Println("Starting hustle-backend...")
  	r := gin.Default()
	log.Println("Connecting to database...")
	models.ConnectDatabase()
	log.Println("Setting up routes...")
	r.GET("/healthcheck", controllers.HealthCheck)
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
    log.Println("Server starting on :8080")
  	r.Run()
}


// TODO

// Copy over .github folders
// Docker file and docker compose ymlhow