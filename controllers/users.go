package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/cuappdev/hustle-backend/models"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
  // Validate input
  var input models.CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  uid := middleware.UIDFrom(c)
  if uid == "" {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "midding firebase uid"})
    return
  }

  // Create user
  user := models.User{
    FirstName: input.FirstName, 
    LastName: input.LastName, 
    Email: input.Email,
    Firebase_UID: uid
  }
  models.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{"data": user})
}

