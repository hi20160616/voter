package biz

import "time"

type Post struct {
	IsOpen                 int
	Title, Detail          string
	CreateTime, UpdateTime time.Time
}
