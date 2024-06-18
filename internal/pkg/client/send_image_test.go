package client

import (
	"os"
	"testing"
)

func TestSendImage(t *testing.T) {
	img, _ := os.Open("img.png")
	SendImage(client, &SendImageRequest{
		Image:         img,
		ReceiveIdType: "open_id",
		ReceiveId:     "ou_3d3a97494b21e6b6cb60e58ee8e39e00",
	})
}
