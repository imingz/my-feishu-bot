package huihutongclient

func (c *Client) GetRoomBalance() (float64, error) {
	var roomBalance struct {
		Result float64 `json:"result"`
	}
	err := c.client.Get("/proxy/qy/sdcz/getRoomBalance").
		SetQueryParams(map[string]string{
			"apartmentId": "7",
			"roomId":      "2290",
		}).
		Do().
		Into(&roomBalance)
	return roomBalance.Result, err
}
