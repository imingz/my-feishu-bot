package mail

import (
	"log"
	"xiaoxiaojiqiren/config"
	"xiaoxiaojiqiren/internal/pkg/env"

	"github.com/wneessen/go-mail"
)

type Client struct {
	client   *mail.Client
	username string
}

func New(mailConfig *config.MailConfig) *Client {
	c, err := mail.NewClient(mailConfig.SMTP.Host,
		mail.WithPort(mailConfig.SMTP.Port),
		mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(mailConfig.SMTP.Username),
		mail.WithPassword(mailConfig.SMTP.Password))

	if env.Active == env.DEV {
		c.SetDebugLog(true)
	}

	if err != nil {
		log.Fatalf("创建邮件客户端失败: %s", err)
	}

	return &Client{
		client:   c,
		username: mailConfig.SMTP.Username,
	}
}

func (c *Client) newMsg() *mail.Msg {
	m := mail.NewMsg()
	if err := m.From(c.username); err != nil {
		log.Fatalf("无法设置邮件发送地址: %s", err)
	}

	return m
}
