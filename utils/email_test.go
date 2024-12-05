package utils

import (
	"testing"

	"desq.com.ru/testjwtauth/config"
)

func TestNotifyEmail(t *testing.T) {
	type args struct {
		m string
		r string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "NotifyEmail",
			args: args{
				m: "127.0.0.1",
				r: config.MockSMTPRecipient,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NotifyEmail(tt.args.m, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("NotifyEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
