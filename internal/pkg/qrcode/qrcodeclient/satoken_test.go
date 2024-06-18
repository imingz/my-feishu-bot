package qrcodeclient

import (
	"context"
	"testing"
)

func TestGetSatoken(t *testing.T) {
	satoken, err := GetSatoken(context.Background())
	if err != nil {
		t.Errorf("GetSatoken() error = %v", err)
	}
	t.Logf("satoken: %v\n", satoken)
}
