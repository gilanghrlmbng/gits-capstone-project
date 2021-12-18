package utils

import (
	"fmt"
	"src/config"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

const (
	CONFIG_SMTP_PORT = 587
	CONFIG_SMTP_HOST = "smtp.gmail.com"
)

func SendEmail(ctx echo.Context, email, subject, body string) error {
	config := config.GetConfigs(ctx)
	CONFIG_SENDER_NAME := fmt.Sprintf("Sma-RT No-Reply <%s>", config.Email)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		config.Email,
		config.PasswordEmail,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		ctx.Logger().Error("Mail Not Send couse of ", err.Error())
		return err
	}

	return nil
}
