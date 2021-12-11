package model

import "time"

type Comment struct {
	Id          int
	PostId      int
	Txt         string `validate:"max=200"`
	Username    string
	CommentTime time.Time
}
type Comment_next struct {
	Id          int
	CommentId   int
	Txt         string
	Username    string
	CommentTime time.Time
}
