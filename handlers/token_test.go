package handlers

import (
	"testing"
	"time"

	c "desq.com.ru/testjwtauth/config"
	"desq.com.ru/testjwtauth/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateToken(t *testing.T) {
	DB, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	DB.AutoMigrate(&models.RefreshToken{})
	userID := uuid.New()
	IP := "127.0.0.1"
	expires := time.Now().Add(time.Hour * 24).Unix()
	type args struct {
		DB      *gorm.DB
		userID  uuid.UUID
		IP      string
		expires int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateToken",
			args: args{
				DB:      DB,
				userID:  userID,
				IP:      IP,
				expires: expires,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := CreateToken(tt.args.DB, tt.args.userID, tt.args.IP, tt.args.expires)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRefreshToken(t *testing.T) {
	DB, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	DB.AutoMigrate(&models.RefreshToken{})
	userID := uuid.New()
	IP := "127.0.0.1"
	expires := time.Now().Add(time.Hour * 24)
	_, RT, err := CreateToken(DB, userID, IP, expires.Unix())
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}
	claims := jwt.MapClaims{
		"userId": userID,
		"ip":     IP,
		"sub":    RT,
		"exp":    expires,
	}
	type args struct {
		DB     *gorm.DB
		userID uuid.UUID
		claims jwt.MapClaims
		RT     string
		IP     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RefreshToken",
			args: args{
				DB:     DB,
				userID: userID,
				claims: claims,
				RT:     RT,
				IP:     IP,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RefreshToken(tt.args.DB, tt.args.userID, tt.args.claims, tt.args.RT, tt.args.IP); (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
