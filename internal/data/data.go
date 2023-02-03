package data

import (
	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
)

var _ biz.PostRepo = new(postRepo)

type Data struct {
	DBClient *mysql.Client
}
