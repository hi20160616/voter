package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data"
	"github.com/hi20160616/voter/internal/data/db/mysql"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	pb.UnimplementedUsersAPIServer
	uc *biz.UserUsecase
}

func NewUserService() (*UserService, error) {
	dbc, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewUserRepo(db, log.Default())
	userUsecase := biz.NewUserUsecase(repo, *log.Default())
	return &UserService{uc: userUsecase}, nil
}

func (as *UserService) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListUsers: \n%v\n", r)
		}
	}()
	bizus, err := as.uc.ListUsers(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.User{}
	for _, u := range bizus.Collection {
		resp = append(resp, &pb.User{
			UserId:     int32(u.UserId),
			Username:   u.Username,
			Password:   u.Password,
			Realname:   u.Realname,
			Nickname:   u.Nickname,
			AvatarUrl:  u.AvatarUrl,
			Phone:      u.Phone,
			UserIp:     u.UserIp,
			State:      int32(u.State),
			Deleted:    int32(u.Deleted),
			CreateTime: timestamppb.New(u.CreateTime),
			UpdateTime: timestamppb.New(u.UpdateTime),
		})
	}
	return &pb.ListUsersResponse{Users: resp}, nil
}

func (us *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetUser: %s\n%v\n", in.Name, r)
		}
	}()
	bizu, err := us.uc.GetUser(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		UserId:     int32(bizu.UserId),
		Username:   bizu.Username,
		Password:   bizu.Password,
		Realname:   bizu.Realname,
		Nickname:   bizu.Nickname,
		AvatarUrl:  bizu.AvatarUrl,
		Phone:      bizu.Phone,
		UserIp:     bizu.UserIp,
		State:      int32(bizu.State),
		Deleted:    int32(bizu.Deleted),
		CreateTime: timestamppb.New(bizu.CreateTime),
		UpdateTime: timestamppb.New(bizu.UpdateTime),
	}, nil
}

func (us *UserService) SearchUsers(ctx context.Context, in *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchUsers: \n%v\n", r)
		}
	}()
	bizus, err := us.uc.SearchUsers(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchUsersResponse{}
	for _, e := range bizus.Collection {
		resp.Users = append(resp.Users, &pb.User{
			UserId:     int32(e.UserId),
			Username:   e.Username,
			Password:   e.Password,
			Realname:   e.Realname,
			Nickname:   e.Nickname,
			AvatarUrl:  e.AvatarUrl,
			Phone:      e.Phone,
			UserIp:     e.UserIp,
			State:      int32(e.State),
			Deleted:    int32(e.Deleted),
			CreateTime: timestamppb.New(e.CreateTime),
			UpdateTime: timestamppb.New(e.UpdateTime),
		})
	}
	return resp, nil
}

func (us *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateUsers: \n%v\n", r)
		}
	}()
	u, err := us.uc.CreateUser(ctx, &biz.User{
		UserId:    int(in.User.UserId),
		Username:  in.User.Username,
		Password:  in.User.Password,
		Realname:  in.User.Realname,
		Nickname:  in.User.Nickname,
		AvatarUrl: in.User.AvatarUrl,
		Phone:     in.User.Phone,
		UserIp:    in.User.UserIp,
		State:     int(in.User.State),
	})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		UserId:     int32(u.UserId),
		Username:   u.Username,
		Password:   u.Password,
		Realname:   u.Realname,
		Nickname:   u.Nickname,
		AvatarUrl:  u.AvatarUrl,
		Phone:      u.Phone,
		UserIp:     u.UserIp,
		State:      int32(u.State),
		Deleted:    int32(u.Deleted),
		CreateTime: timestamppb.New(u.CreateTime),
		UpdateTime: timestamppb.New(u.UpdateTime),
	}, nil
}

func (us *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateUsers: \n%v\n", r)
		}
	}()
	u, err := us.uc.UpdateUser(ctx, &biz.User{
		UserId:    int(in.User.UserId),
		Username:  in.User.Username,
		Password:  in.User.Password,
		Realname:  in.User.Realname,
		Nickname:  in.User.Nickname,
		AvatarUrl: in.User.AvatarUrl,
		Phone:     in.User.Phone,
		UserIp:    in.User.UserIp,
		State:     int(in.User.State),
	})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		UserId:     int32(u.UserId),
		Username:   u.Username,
		Password:   u.Password,
		Realname:   u.Realname,
		Nickname:   u.Nickname,
		AvatarUrl:  u.AvatarUrl,
		Phone:      u.Phone,
		UserIp:     u.UserIp,
		State:      int32(u.State),
		UpdateTime: timestamppb.New(u.UpdateTime),
		CreateTime: timestamppb.New(u.CreateTime),
	}, nil
}

func (as *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateUsers: \n%v\n", r)
		}
	}()
	return nil, nil
	// return as.ac.DeleteUser(ctx, in.Name)
}
