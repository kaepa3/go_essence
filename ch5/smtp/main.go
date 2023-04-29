package main

import (
	"log"
	"net/smtp"

	"github.com/jhillyerd/enmime"
)

func main() {
	smtpHost := "my-mail-server:25"
	smtpAuth := smtp.PlainAuth(
		"example.com",
		"example-user",
		"example-password",
		"auth.example.com",
	)
	sender := enmime.NewSMTP(smtpHost, smtpAuth)
	master := enmime.Builder().From("宗慎太郎", "hoge@example.com"). 
		Subject("title").Text([]byte("text body")).
		HTML([]byte("<p>html mail<b> body </b> </p>"))
	//	AddFileAttachment("document.odf")
	msg := master.To("あてさき", "atge@example.com")
	err := msg.Send(sender)
	if err != nil {
		log.Fatal(err)
	}
}
