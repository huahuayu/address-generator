package user

import (
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/address-generator/common/request"
	"github.com/huahuayu/address-generator/service/user"
	"github.com/huahuayu/address-generator/view"
)

func Register(c *gin.Context) {
	req := new(view.RegisterReq)
	request.GetReq(c, req)
	user.Register(c, req)
}

func Login(c *gin.Context) {
	req := new(view.LoginReq)
	request.GetReq(c, req)
	user.Login(c, req)
}

func UpdatePassword(c *gin.Context) {
	req := new(view.UpdatePasswordReq)
	request.GetReq(c, req)
	user.UpdatePassword(c, req)
}

func UpdateUsername(c *gin.Context) {
	req := new(view.UpdateUsernameReq)
	request.GetReq(c, req)
	user.UpdateUsername(c, req)
}

func Delete(c *gin.Context) {
	user.Delete(c)
}

func Logout(c *gin.Context) {
	user.Logout(c)
}

func Info(c *gin.Context) {
	user.Info(c)
}

func AddressGenerator(c *gin.Context){
	user.AddressGenerator(c)
}
