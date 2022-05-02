package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
)

func Register(ctx *gin.Context) {
	// 获取参数
	// name := ctx.PostForm("name")
	// password := ctx.PostForm("password")
	// phone := ctx.PostForm("phone")

	// 数据验证

	// 创建用户

	// 返回结果

	ctx.JSON(utils.NewSucc("Register success!", gin.H{
		"msg": "Register success!",
	}))
}
