package qrcodeclient

import (
	"context"
	"testing"
)

func TestGetQrcode(t *testing.T) {
	gotQrcode, err := GetQrcode(context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("GetQrcode() = %v", gotQrcode)
}
