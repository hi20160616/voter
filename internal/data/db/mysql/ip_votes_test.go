package mysql

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var ipVoteid = 1

func TestListIpVotes(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryIpVote().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertIpVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*IpVote{
		{
			Ip:     "127.0.0.1",
			VoteId: 1,
			Opts:   "01100000",
		},
		{
			Ip:     "127.0.0.2",
			VoteId: 2,
			Opts:   "10010000",
		},
		{
			Ip:     "127.0.0.3",
			VoteId: 3,
			Opts:   "11010000",
		},
	}
	for _, tc := range tcs {
		x, err := c.DatabaseClient.InsertIpVote(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
		fmt.Println("id: ", x)
	}
}

func TestUpdateIpVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	ipIpVote := &IpVote{
		Id: ipVoteid,
		Ip: "123.123.123.124",
	}
	getIpVote := func() *IpVote {
		ps := [4]string{"id", "=", strconv.Itoa(ipIpVote.Id), "or"}
		got, err := c.DatabaseClient.QueryIpVote().Where(ps).
			First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getIpVote()
	if err := c.DatabaseClient.UpdateIpVote(context.Background(),
		ipIpVote); err != nil {
		t.Error(err)
		return
	}
	after := getIpVote()
	if before.Ip != after.Ip {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				ipIpVote.Ip, after.Ip))
		}
	}
}

func TestDeleteIpVote(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeleteIpVote(context.Background(), id); err != nil {
		t.Fatalf("DeleteIpVote err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryIpVote().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryIpVote err: %v", err)
	}
	if got != nil {
		t.Error(errors.New("Delete failed."))
	}
}
