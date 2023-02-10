package data

import (
	"context"
	"log"
	"testing"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
)

var repo = func() biz.PostRepo {
	dc, err := mysql.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewPostRepo(&Data{DBClient: dc}, log.Default())
}()

func TestCreatePost(t *testing.T) {
	_, err := repo.CreatePost(context.Background(), &biz.Post{
		Title:  "create post test",
		Detail: "yeah",
	})
	if err != nil {
		t.Error(err)
	}
}
