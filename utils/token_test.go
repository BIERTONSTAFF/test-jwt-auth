package utils

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GenerateValidRefreshToken",
			args: args{
				l: 32,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateToken(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEncodeToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
