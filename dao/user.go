package dao

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"message-board/model"
)

func UpdatePassword(username, newPassword string) error {
	_, err := dB.Exec("UPDATE user SET password = ? WHERE username = ?", newPassword, username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	row := dB.QueryRow("SELECT id, password FROM user WHERE username = ? ", username) //最多返回一行的查询
	if row.Err() != nil {                                                             //检查查询错误而不调用scan，返回错误或nil
		return user, row.Err()
	}

	err := row.Scan(&user.Id, &user.Password) //返回dB.QueryRow的错误
	if err != nil {
		return user, err //没有匹配查询的行
	}

	return user, nil //获取数据库中的user对应的信息
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("INSERT INTO user(username, password) "+"values(?, ?);", user.Username, user.Password)
	return err
}

//func PasswordSalt(user model.User) string {
//	h := md5.New()
//
//	pmed5 := fmt.Sprintf("%x", h.Sum(nil))
//
//	salt1 := time.Now().Unix() //使用当前时间作为salt
//	salt2 := "%^&*"
//	//指定两个salt
//
//	//salt1+用户名+salt2+MD5拼接
//	io.WriteString(h, strconv.FormatInt(salt1, 10))
//	io.WriteString(h, user.Username)
//	io.WriteString(h, salt2)
//	io.WriteString(h, pmed5)
//
//	last := fmt.Sprintf("%x", h.Sum(nil))
//
//	return last
//} //不太确定这样能不能加上所以用了一种看不懂的（》

// Cipher 但是这里用的bcrypt加密原理和上面的好像差不多
//password传过来后通过10次（默认）循环加盐加密后得到myHash,然后拼接Bcrypt版本号（相当于上面的自定义salt吧）
//+salt+myHash等到最终的bcrypt密码，再存入数据库中
func Cipher(user model.User) error {
	hash := Bcrypt{
		cost: bcrypt.DefaultCost,
	}
	_, err := hash.Make([]byte(user.Password))
	if err != nil {
		return errors.New("加密失败")
	}
	return nil
}

type Bcrypt struct {
	cost int
}

// Make 加密
func (b *Bcrypt) Make(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.cost)
}
