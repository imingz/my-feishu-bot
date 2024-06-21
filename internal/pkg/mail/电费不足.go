package mail

import "fmt"

func (c *Client) Send电费不足(enterpriseEmail string, balance float64) error {
	m := c.newMsg()
	if err := m.To(enterpriseEmail); err != nil {
		return fmt.Errorf("设置收件人地址发送错误: %s", err)
	}
	m.Subject(fmt.Sprintf("电费不足，当前电费余额：%v", balance))

	return c.client.DialAndSend(m)
}
