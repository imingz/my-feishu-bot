package qrcode

import (
	"testing"
	"xiaoxiaojiqiren/internal/pkg/client"
)

func TestQrcode(t *testing.T) {
	img, err := getQrcodeFile()
	if err != nil {
		t.Errorf("error = %v", err)
	}
	client.SendImage(client.Get(), &client.SendImageRequest{
		Image:         img,
		ReceiveIdType: "open_id",
		ReceiveId:     "ou_3d3a97494b21e6b6cb60e58ee8e39e00",
	})
}
