package utils

import (
	"fmt"
	"net/smtp"

	"desq.com.ru/testjwtauth/config"
)

func NotifyEmail(m string, r string) error {
	a := smtp.PlainAuth("", config.SMTPEmail, config.SMTPPassword, config.SMTPHost)

	message := "From: " + config.SMTPEmail + "\n" +
		"To: " + r + "\n" +
		"Subject: Access from an unknown IP\n\n" +
		m

	if err := smtp.SendMail(fmt.Sprintf("%s:%d", config.SMTPHost, 587), a, config.SMTPEmail, []string{r}, []byte(message)); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
