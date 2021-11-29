package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
)

func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println("judge repeat username err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if flag {
		tool.RespErrorWithDate(ctx, "用户名已经存在")
		return
	}

	err = service.Register(user)
	if err != nil {
		fmt.Println("register err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
