package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"pingPong/internal/config"
)

func SendConfirmation(toEmail, code string) error {
	cfg := config.LoadConfig()
	subject := "Код подтверджения"
	body := fmt.Sprintf("Ваш код подтверджения: %s\nЕсли вы не регистрировались - проигнорируйте это письмо.", code)
	message := []byte("Subject: " + subject + "\r\n " +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\";\r\n" +
		"\r\n" + body + "\r\n")

	if cfg.SMTPHost == "" || cfg.SMTPUser == "" || cfg.SMTPPass == "" {
		log.Printf("[DEV EMAIL] To: %s Code: %s\n", toEmail, code)
		return nil
	}

	address := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)
	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
	if err := smtp.SendMail(address, auth, cfg.MailFrom, []string{toEmail}, message); err != nil {
		return err
	}
	return nil
}
