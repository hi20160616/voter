package mysql

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

// var id = 5

func TestListPostVotes(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryPostVote().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertPostVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*PostVote{
		{
			PostId: 2,
			VoteId: 1,
		},
		{
			PostId: 2,
			VoteId: 2,
		},
		{
			PostId: 2,
			VoteId: 3,
		},
	}
	for _, tc := range tcs {
		err := c.DatabaseClient.InsertPostVote(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUpdatePostVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	postPostVote := &PostVote{
		Id:     id,
		PostId: 1,
	}
	getPostVote := func() *PostVote {
		ps := [4]string{"id", "=", strconv.Itoa(postPostVote.Id), "or"}
		got, err := c.DatabaseClient.QueryPostVote().Where(ps).First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getPostVote()
	if err := c.DatabaseClient.UpdatePostVote(context.Background(), postPostVote); err != nil {
		t.Error(err)
		return
	}
	after := getPostVote()
	if before.PostId != after.PostId {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				postPostVote.PostId, after.PostId))
		}
	}
}

func TestDeletePostVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeletePostVote(context.Background(), id); err != nil {
		t.Fatalf("DeletePostVote err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryPostVote().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryPostVote err: %v", err)
	}
	if got != nil {
		t.Error(errors.New("Delete failed."))
	}
}
