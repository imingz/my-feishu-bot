package mail

import (
	"log"
	"testing"
	"xiaoxiaojiqiren/config"

	"github.com/wneessen/go-mail"
)

func TestGoMail(t *testing.T) {
	config := config.NewConfig().Mail.SMTP

	m := mail.NewMsg()
	if err := m.From(config.Username); err != nil {
		t.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To("wmz@mail.xrqw.site"); err != nil {
		t.Fatalf("failed to set To address: %s", err)
	}

	m.Subject("This is my first mail with go-mail!")
	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")

	c, err := mail.NewClient(config.Host,
		mail.WithDebugLog(),
		mail.WithPort(config.Port),
		mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(config.Username),
		mail.WithPassword(config.Password))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
}
