package huihutongclient

import "fmt"

type huiHuTongResponse struct {
	Code      int64  `json:"code"`
	Data      any    `json:"data"`
	Message   string `json:"message"`
	Result    any    `json:"result"`
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
}

func (hht *huiHuTongResponse) Error() string {
	return fmt.Sprintf("[慧湖通服务端] code: %d, message: %s", hht.Code, hht.Message)
}
