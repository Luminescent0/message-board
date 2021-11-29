package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register") //注册
		userGroup.POST("/login")    //登陆
		userGroup.POST("/password") //修改密码

		userGroup.POST("/info") //修改个人信息
		userGroup.GET("/info")  //查看个人信息
	}

	postGroup := engine.Group("/post")
	{
		postGroup.POST("/")         //发布新留言
		postGroup.POST("/:post_id") //修改留言

		postGroup.GET("/")         //查看全部留言概略
		postGroup.GET("/:post_id") //查看一条留言详细信息和其下属评论
	}

	commentGroup := engine.Group("/comment")
	{
		commentGroup.POST("/")              //发送评论
		commentGroup.DELETE("/:comment_id") //删除评论
	}

	engine.Run()
}
