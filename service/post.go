package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddPost(post model.Post) error {
	err := dao.InsertPost(post)
	return err
}
