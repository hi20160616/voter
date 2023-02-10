package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/hi20160616/voter/api/voter/v1"
	_ "github.com/hi20160616/voter/configs"
	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ biz.UserRepo = new(userRepo)

type userRepo struct {
	data *Data
	log  *log.Logger
}

func NewUserRepo(data *Data, logger *log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.Default(),
	}
}

// parent=categories/*/users
// TODO parent=tags/*/users
// parent=users/*/users
func (ur *userRepo) ListUsers(ctx context.Context, parent string) (*biz.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	bizas := &biz.Users{}
	us := &mysql.Users{}
	var err error
	re := regexp.MustCompile(`^(departments|roles|usergroups)/(.+)/users$`)
	x := re.FindStringSubmatch(parent)
	y, err := regexp.MatchString(parent, `^users$`)
	if err != nil {
		return nil, err
	}
	if len(x) != 3 && y {
		us, err = ur.data.DBClient.DatabaseClient.QueryUser().All(ctx)
	} else {
		clause := [4]string{}
		switch x[1] {
		case "departments":
			clause = [4]string{"department_id", "=", x[2], "and"}
		case "roles":
			clause = [4]string{"role_id", "=", x[2], "and"}
		case "usergroups":
			clause = [4]string{"usergroup_id", "=", x[2], "and"}
		}
		us, err = ur.data.DBClient.DatabaseClient.QueryUser().
			Where(clause).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, u := range us.Collection {
		bizas.Collection = append(bizas.Collection, &biz.User{
			UserId:     u.Id,
			Username:   u.Username,
			Password:   u.Password,
			Realname:   u.Realname,
			Nickname:   u.Nickname,
			AvatarUrl:  u.AvatarUrl,
			Phone:      u.Phone,
			UserIp:     u.UserIp,
			State:      u.State,
			Deleted:    u.Deleted,
			CreateTime: u.CreateTime,
			UpdateTime: u.UpdateTime,
		})
	}
	return bizas, nil
}

func (ur *userRepo) GetUser(ctx context.Context, name string) (*biz.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// name=users/1
	re := regexp.MustCompile(`^users/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id, "and"}
	u, err := ur.data.DBClient.DatabaseClient.QueryUser().
		Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		UserId:     u.Id,
		Username:   u.Username,
		Password:   u.Password,
		Realname:   u.Realname,
		Nickname:   u.Nickname,
		AvatarUrl:  u.AvatarUrl,
		Phone:      u.Phone,
		UserIp:     u.UserIp,
		State:      u.State,
		Deleted:    u.Deleted,
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
	}, nil
}

func (ur *userRepo) SearchUsers(ctx context.Context, name string) (*biz.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^users/(.+)/search$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	kws := strings.Split(
		strings.TrimSpace(strings.ReplaceAll(x[1], "　", " ")), ",")
	cs := [][4]string{}
	for _, kw := range kws {
		cs = append(cs,
			// cs will be filtered by Where(clauses...)
			// the last `or` `and` in clause will cut off.
			// so, every clause need `or` or `and` for last element.
			[4]string{"username", "like", kw, "or"},
			[4]string{"realname", "like", kw, "or"},
			[4]string{"nickname", "like", kw, "or"},
			[4]string{"user_ip", "like", kw, "and"},
		)
	}
	us, err := ur.data.DBClient.DatabaseClient.QueryUser().
		Where(cs...).All(ctx)
	if err != nil {
		return nil, err
	}
	bizas := &biz.Users{Collection: []*biz.User{}}
	for _, e := range us.Collection {
		bizas.Collection = append(bizas.Collection, &biz.User{
			UserId:     e.Id,
			Username:   e.Username,
			Password:   e.Password,
			Realname:   e.Realname,
			Nickname:   e.Nickname,
			AvatarUrl:  e.AvatarUrl,
			Phone:      e.Phone,
			UserIp:     e.UserIp,
			State:      e.State,
			Deleted:    e.Deleted,
			CreateTime: e.CreateTime,
			UpdateTime: e.UpdateTime,
		})
	}
	return bizas, nil
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := ur.data.DBClient.DatabaseClient.
		InsertUser(ctx, &mysql.User{
			Username:  user.Username,
			Password:  user.Password,
			Realname:  user.Realname,
			Nickname:  user.Nickname,
			AvatarUrl: user.AvatarUrl,
			Phone:     user.Phone,
			UserIp:    user.UserIp,
		}); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbUser, err := ur.data.DBClient.DatabaseClient.QueryUser().
		Where([4]string{"id", "=", strconv.Itoa(user.UserId), "and"}).
		First(ctx)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(
			"userRepo: UpdateUser: query user by id error: %v: %v",
			user, err))
	}
	if &user.Username != nil {
		dbUser.Username = user.Username
	}
	if &user.Password != nil {
		dbUser.Password = user.Password
	}
	if &user.Realname != nil {
		dbUser.Realname = user.Realname
	}
	if &user.Nickname != nil {
		dbUser.Nickname = user.Nickname
	}
	if &user.AvatarUrl != nil {
		dbUser.AvatarUrl = user.AvatarUrl
	}
	if &user.Phone != nil {
		dbUser.Phone = user.Phone
	}
	if &user.UserIp != nil {
		dbUser.UserIp = user.UserIp
	}
	if &user.State != nil {
		dbUser.State = user.State
	}
	if err := ur.data.DBClient.DatabaseClient.
		UpdateUser(ctx, dbUser); err != nil {
		return nil, err
	}
	return &biz.User{
		UserId:     dbUser.Id,
		Username:   dbUser.Username,
		Password:   dbUser.Password,
		Realname:   dbUser.Realname,
		Phone:      dbUser.Phone,
		UserIp:     dbUser.UserIp,
		State:      dbUser.State,
		Deleted:    dbUser.Deleted,
		CreateTime: dbUser.CreateTime,
		UpdateTime: dbUser.UpdateTime,
	}, nil
}

// DeleteUser is soft delete, that can be undeleted, it just update deleted field to 1
func (ur *userRepo) DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^users/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("userRepo: DeleteUser: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("userRepo: DeleteUser: user id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.DeleteUser(ctx, id)
}

func (ur *userRepo) UndeleteUser(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	re := regexp.MustCompile(`^users/([\d.]+)/undelete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("userRepo: DeleteUser: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("userRepo: DeleteUser: user id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.UndeleteUser(ctx, id)
}

// DeleteUser2 is true delete row from database permanently, be careful
func (ur *userRepo) DeleteUser2(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^users/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("userRepo: DeleteUser: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("userRepo: DeleteUser: user id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.DeleteUser2(ctx, id)
}
