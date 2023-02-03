package biz

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	Username, Password, Realname, Nickname, AvatarUrl, Phone string
	UserIP, State, Deleted                                   int
	CreateTime, UpdateTime                                   time.Time
}

type Users struct {
	Collention    []*User
	NextPageToken string
}

type UserRepo interface {
	ListUsers(ctx context.Context, parent string) (*Users, error)
	GetUser(ctx context.Context, name string) (*User, error)
	SearchUsers(ctx context.Context, name string) (*Users, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (pu *UserUsecase) ListUsers(ctx context.Context, parent string) (*Users, error) {
	return pu.repo.ListUsers(ctx, parent)
}

func (pu *UserUsecase) GetUser(ctx context.Context, name string) (*User, error) {
	return pu.repo.GetUser(ctx, name)
}

func (pu *UserUsecase) SearchUsers(ctx context.Context, name string) (*Users, error) {
	return pu.repo.SearchUsers(ctx, name)
}

func (pu *UserUsecase) CreateUser(ctx context.Context, user *User) (*User, error) {
	return pu.repo.CreateUser(ctx, user)
}

func (pu *UserUsecase) UpdateUser(ctx context.Context, user *User) (*User, error) {
	return pu.repo.UpdateUser(ctx, user)
}

func (pu *UserUsecase) DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error) {
	return pu.repo.DeleteUser(ctx, name)
}
