package main

import (
	"github.com/gin-gonic/gin"

  	"github.com/cuappdev/hustle-backend/models"
  	"github.com/cuappdev/hustle-backend/controllers"
)

func main() {
  	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/healthcheck", controllers.HealthCheck)
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)

  	r.Run()
}


// TODO

// Copy over .github folders
// Docker file and docker compose ymlhow