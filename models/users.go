package models

type User struct {
  ID         uint   `json:"id" gorm:"primary_key"`
  FirstName  string `json:"firstname"`
  LastName   string `json:"lastname"`
  Email      string `json:"email"`
}

type CreateUserInput struct {
  FirstName  string `json:"firstname" binding:"required"`
  LastName  string `json:"lastname" binding:"required"`
  Email  string `json:"email" binding:"required"`
}