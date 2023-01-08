package mail

import (
	"errors"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail struct {
	Subject     string
	To          string
	TextContent string
	HtmlContent string
}

func SendMail(m Mail) error {
	from := mail.NewEmail("Funcy_ICT", os.Getenv("FROM_ADDRESS"))
	to := mail.NewEmail("", m.To)

	plainTextContent := m.TextContent

	htmlContent := m.HtmlContent

	message := mail.NewSingleEmail(from, m.Subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return errors.New("wrong send mail")
	}

	return nil
}
