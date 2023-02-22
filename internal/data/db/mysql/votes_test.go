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

func TestListVotes(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryVote().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*Vote{
		{
			Title:       "testInsertVote1",
			A:           "Mazzya",
			B:           "Mazzyb",
			C:           "Mazzyc",
			D:           "Mazzyd",
			HasTxtField: 1,
		},
		{
			Title: "testInsertVote2",
			A:     "Mazza",
			B:     "Mazzb",
			C:     "Mazzc",
			D:     "Mazzd",
			E:     "Mazze",
		},
		{
			Title:   "testInsertVote3",
			IsRadio: 0,
			A:       "Mazzy3a",
			B:       "Mazzy3b",
			C:       "Mazzy3c",
		},
	}
	for _, tc := range tcs {
		err := c.DatabaseClient.InsertVote(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUpdateVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	vote := &Vote{
		Id:      id,
		Title:   "tttest",
		IsRadio: 0,
		A:       "updated Opt A",
	}
	getVote := func() *Vote {
		ps := [4]string{"id", "=", strconv.Itoa(vote.Id), "or"}
		got, err := c.DatabaseClient.QueryVote().Where(ps).First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getVote()
	if err := c.DatabaseClient.UpdateVote(context.Background(), vote); err != nil {
		t.Error(err)
		return
	}
	after := getVote()
	if before.Title != after.Title {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				vote.Title, after.Title))
		}
	}
	if before.A != after.A {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				vote.A, after.A))
		}
	}
	if before.IsRadio != after.IsRadio {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				vote.IsRadio, after.IsRadio))
		}
	}
}

func TestDeleteVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeleteVote(context.Background(), id); err != nil {
		t.Fatalf("DeleteVote err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryVote().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryVote err: %v", err)
	}
	if got != nil {
		t.Error(errors.New("Delete failed."))
	}
}
