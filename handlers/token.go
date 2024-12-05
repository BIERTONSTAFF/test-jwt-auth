package handlers

import (
	"fmt"

	"desq.com.ru/testjwtauth/config"
	"desq.com.ru/testjwtauth/models"
	"desq.com.ru/testjwtauth/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateToken(DB *gorm.DB, userID uuid.UUID, IP string, expires int64) (string, string, error) {
	RT, err := utils.GenerateToken(32)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token")
	}

	claims := jwt.MapClaims{
		"userId": userID,
		"ip":     IP,
		"sub":    RT,
		"exp":    expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign token")
	}

	RTEncoded, err := utils.EncodeToken(RT)
	if err != nil {
		return "", "", fmt.Errorf("failed to encode refresh token")
	}

	if err := DB.Create(&models.RefreshToken{
		UserID: userID,
		Token:  RTEncoded,
		Valid:  true,
	}).Error; err != nil {
		return "", "", fmt.Errorf("failed to create refresh token")
	}

	return t, RT, nil
}

func RefreshToken(DB *gorm.DB, userID uuid.UUID, claims jwt.MapClaims, RT string, IP string) error {
	if claims["sub"].(string) != RT {
		return fmt.Errorf("refresh token does not belong to this token")
	}

	validRTs := []models.RefreshToken{}
	if err := DB.Where("user_id = ? AND valid = ?", userID, true).Find(&validRTs).Error; err != nil {
		return fmt.Errorf("failed to retrieve valid tokens")
	}

	var matchedRT *models.RefreshToken
	for i := range validRTs {
		if utils.CompareToken(validRTs[i].Token, RT) {
			matchedRT = &validRTs[i]

			break
		}
	}

	if matchedRT == nil || !matchedRT.Valid {
		return fmt.Errorf("token is invalid")
	}

	matchedRT.Valid = false
	if err := DB.Save(&matchedRT).Error; err != nil {
		return fmt.Errorf("failed to invalidate refresh token")
	}

	if claims["ip"].(string) != IP {
		if err := utils.NotifyEmail(IP, config.MockSMTPRecipient); err != nil {
			fmt.Printf("Failed to send email: %v\n", err)
		}
	}

	return nil
}
