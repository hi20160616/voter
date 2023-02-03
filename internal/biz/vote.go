package biz

import "time"

type Vote struct {
	IsRadio                int
	Title, Detail          string
	CreateTime, UpdateTime time.Time
}
