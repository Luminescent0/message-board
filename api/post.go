package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func postDetail(ctx *gin.Context) {
	postIdString := ctx.Param("post_id") //返回url参数的值
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "post_id格式有误")
		return
	}

	//根据postId拿到post
	post, err := service.GetPostById(postId)
	if err != nil {
		fmt.Println("get post by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	//找到它的评论
	comments, err := service.GetPostComments(postId)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("get post comments err: ", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	var postDetail model.PostDetail
	postDetail.Post = post
	postDetail.Comments = comments

	fmt.Println("123")
	tool.RespSuccessfulWithDate(ctx, postDetail)
}

func briefPosts(ctx *gin.Context) {
	posts, err := service.GetPosts()
	if err != nil {
		fmt.Println("get posts err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, posts)
}

func addPost(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //从上下文跨中间件取值
	username := iUsername.(string)

	txt := verifypost(ctx)

	post := model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}

	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
func changePost(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //从上下文跨中间件取值
	username := iUsername.(string)

	txt := verifypost(ctx)

	post := model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.ChangePost(post)
	if err != nil {
		fmt.Println("change post err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)

}

func deletePost(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //从上下文跨中间件取值
	username := iUsername.(string)

	post := model.Post{
		Txt:        "该留言已删除",
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.DeletePost(post)
	if err != nil {
		fmt.Println("delete post err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)

}
func verifypost(ctx *gin.Context) string {
	validate := validator.New() //创建验证器
	txt := ctx.PostForm("txt")
	err := validate.Struct(txt)
	if err != nil {
		tool.RespErrorWithDate(ctx, err)
		return "留言长度错误"
	}
	return txt
}
