package hhtclient

import (
	"context"
	"fmt"
	"testing"
)

func TestGetRoomBalance(t *testing.T) {
	balance, err := GetRoomBalance(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("balance: %v\n", balance)
}
