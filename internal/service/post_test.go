package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	v1 "github.com/hi20160616/voter/api/voter/v1"
)

var ps = func() *PostService {
	ps, err := NewPostService()
	if err != nil {
		log.Fatal(err)
	}
	return ps
}()

func TestCreatePosts(t *testing.T) {

	a, err := ps.CreatePost(context.Background(), &v1.CreatePostRequest{
		Post: &v1.Post{
			Title:  "Test CreatePost Service",
			IsOpen: 1,
			Detail: "Test CreatePost Service Detail",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestListPosts(t *testing.T) {
	ps, err := ps.ListPosts(context.Background(), &v1.ListPostsRequest{})
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range ps.Posts {
		fmt.Println(a)
	}
}

func TestGetPost(t *testing.T) {
	id := "1"
	p, err := ps.GetPost(context.Background(), &v1.GetPostRequest{Name: "posts/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("post: ", p.Title)
}

func TestSearchPosts(t *testing.T) {
	name := "posts/test3/search"
	posts, err := ps.SearchPosts(context.Background(), &v1.SearchPostsRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range posts.Posts {
		fmt.Println(v)
	}
}

func TestUpdatePost(t *testing.T) {
	a, err := ps.UpdatePost(context.Background(), &v1.UpdatePostRequest{
		Post: &v1.Post{
			PostId: 1,
			Title:  "UpdateViaPS",
			IsOpen: 1,
			Detail: "UpdateViaPS",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestDeletePost(t *testing.T) {
	id := "1"
	name := "posts/" + id + "/delete"
	if _, err := ps.DeletePost(context.Background(), &v1.DeletePostRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	_, err := ps.GetPost(context.Background(), &v1.GetPostRequest{Name: "posts/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}
