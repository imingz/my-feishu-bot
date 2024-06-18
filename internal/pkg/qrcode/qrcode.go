package qrcode

import (
	"bytes"
	"context"
	"io"
	"xiaoxiaojiqiren/internal/pkg/qrcodeclient"

	"github.com/skip2/go-qrcode"
)

func GetQrcodeFile() (io.Reader, error) {
	gotQrcode, err := qrcodeclient.GetQrcode(context.Background())
	if err != nil {
		return nil, err
	}
	content, err := qrcode.Encode(gotQrcode, qrcode.Medium, 256)
	return bytes.NewReader(content), err
}
