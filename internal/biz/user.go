package biz

import "time"

type User struct {
	Username, Password, Realname, Nickname, AvatarUrl, Phone string
	UserIP, State, Deleted                                   int
	CreateTime, UpdateTime                                   time.Time
}
