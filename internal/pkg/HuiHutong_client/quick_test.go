package huihutongclient

import (
	"fmt"
	"testing"
)

func TestGetRoomId(t *testing.T) {
	c := NewClient()
	res, err := c.getRoomId("6", "5", "59")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("res: %v\n", res)
}

func TestGetRoomBalance(t *testing.T) {
	c := NewClient()
	res, err := c.Get房间余额("6", "5", "59")
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
