package toolkit

import (
	"github.com/gin-gonic/gin"
)

func Render() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		MakeResp(c)
	}
}
