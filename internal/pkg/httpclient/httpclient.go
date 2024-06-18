package httpclient

import (
	"sync"
	"xiaoxiaojiqiren/internal/pkg/httpclient/handler"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
)

var instance *gin.Engine = gin.Default()
var once sync.Once

func Get() *gin.Engine {
	once.Do(func() {
		instance.POST("/webhook/card", sdkginext.NewCardActionHandlerFunc(handler.CardHandler))
	})
	return instance
}
