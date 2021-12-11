package tree

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/service"
	"strconv"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Data  int
}

func traverse(t *TreeNode) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Data, "")
	traverse(t.Right)

}
func create(v int) *TreeNode {
	var t *TreeNode
	t = insertRight(t, v)
	return t
}

func insertRight(t *TreeNode, v int) *TreeNode {

	if t == nil {
		return &TreeNode{nil, nil, v}
	}
	if v == t.Data {
		return t
	}
	t.Right = insertRight(t.Right, v)
	return t

}
func insertLeft(t *TreeNode, v int) *TreeNode {
	t.Left = insertLeft(t.Left, v)
	return t
}

func tree(ctx *gin.Context) {
	arr1, err1 := service.GetPosts()
	if err1 != nil {
		return
	}
	for _, v1 := range arr1 {
		root := v1.Id
		t := create(root)
		arr2, _ := service.GetPostComments(root)
		for _, v2 := range arr2 {
			t = insertRight(t, v2.Id)
		}
		nextPostidstring := ctx.Param("Comment_Id")
		nextPostid, _ := strconv.Atoi(nextPostidstring)
		for _, v2 := range arr2 {
			if nextPostid == v2.Id {
				t = insertRight(t, nextPostid) //同级
			} else {
				t = insertLeft(t, nextPostid) //只写了一层的话不是同级就是下一级(笃定)
			}
			traverse(t)
		}

	}

}
