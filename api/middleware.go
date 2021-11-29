package api

import (
	"github.com/gin-gonic/gin"
	"message-board/tool"
)

func auth(ctx *gin.Context) {
	_, exist := ctx.Get("username")
	if !exist {
		tool.RespErrorWithDate(ctx, "请登陆后进行操作")
		ctx.Abort()
	}

	ctx.Next()
}
