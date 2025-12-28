package Smtp

import (
	"Polybub/Utilities"
	"net/smtp"
	"strconv"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(recipients []string, message []byte) error {
	auth := smtp.PlainAuth("", "", "", "")

	err := smtp.SendMail("", auth, "", []string{}, []byte{})
	if err != nil {
		return err
	}

	return nil
}

func SendgridEmail(userId int32, resetKey string, givenEmail string) (*rest.Response, error) {
	gc := Utilities.GlobalConfig
	baseUrl := Utilities.GetBaseUrl(Utilities.GlobalConfig)
	resetUrl := baseUrl + "/forgot-password?id=" + strconv.Itoa(int(userId)) + "&key=" + resetKey
	body := "To reset your password, click <a href=\"" + resetUrl + "\">This Link</a>"

	from := mail.NewEmail(gc.SendgridName, gc.SendgridEmail)
	subject := "Subject: Password Reset - " + baseUrl
	to := mail.NewEmail("", givenEmail)
	message := mail.NewSingleEmail(from, subject, to, "", body)
	client := sendgrid.NewSendClient(gc.SendgridKey)

	response, err := client.Send(message)
	return response, err
}
