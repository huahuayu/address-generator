package view

import (
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/address-generator/global"
)

func ResponseOK(c *gin.Context, data interface{}, msgArg ...string) {
	var msg string
	if data == nil {
		data = ""
	}

	if msgArg == nil {
		msg = ""
	} else {
		msg = msgArg[0]
	}

	c.JSON(200, gin.H{
		"data": data,
		"code": "0000",
		"msg":  msg,
	})
}

func ResponseErr(c *gin.Context, err *global.AppErr, msg string, httpCodeArg ...int) {
	var httpCode int
	if httpCodeArg == nil {
		httpCode = 200
	} else {
		httpCode = httpCodeArg[0]
	}

	if msg == "" {
		msg = err.Msg
	}

	c.JSON(httpCode, map[string]interface{}{
		"data": "",
		"code": err.Code,
		"msg":  msg,
	})
}
