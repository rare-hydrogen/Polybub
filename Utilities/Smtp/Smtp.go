package Smtp

import (
	"Polybub/Utilities"
	"net/smtp"
)

func SendEmail(recipients []string, message []byte) error {
	c := Utilities.GlobalConfig

	auth := smtp.PlainAuth("", c.SmtpAddress, c.SmtpPassword, c.SmtpHost)

	// TODO: For Gmail, if using remote server use port 465, otherwise use 587 for tcp (local)
	err := smtp.SendMail(c.SmtpHost+":"+c.SmtpPort, auth, c.SmtpAddress, recipients, message)
	if err != nil {
		return err
	}

	return nil
}
