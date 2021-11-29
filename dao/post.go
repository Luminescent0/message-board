package dao

import "message-board/model"

func InsertPost(post model.Post) error {
	_, err := dB.Exec("INSERT INTO post(username, txt, post_time, update_time) "+"values(?, ?, ?, ?);", post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}
