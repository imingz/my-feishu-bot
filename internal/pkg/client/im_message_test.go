package client

import (
	"testing"
)

func TestClient_Im_Message_Reply(t *testing.T) {
	Get().Im_Message_Reply("om_b0d3002db89790f29b1c2f96fd7881fe", Im_Message_Reply_Request{
		Content: `{
    "image_key": "img_v3_02bv_a9ea4eb5-2a2e-498f-8ba9-7b2f95e407eg"
}`,
		MsgType:       "image",
		ReplyInThread: true,
		Uuid:          "",
	})
}
