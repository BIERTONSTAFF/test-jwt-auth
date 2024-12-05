package config

import (
	"log"
	"os"
)

var (
	DSN               = getEnv("DSN")
	SMTPHost          = getEnv("SMTP_HOST")
	SMTPEmail         = getEnv("SMTP_EMAIL")
	SMTPPassword      = getEnv("SMTP_PASSWORD")
	MockSMTPRecipient = getEnv("MOCK_SMTP_RECIPIENT")
	JWTSecret         = getEnv("JWT_SECRET")
)

func getEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("enviroment variable %s is required", k)
	}

	return v
}
