package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	v1 "github.com/hi20160616/voter/api/voter/v1"
)

var ips = func() *IpPostService {
	ips, err := NewIpPostService()
	if err != nil {
		log.Fatal(err)
	}
	return ips
}()

func TestCreateIpPost(t *testing.T) {

	a, err := ips.CreateIpPost(context.Background(), &v1.CreateIpPostRequest{
		IpPost: &v1.IpPost{
			Ip:     "127.0.0.5",
			PostId: 1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestListIpPosts(t *testing.T) {
	x, err := ips.ListIpPosts(context.Background(), &v1.ListIpPostsRequest{})
	y, err := ips.ListIpPosts(context.Background(), &v1.ListIpPostsRequest{
		Parent: "ip/127.0.0.1/ip_posts",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("x::::::::::::::::::")
	for _, a := range x.IpPosts {
		fmt.Println(a)
	}
	fmt.Println("y::::::::::::::::::")
	for _, a := range y.IpPosts {
		fmt.Println(a)
	}
}

func TestGetIpPost(t *testing.T) {
	id := "3"
	x, err := ips.GetIpPost(context.Background(), &v1.GetIpPostRequest{Name: "ip_posts/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Ip: ", x.Ip)
	fmt.Println("IpPostId: ", x.IpPostId)
}

func TestUpdateIpPost(t *testing.T) {
	a, err := ips.UpdateIpPost(context.Background(), &v1.UpdateIpPostRequest{
		IpPost: &v1.IpPost{
			IpPostId: 1,
			Ip:       "192.168.1.123",
			PostId:   1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestDeleteIpPost(t *testing.T) {
	id := "7"
	name := "ip_posts/" + id + "/delete"
	if _, err := ips.DeleteIpPost(context.Background(), &v1.DeleteIpPostRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	_, err := ips.GetIpPost(context.Background(), &v1.GetIpPostRequest{Name: "ip_posts/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}
