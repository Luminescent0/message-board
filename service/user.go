package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

//判断用户名是否重复
func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}
