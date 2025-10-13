package models

type User struct {
  ID            uint   `json:"id" gorm:"primary_key"`
  Firebase_UID  string `json:"firebase_uid" gorm:"uniqueIndex"`
  Refresh_Token string `json:"refresh_token"`
  FirstName     string `json:"firstname"`
  LastName      string `json:"lastname"`
  Email         string `json:"email"`
  CreatedAt     string `json:"created_at"`
  UpdatedAt     string `json:"updated_at"`
}

type CreateUserInput struct {
  FirstName  string `json:"firstname" binding:"required"`
  LastName   string `json:"lastname" binding:"required"`
  Email      string `json:"email" binding:"required"`
}

// FindOrCreateUser finds an existing user by Firebase UID or creates a new one
func FindOrCreateUser(firebaseUID, email, firstName, lastName string) (*User, error) {
	var user User
	
	// Try to find existing user
	result := DB.Where("firebase_uid = ?", firebaseUID).First(&user)
	
	if result.Error != nil {
		// User doesn't exist, create new one
		user = User{
			Firebase_UID: firebaseUID,
			Email:        email,
			FirstName:    firstName,
			LastName:     lastName,
		}
		
		if err := DB.Create(&user).Error; err != nil {
			return nil, err
		}
	}
	
	return &user, nil
}

// UpdateRefreshToken updates the user's refresh token
func (u *User) UpdateRefreshToken(refreshToken string) error {
	return DB.Model(u).Update("refresh_token", refreshToken).Error
}