package huihutongclient

func (c *Client) getSatoken() (string, error) {
	var data struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	err := c.client.Get("/web-app/auth/certificateLogin").
		SetQueryParam("openId", "or3265XAF-OMGGErEx7BW1bosxzk").
		Do().
		Into(&data)
	return data.Data.Token, err
}

func (c *Client) GetQrcodeData() (string, error) {
	satoken, err := c.getSatoken()
	if err != nil {
		return "", err
	}
	var data struct {
		Data struct {
			QrCode string `json:"qrCode"`
		} `json:"data"`
	}
	err = c.client.Get("/pms/welcome/make-code-info").
		SetHeader("satoken", satoken).
		Do().
		Into(&data)
	return data.Data.QrCode, err
}
