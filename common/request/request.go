package request

import (
	"github.com/bytedance/go-tagexpr/validator"
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/address-generator/global"
	"github.com/huahuayu/address-generator/view"
	"github.com/sirupsen/logrus"
)

var vd = validator.New("vd")

func GetReq(c *gin.Context, req interface{}) {
	if err := c.Bind(req); err != nil {
		view.ResponseErr(c, global.ErrInvalidParam, err.Error())
		return
	}
	logrus.Info("request: "+c.Request.RequestURI+"|", req)
	if err := vd.Validate(req); err != nil {
		view.ResponseErr(c, global.ErrRequestValidationNotPass, err.Error())
		return
	}
}
