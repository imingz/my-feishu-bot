package handler

import (
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
)

var handler = dispatcher.NewEventDispatcher("", "").
	OnP2MessageReceiveV1(P2MessageReceive)

func Get() *dispatcher.EventDispatcher {
	return handler
}
