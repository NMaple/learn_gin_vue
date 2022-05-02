package router

import (
	"github.com/NMaple/learn_gin_vue/controller"
	"github.com/gin-gonic/gin"
)

func LoginRouter(c *gin.Engine) {
	c.GET("login", controller.Login)
	c.Group("group", controller.Ping)
	{

	}
}
