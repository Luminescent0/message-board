package dao

import "message-board/model"

func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(username, txt, comment_time, post_id) "+"values(?, ?, ?, ?);", comment.Username, comment.Txt, comment.CommentTime, comment.PostId)
	return err
}
