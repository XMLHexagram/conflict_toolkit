package toolkit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Worker struct {
	*zap.Logger
}

var worker Worker

func Init(logger *zap.Logger) {
	worker.Logger = logger
}

func Log(c *gin.Context, resp interface{}, responseType ResponseType) {
	var err error
	interfaceErr, exist := c.Get(Error)
	if exist {
		err = interfaceErr.(error)
	}
	worker.Logger.WithOptions(zap.AddCallerSkip(1)).Info("response",
		zap.String("requestId", ""),
		zap.String("staffId", ""),
		zap.String("method", c.Request.Method),
		zap.String("URI", c.Request.RequestURI),
		zap.Int8("responseType", int8(responseType)),
		zap.Reflect("body", resp),
		zap.String("UA", c.GetHeader("user-agent")),
		zap.String("remoteIP", c.ClientIP()),
		zap.Error(err),
		//zap.Int64("timestamp", time.Now().Unix()),
	)
}
