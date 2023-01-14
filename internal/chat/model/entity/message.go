package entity

import "github.com/kkakoz/pkg/timex"

type Message struct {
	ID       int64      `json:"id"`
	Content  string     `json:"content"`
	TheadId  int64      `json:"thead_id"`
	UserId   int64      `json:"user_id"`
	ToUserId int64      `json:"to_user_id"`
	Created  timex.Time `json:"created"`
}
