package api

import (
	"github.com/gin-gonic/gin"
	"message-board/tool"
)

func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithDate(ctx, "请登陆后进行操作")
		ctx.Abort() //终止后续操作
	}

	ctx.Set("username", username) //将获取到的cookie的值写入上下文
	ctx.Next()
}
