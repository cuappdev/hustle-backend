package models

import "time"

type FCMToken struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    UserID    uint      `json:"user_id" gorm:"index;not null"`
    Token     string    `json:"token" gorm:"uniqueIndex;not null"`
    Platform  string    `json:"platform"` // "android" or "ios"
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    User      User      `json:"-" gorm:"foreignKey:UserID"`
}

// Save or update a token
func SaveOrUpdateToken(userID uint, token, platform string) error {
    var fcmToken FCMToken
    
    // Check if this exact token already exists
    result := DB.Where("token = ?", token).First(&fcmToken)
    
    if result.Error != nil {
        // Create new token
        fcmToken = FCMToken{
            UserID:   userID,
            Token:    token,
            Platform: platform,
        }
        return DB.Create(&fcmToken).Error
    }
    
    // Token exists - just update user/platform if needed
    return DB.Model(&fcmToken).Updates(FCMToken{UserID: userID, Platform: platform}).Error
}

// retrieves all FCM tokens for a user
func GetUserTokens(userID uint) ([]string, error) {
    var tokens []FCMToken
    err := DB.Where("user_id = ?", userID).Find(&tokens).Error
    if err != nil {
        return nil, err
    }
    
    tokenStrings := make([]string, len(tokens))
    for i, t := range tokens {
        tokenStrings[i] = t.Token
    }
    return tokenStrings, nil
}

// removes an FCM token
func DeleteToken(token string) error {
    return DB.Where("token = ?", token).Delete(&FCMToken{}).Error
}