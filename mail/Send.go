package mail

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"

	"mail-account-keeper/config"
)

func Send(a *config.AccountConfig) {
	time.Sleep(1 * time.Second)

	m := gomail.NewMessage()
	m.SetHeader("From", a.Email)
	m.SetHeader("To", a.MailTo)
	m.SetHeader("Subject", "mail-account-keeper")
	m.SetBody("text/plain", "This message was sent via github.com/dlford/mail-account-keeper to protect the account \""+a.Title+"\" from auto-deletion.")

	d := gomail.NewDialer(a.Host, a.Port, a.Email, a.Password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Failed to send mail from account \"" + a.Title + "\"!")
		fmt.Println(err)
	} else {
		fmt.Printf("Successfully sent mail from account \"%s\"!\n", a.Title)
	}
}
