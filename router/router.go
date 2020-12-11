package router

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/huahuayu/address-generator/handler/user"
	"github.com/huahuayu/address-generator/middleware"
)

func Init(r *gin.Engine) {
	r.LoadHTMLGlob("./public/*.html")
	r.Use(static.Serve("/", static.LocalFile("./public", false)))
	r.GET("/health", HealthGET)
	r.GET("/", Index)
	r.NoRoute(NotFoundError)
	r.GET("/500", InternalServerError)

	userGroup := r.Group("/user")
	userGroup.POST("/register", user.Register)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/updatePassword", middleware.AuthMiddleware(), user.UpdatePassword)
	userGroup.POST("/updateUsername", middleware.AuthMiddleware(), user.UpdateUsername)
	userGroup.GET("/info", middleware.AuthMiddleware(), user.Info)
	userGroup.GET("/logout", middleware.AuthMiddleware(), user.Logout)
	userGroup.GET("/delete", middleware.AuthMiddleware(), user.Delete)
	userGroup.GET("/address", user.AddressGenerator)
}
