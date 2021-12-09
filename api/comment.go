package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func addComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	txt := ctx.PostForm("txt")
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "文章id有误")
		return
	}

	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
func amendComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	txt := ctx.PostForm("txt")
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err:", err)
		tool.RespErrorWithDate(ctx, "文章id有误")
		return
	}
	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
	}
	err = service.AmendComment(comment)
	if err != nil {
		fmt.Println("amend comment err : ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
func deleteComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //iUsername为空接口类型
	username := iUsername.(string)

	txt := "该评论已删除"
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err", err)
		tool.RespErrorWithDate(ctx, "文章id有误")
		return
	}
	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
	}
	err = service.AmendComment(comment)
	if err != nil {
		fmt.Println("delete comment err : ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
