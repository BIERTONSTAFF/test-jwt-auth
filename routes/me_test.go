package routes

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestCreateToken(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateToken(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRefreshToken(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RefreshToken(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
