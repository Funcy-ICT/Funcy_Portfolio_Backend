package mail

import "log"

type Mail struct {
	Subject     string
	To          string
	TextContent string
	HtmlContent string
}

func SendMail(m Mail) error {

	// ローカルではlogに出力
	log.Println(m.To)
	log.Println(m.Subject)
	log.Println(m.TextContent)
	log.Println(m.HtmlContent)

	return nil
}
