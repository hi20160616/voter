package data

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
)

var repo4 = func() biz.IpPostRepo {
	dc, err := mysql.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewIpPostRepo(&Data{DBClient: dc}, log.Default())
}()

func TestGetIpPost(t *testing.T) {
	x, err := repo4.GetIpPost(context.Background(), "ip_posts/3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Ip: ", x.Ip)
	fmt.Println("IpPostId: ", x.IpPostId)
}

func TestCreateIpPost(t *testing.T) {
	tcs := []*biz.IpPost{
		{Ip: "127.0.0.1", PostId: 1},
		{Ip: "127.0.0.1", PostId: 2},
		{Ip: "127.0.0.1", PostId: 3},
	}

	for _, tc := range tcs {
		x, err := repo4.CreateIpPost(context.Background(), tc)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(x)
	}
}

func TestUpdateIpPost(t *testing.T) {
	x, err := repo4.UpdateIpPost(context.Background(), &biz.IpPost{
		IpPostId: 1,
		Ip:       "12.12.12.123",
		PostId:   1,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Ip: ", x.Ip)
	fmt.Println("IpPostId: ", x.IpPostId)
	fmt.Println("PostId: ", x.PostId)
}

func TestListIpPosts(t *testing.T) {
	x, err := repo4.ListIpPosts(context.Background(), "ip/127.0.0.1/ip_posts")
	y, err := repo4.ListIpPosts(context.Background(), "ip_posts")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("=========================ip: 127.0.0.1 posts:")
	for _, e := range x.Collection {
		fmt.Println("Id: ", e.IpPostId)
		fmt.Println("Ip: ", e.Ip)
		fmt.Println("PostId: ", e.PostId)
		fmt.Println("------------------")
	}
	fmt.Println("=========================ip: * posts:")
	for _, e := range y.Collection {
		fmt.Println("Id: ", e.IpPostId)
		fmt.Println("Ip: ", e.Ip)
		fmt.Println("PostId: ", e.PostId)
		fmt.Println("------------------")
	}
}
