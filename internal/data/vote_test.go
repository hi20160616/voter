package data

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
)

var repoTestVotes = func() biz.VoteRepo {
	dc, err := mysql.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewVoteRepo(&Data{DBClient: dc}, log.Default())
}()

func TestListVotes(t *testing.T) {
	votes, err := repoTestVotes.ListVotes(context.Background(), "pid/6/votes")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range votes.Collection {
		fmt.Println("Title: ", e.Title)
		fmt.Println("A: ", e.A)
		fmt.Println("B: ", e.B)
		fmt.Println("C: ", e.C)
		fmt.Println("D: ", e.D)
	}
}
