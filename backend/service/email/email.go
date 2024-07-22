package email

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/ajtroup1/speakeasy/config"
)

// SendEmail sends an email using the provided recipient, subject, and body.
func SendEmail(to string, subject string, body string) error {
	from := config.Envs.Email
	password := config.Envs.EmailPassword

	log.Printf("EMAIL CREDENTIALS: %s // %s", from, password)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
		if err == nil {
			log.Printf("Email sent successfully on attempt %d", attempt)
			return nil // Email sent successfully
		}

		log.Printf("Attempt %d: Failed to send email: %v", attempt, err)

		if attempt < maxRetries {
			time.Sleep(2 * time.Second)
		}
	}

	return fmt.Errorf("failed to send email after %d attempts", maxRetries)
}
