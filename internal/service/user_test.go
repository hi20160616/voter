package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	v1 "github.com/hi20160616/voter/api/voter/v1"
)

var us = func() *UserService {
	us, err := NewUserService()
	if err != nil {
		log.Fatal(err)
	}
	return us
}()

func TestCreateUser(t *testing.T) {

	a, err := us.CreateUser(context.Background(), &v1.CreateUserRequest{
		User: &v1.User{
			Username: "user1",
			Realname: "Richael",
			UserIp:   "123.123.123.123",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestListUsers(t *testing.T) {
	us, err := us.ListUsers(context.Background(), &v1.ListUsersRequest{})
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range us.Users {
		fmt.Println(a)
	}
}

func TestGetUser(t *testing.T) {
	id := "1"
	u, err := us.GetUser(context.Background(), &v1.GetUserRequest{Name: "users/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("user ip: ", u.UserIp)
}

func TestSearchUsers(t *testing.T) {
	name := "users/user1/search"
	users, err := us.SearchUsers(context.Background(), &v1.SearchUsersRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range users.Users {
		fmt.Println(v)
	}
}

func TestUpdateUser(t *testing.T) {
	u, err := us.UpdateUser(context.Background(), &v1.UpdateUserRequest{
		User: &v1.User{
			UserId:   2,
			Password: "abcdefg",
			UserIp:   "111.111.111.111",
			Nickname: "Krasto",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func TestDeleteUser(t *testing.T) {
	id := "1"
	name := "users/" + id + "/delete"
	if _, err := us.DeleteUser(context.Background(),
		&v1.DeleteUserRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	_, err := us.GetUser(context.Background(), &v1.GetUserRequest{Name: "users/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}

func TestUndeleteUser(t *testing.T) {
	id := "1"
	name := "users/" + id + "/undelete"
	if _, err := us.UndeleteUser(context.Background(),
		&v1.UndeleteUserRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	u, err := us.GetUser(context.Background(), &v1.GetUserRequest{Name: "users/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
	if u != nil {
		fmt.Println(u)
	}
}

func TestPerDelUser(t *testing.T) {
	id := "1"
	name := "users/" + id + "/permanently_delete"
	if _, err := us.PermanentlyDeleteUser(context.Background(),
		&v1.PermanentlyDeleteUserRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	_, err := us.GetUser(context.Background(), &v1.GetUserRequest{Name: "users/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}
