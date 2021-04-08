package toolkit

import (
	"github.com/gin-gonic/gin"
)

type JsonResp struct {
	HttpCode int
	Data     interface{}
}

func SetResp(c *gin.Context, resp interface{}) {
	//_, found := c.Get(RESPONSE_FLAG_SET)
	c.Set(RESPONSE_FLAG, resp)
	c.Set(RESPONSE_FLAG_SET, true)
	return
}

func SetJson(c *gin.Context, httpCode int, data interface{}, err error) {
	if err != nil {
		SetJsonErr(c, httpCode, data, err)
	}
	SetJsonSuccess(c, httpCode, data)
}

func SetJsonErr(c *gin.Context, httpCode int, data interface{}, err error) {
	c.Set(Error, err)
	SetResp(c, &JsonResp{
		HttpCode: httpCode,
		Data:     data,
	})
}

func SetJsonSuccess(c *gin.Context, httpCode int, data interface{}) {
	SetResp(c, &JsonResp{
		HttpCode: httpCode,
		Data:     data,
	})
}

func GetResp(c *gin.Context) interface{} {
	res, ok := c.Get(RESPONSE_FLAG)
	if !ok {
		return nil
	}
	return res
}

func MakeResp(c *gin.Context) {
	resp := GetResp(c)
	switch v := resp.(type) {
	case *JsonResp:
		go Log(c, resp, Json)
		c.PureJSON(v.HttpCode, v.Data)
	}
	return
}
