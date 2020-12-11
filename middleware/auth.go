package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/address-generator/common/session"
	"github.com/huahuayu/address-generator/global"
	"github.com/huahuayu/address-generator/view"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			view.ResponseErr(c, global.ErrInvalidToken, "")
			c.Abort()
			return
		}

		user := session.Get(token)
		if user == nil {
			view.ResponseErr(c, global.ErrLogin, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
