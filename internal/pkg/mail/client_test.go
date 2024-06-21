package mail

import (
	"log"
	"testing"
	"xiaoxiaojiqiren/config"
)

func TestClient_newMsg(t *testing.T) {
	config := config.NewConfig().Mail
	c := New(&config)

	m := c.newMsg()
	if err := m.To("wmz@mail.xrqw.site"); err != nil {
		t.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("This is test newMsg!")

	if err := c.client.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
}
