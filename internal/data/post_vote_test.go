package data

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
)

var repo2 = func() biz.PostVoteRepo {
	dc, err := mysql.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewPostVoteRepo(&Data{DBClient: dc}, log.Default())
}()

func TestCreatePostVote(t *testing.T) {
	_, err := repo2.CreatePostVote(context.Background(), &biz.PostVote{
		PostId: 1,
		VoteId: 1,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestListVidsByPid(t *testing.T) {
	vids, err := repo2.ListVidsByPid(context.Background(), "post_votes/1/list")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("pid:", vids.Pid)
	for _, e := range vids.Vids {
		fmt.Println(e)
	}
}
