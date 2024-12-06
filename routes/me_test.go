package routes

import (
	"net/http"
	"testing"
	"time"

	"desq.com.ru/testjwtauth/config"
	"desq.com.ru/testjwtauth/handlers"
	"desq.com.ru/testjwtauth/models"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupFiber(t *testing.T) *fiber.App {
	DB, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	DB.AutoMigrate(&models.RefreshToken{})
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", DB)

		return c.Next()
	})
	app.Post("/me/:id/token", CreateToken)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.JWTSecret)},
	}))
	app.Post("/me/token/refresh", RefreshToken)

	return app
}

func TestCreateToken(t *testing.T) {
	userID := uuid.New().String()
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "CreateTokenRoute",
			route:         "/me/" + userID + "/token",
			expectedError: false,
			expectedCode:  201,
		},
	}
	app := setupFiber(t)
	for _, test := range tests {
		req, _ := http.NewRequest("POST", test.route, nil)
		res, err := app.Test(req, -1)
		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
	}
}

func TestRefreshToken(t *testing.T) {
	DB, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	userID := uuid.New()
	IP := "127.0.0.1"
	expires := time.Now().Add(time.Hour * 24)
	token, RT, err := handlers.CreateToken(DB, userID, IP, expires.Unix())
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}
	RTCookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    RT,
		Path:     "/",
		HttpOnly: true,
		Expires:  expires,
	}
	tests := []struct {
		description   string
		route         string
		RTCookie      *http.Cookie
		token         string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "WithNoRefreshToken",
			route:         "/me/token/refresh",
			RTCookie:      &http.Cookie{},
			token:         token,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "WithNoToken",
			route:         "/me/token/refresh",
			RTCookie:      RTCookie,
			token:         "",
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "WithInvalidToken",
			route:         "/me/token/refresh",
			RTCookie:      RTCookie,
			token:         "InvalidToken",
			expectedError: false,
			expectedCode:  401,
		},
		{
			description:   "WithValidTokenPair",
			route:         "/me/token/refresh",
			RTCookie:      RTCookie,
			token:         token,
			expectedError: false,
			expectedCode:  200,
		},
	}
	app := setupFiber(t)
	for _, test := range tests {
		req, _ := http.NewRequest("POST", test.route, nil)
		req.Header.Add("Authorization", "Bearer "+test.token)
		req.AddCookie(test.RTCookie)
		res, err := app.Test(req, -1)
		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
	}
}
