package service

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/configs"
)

var cfg = configs.NewConfig("voter")

func TestListPosts(t *testing.T) {
	in := &pb.ListPostsRequest{}
	as, err := ListPosts(context.Background(), in, cfg)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(as.Posts)
	for i, a := range as.Posts {
		fmt.Println(i, a.Title, " ", a.Detail, " ", a.IsOpen)
	}
}

func TestGetPost(t *testing.T) {
	id := 1
	in := &pb.GetPostRequest{Name: "posts/" + strconv.Itoa(id)}
	a, err := GetPost(context.Background(), in, cfg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
