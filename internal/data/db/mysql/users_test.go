package mysql

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var id = 1

func TestPrepareQuery(t *testing.T) {
	qc := &UserQuery{query: "SELECT * FROM users"}
	qc.Where(
		[4]string{"name", "like", "test", "and"},
		[4]string{"name", "like", "test1", "and"},
		[4]string{"name", "like", "test2", "and"},
		[4]string{"name", "like", "test3", "and"},
	)
	if err := qc.prepareQuery(context.Background()); err != nil {
		t.Error(err)
	}
	fmt.Println(qc.query, qc.args)
}

func TestListUsers(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryUser().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertUser(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*User{
		{
			Username:  "testInsert1",
			Password:  "testInsert1",
			Realname:  "Mazzy1",
			Nickname:  "donkey1",
			AvatarUrl: "testInsert1.jpg",
			Phone:     "13512345678",
			UserIp:    "123.123.123.1",
		},
		{
			Username:  "testInsert2",
			Password:  "testInsert2",
			Realname:  "Mazzy2",
			Nickname:  "donkey2",
			AvatarUrl: "testInsert2.jpg",
			Phone:     "13512345678",
			UserIp:    "123.123.123.2",
		},
		{
			Username:  "testInsert3",
			Password:  "testInsert3",
			Realname:  "Mazzy3",
			Nickname:  "donkey3",
			AvatarUrl: "testInsert3.jpg",
			Phone:     "13512345678",
			UserIp:    "123.123.123.3",
		},
	}
	for _, tc := range tcs {
		err := c.DatabaseClient.InsertUser(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUpdateUser(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{
		Id:        id,
		Username:  "tttest",
		Password:  "testNewPwd",
		Realname:  "real test",
		Nickname:  "nick test",
		AvatarUrl: "avatar_url.test.jpg",
		Phone:     "13512345678",
		UserIp:    "111.111.111.111",
		State:     1,
	}
	getUser := func() *User {
		ps := [4]string{"id", "=", strconv.Itoa(user.Id), "or"}
		got, err := c.DatabaseClient.QueryUser().Where(ps).First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getUser()
	if err := c.DatabaseClient.UpdateUser(context.Background(), user); err != nil {
		t.Error(err)
		return
	}
	after := getUser()
	if before.Password != after.Password {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				user.Password, after.Password))
		}
	}
	if before.Realname != after.Realname {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				user.Realname, after.Realname))
		}
	}
	if before.Nickname != after.Nickname {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				user.Nickname, after.Nickname))
		}
	}
	if before.AvatarUrl != after.AvatarUrl {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				user.AvatarUrl, after.AvatarUrl))
		}
	}
	if before.Phone != after.Phone {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				user.Phone, after.Phone))
		}
	}
	if before.UserIp != after.UserIp {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				user.UserIp, after.UserIp))
		}
	}
	if before.State != after.State {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				user.State, after.State))
		}
	}
}

func TestDeleteUser(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeleteUser(context.Background(), id); err != nil {
		t.Fatalf("DeleteUser err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryUser().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryUser err: %v", err)
	}
	if got.Deleted != 1 {
		t.Error(errors.New("Delete failed."))
	}
}

func TestUnDeleteUser(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.UndeleteUser(context.Background(), id); err != nil {
		t.Fatalf("UndeleteUser err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryUser().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryUser err: %v", err)
	}
	if got.Deleted != 0 {
		t.Error(errors.New("UnDelete failed."))
	}

}
