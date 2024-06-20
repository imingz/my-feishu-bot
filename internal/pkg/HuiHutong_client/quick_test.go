package huihutongclient

import (
	"fmt"
	"testing"
)

func TestGetRoomBalance(t *testing.T) {
	c := NewClient()
	res, err := c.GetRoomBalance()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("res: %v\n", res)
}

func TestGetSatoken(t *testing.T) {
	c := NewClient()
	res, err := c.getSatoken()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("res: %v\n", res)
}

func TestGetQrcodeData(t *testing.T) {
	c := NewClient()
	res, err := c.GetQrcodeData()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("res: %v\n", res)
}
