package biz

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	UserId, State, Deleted               int
	UserIp, Username, Password           string
	Realname, Nickname, AvatarUrl, Phone string
	CreateTime, UpdateTime               time.Time
}

type Users struct {
	Collection    []*User
	NextPageToken string
}

type UserRepo interface {
	ListUsers(ctx context.Context, parent string) (*Users, error)
	GetUser(ctx context.Context, name string) (*User, error)
	SearchUsers(ctx context.Context, name string) (*Users, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error)
	UndeleteUser(ctx context.Context, name string) (*emptypb.Empty, error)
	DeleteUser2(ctx context.Context, name string) (*emptypb.Empty, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uu *UserUsecase) ListUsers(ctx context.Context, parent string) (*Users, error) {
	return uu.repo.ListUsers(ctx, parent)
}

func (uu *UserUsecase) GetUser(ctx context.Context, name string) (*User, error) {
	return uu.repo.GetUser(ctx, name)
}

func (uu *UserUsecase) SearchUsers(ctx context.Context, name string) (*Users, error) {
	return uu.repo.SearchUsers(ctx, name)
}

func (uu *UserUsecase) CreateUser(ctx context.Context, user *User) (*User, error) {
	return uu.repo.CreateUser(ctx, user)
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, user *User) (*User, error) {
	return uu.repo.UpdateUser(ctx, user)
}

func (uu *UserUsecase) DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error) {
	return uu.repo.DeleteUser(ctx, name)
}

func (uu *UserUsecase) UndeleteUser(ctx context.Context, name string) (*emptypb.Empty, error) {
	return uu.repo.UndeleteUser(ctx, name)
}

func (uu *UserUsecase) DeleteUser2(ctx context.Context, name string) (*emptypb.Empty, error) {
	return uu.repo.DeleteUser2(ctx, name)
}
