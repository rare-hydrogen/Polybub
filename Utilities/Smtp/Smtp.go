package Smtp

import (
	"Polybub/Utilities"
	"net/smtp"
)

func SendEmail(recipients []string, message []byte) error {
	c := Utilities.GlobalConfig

	auth := smtp.PlainAuth("", c.SmtpAddress, c.SmtpPassword, c.SmtpHost)

	err := smtp.SendMail(c.SmtpHost+":"+c.SmtpPort, auth, c.SmtpAddress, recipients, message)
	if err != nil {
		return err
	}

	return nil
}
