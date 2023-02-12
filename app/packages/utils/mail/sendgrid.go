package mail

import (
	"gopkg.in/gomail.v2"
	"log"
)

type Mail struct {
	Subject     string
	To          string
	TextContent string
	HtmlContent string
}

func SendMail(m Mail) error {

	log.Println(m.Subject, m.To, m.TextContent, m.HtmlContent)

	mail := gomail.NewMessage()
	mail.SetHeader("From", "test@example.com")
	mail.SetHeader("To", m.To)
	mail.SetHeader("Subject", m.Subject)
	mail.SetBody("text/html", m.HtmlContent)

	d := gomail.Dialer{Host: "mailhog", Port: 1025}
	if err := d.DialAndSend(mail); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
