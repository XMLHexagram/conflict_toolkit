package toolkit

import "github.com/gin-gonic/gin"

func DefaultHttpEngine() *gin.Engine {
	e := gin.New()

	e.Use(gin.Recovery(), gin.Logger(), Render())

	return e
}
