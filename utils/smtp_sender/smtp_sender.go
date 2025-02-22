package smtp_sender

import (
	"STDE_proj/utils/converter"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"os"
	"time"
)

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// SendEmail sends an email with the verification code
func SendEmail(to string, code string) error {
	// Конфигурация SMTP клиента
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Код подтверждения")
	m.SetBody("text/plain", fmt.Sprintf("Ваш код подтверждения: %s", code))
	d := gomail.NewDialer(smtpHost, converter.StoI(smtpPort), smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
