package model

import "time"

type Post struct {
	Id         int
	CommentNum int
	Txt        string
	Username   string
	PostTime   time.Time
	UpdateTime time.Time
}
