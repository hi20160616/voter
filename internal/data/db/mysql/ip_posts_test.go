package mysql

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var ipPostid = 1

func TestListIpPosts(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryIpPost().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertIpPost(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*IpPost{
		{
			Ip:     "127.0.0.1",
			PostId: 1,
		},
		{
			Ip:     "127.0.0.2",
			PostId: 2,
		},
		{
			Ip:     "127.0.0.3",
			PostId: 3,
		},
	}
	for _, tc := range tcs {
		x, err := c.DatabaseClient.InsertIpPost(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
		fmt.Println("id: ", x)
	}
}

func TestUpdateIpPost(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	ipIpPost := &IpPost{
		Id:     ipPostid,
		Ip:     "123.123.123.124",
		PostId: 4,
	}
	getIpPost := func() *IpPost {
		ps := [4]string{"id", "=", strconv.Itoa(ipIpPost.Id), "or"}
		got, err := c.DatabaseClient.QueryIpPost().Where(ps).
			First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getIpPost()
	if err := c.DatabaseClient.UpdateIpPost(context.Background(),
		ipIpPost); err != nil {
		t.Error(err)
		return
	}
	after := getIpPost()
	if before.Ip != after.Ip {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				ipIpPost.Ip, after.Ip))
		}
	}
}

func TestDeleteIpPost(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeleteIpPost(context.Background(), id); err != nil {
		t.Fatalf("DeleteIpPost err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryIpPost().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryIpPost err: %v", err)
	}
	if got != nil {
		t.Error(errors.New("Delete failed."))
	}
}
