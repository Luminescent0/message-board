package dao

import "message-board/model"

func InsertPost(post model.Post) error {
	_, err := dB.Exec("INSERT INTO post(username,"+
		" txt, post_time, update_time) "+"values(?, ?, ?, ?);",
		post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}
func ChangePost(post model.Post) error {
	_, err := dB.Exec("UPDATE post SET post_time = ?,update_time = ?,txt = ? where username = ?",
		post.PostTime, post.UpdateTime, post.Txt, post.Username)
	return err
}

//func DeletePost(post model.Post) error {
//
//	_,err := dB.Exec("update post set txt = ? where username = ? ",
//		post.Txt,post.Username)
//	return err
//} 只是覆写了所以其实没有必要单独开一个吧（？

func SelectPostById(postId int) (model.Post, error) {
	var post model.Post

	row := dB.QueryRow("SELECT id, username, txt, post_time,update_time, comment_num FROM post WHERE id = ? ",
		postId)
	if row.Err() != nil {
		return post, row.Err()
	}

	err := row.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
	if err != nil {
		return post, err
	}

	return post, nil
}

func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := dB.Query("SELECT id, username, txt, post_time, update_time, comment_num FROM post")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var post model.Post

		err = rows.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
