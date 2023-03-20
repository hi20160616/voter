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

func TestListPosts(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryPost().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertPost(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*Post{
		{
			Title:  "testInsertPost1",
			Detail: "Mazzy1",
		},
		{
			Title:  "testInsertPost2",
			Detail: "Mazzy2",
		},
		{
			Title:    "testInsertPost3",
			IsClosed: 0,
			Detail:   "Mazzy3",
		},
	}
	for _, tc := range tcs {
		err := c.DatabaseClient.InsertPost(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUpdatePost(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	post := &Post{
		Id:       id,
		Title:    "tttest",
		IsClosed: 0,
	}
	getPost := func() *Post {
		ps := [4]string{"id", "=", strconv.Itoa(post.Id), "or"}
		got, err := c.DatabaseClient.QueryPost().Where(ps).First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getPost()
	if err := c.DatabaseClient.UpdatePost(context.Background(), post); err != nil {
		t.Error(err)
		return
	}
	after := getPost()
	if before.Title != after.Title {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				post.Title, after.Title))
		}
	}
	if before.Detail != after.Detail {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				post.Detail, after.Detail))
		}
	}
	if before.IsClosed != after.IsClosed {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				post.IsClosed, after.IsClosed))
		}
	}
}

func TestDeletePost(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeletePost(context.Background(), id); err != nil {
		t.Fatalf("DeletePost err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryPost().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryPost err: %v", err)
	}
	if got != nil {
		t.Error(errors.New("Delete failed."))
	}
}
