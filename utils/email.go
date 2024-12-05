package utils

import (
	"fmt"
	"net/smtp"

	c "desq.com.ru/testjwtauth/config"
)

func NotifyEmail(m string, r string) error {
	a := smtp.PlainAuth("", c.SMTPEmail, c.SMTPPassword, c.SMTPHost)

	message := "From: " + c.SMTPEmail + "\n" +
		"To: " + r + "\n" +
		"Subject: Access from an unknown IP\n\n" +
		m

	if err := smtp.SendMail(fmt.Sprintf("%s:%d", c.SMTPHost, 587), a, c.SMTPEmail, []string{r}, []byte(message)); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
