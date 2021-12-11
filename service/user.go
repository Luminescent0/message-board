package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

func ChangePassword(username, newPassword string) error {
	err := dao.UpdatePassword(username, newPassword)
	return err
}

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows { //用户名不存在
			return false, nil
		}

		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	return true, nil
}

// IsRepeatUsername 判断用户名是否重复
func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows { //没有匹配查询的行
			return false, nil
		}

		return false, err //查询失败
	}

	return true, nil //用户已存在
}

func Register(user model.User) error {
	err := dao.Cipher(user)
	if err != nil {
		return err
	}
	err = dao.InsertUser(user)
	return err
}
