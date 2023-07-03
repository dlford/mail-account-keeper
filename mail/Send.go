package mail

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"

	"mail-account-keeper/config"
)

func Send(c *config.AccountConfig, a *config.AlertConfig) {
	time.Sleep(1 * time.Second)

	m := gomail.NewMessage()
	m.SetHeader("From", c.Email)
	m.SetHeader("To", c.MailTo)
	m.SetHeader("Subject", "mail-account-keeper")
	m.SetBody("text/plain", "This message was sent via github.com/dlford/mail-account-keeper to protect the account \""+c.Title+"\" from auto-deletion.")

	d := gomail.NewDialer(c.Host, c.Port, c.Email, c.Password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Failed to send mail from account \"" + c.Title + "\": ")
		fmt.Println(err)
		if a.Email != "" && a.Password != "" && a.MailTo != "" && a.Host != "" && a.Port != 0 {
			am := gomail.NewMessage()
			am.SetHeader("From", a.Email)
			am.SetHeader("To", a.MailTo)
			am.SetHeader("Subject", "mail-account-keeper failed to send mail")
			am.SetBody("text/plain", "Failed to send mail from account \""+c.Title+"\"!\n\n"+err.Error())
			ad := gomail.NewDialer(a.Host, a.Port, a.Email, a.Password)
			if err := ad.DialAndSend(am); err != nil {
				fmt.Printf("Failed to send alert from account \"" + a.Email + "\": ")
				fmt.Println(err)
			} else {
				fmt.Printf("Successfully sent alert from account \"%s\"!\n", a.Email)
			}
		}
	} else {
		fmt.Printf("Successfully sent mail from account \"%s\"!\n", c.Title)
	}
}
