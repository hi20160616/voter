package data

import (
	"context"
	"log"
	"regexp"
	"time"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ biz.PostRepo = new(postRepo)

type postRepo struct {
	data *Data
	log  *log.Logger
}

func NewPostRepo(data *Data, logger *log.Logger) biz.PostRepo {
	return &postRepo{
		data: data,
		log:  log.Default(),
	}
}

func (pr *postRepo) ListPosts(ctx context.Context, parent string) (*biz.Posts, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	ps := &mysql.Posts{}
	bizps := &biz.Posts{}
	var err error

	re := regexp.MustCompile(`^(categories|tags)/(.+)/posts$`)
	x := re.FindStringSubmatch(parent)
	if len(x) != 3 {
		ps, err = pr.data.DBClient.DatabaseClient.QueryPost().All(ctx)
	} else {
		clause := [4]string{}
		if x[1] == "categories" {
			clause = [4]string{"category_id", "=", x[2], "and"}
		}
		if x[1] == "users" {
			clause = [4]string{"users_id", "=", x[2], "and"}
		}
		ps, err = pr.data.DBClient.DatabaseClient.QueryPost().
			Where(clause).All(ctx)
	}
	if err != nil {
		return nil, err
	}

	for _, p := range ps.Collection {
		bizps.Collection = append(bizps.Collection, &biz.Post{
			PostId:     p.Id,
			Title:      p.Title,
			Detail:     p.Detail,
			CreateTime: p.CreateTime,
			UpdateTime: p.UpdateTime,
		})
	}
	return bizps, nil
}

func (pr *postRepo) GetPost(ctx context.Context, name string) (*biz.Post, error) {
	return nil, nil
}

func (pr *postRepo) SearchPosts(ctx context.Context, name string) (*biz.Posts, error) {
	return nil, nil
}

func (pr *postRepo) CreatePost(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	return nil, nil
}

func (pr *postRepo) UpdatePost(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	return nil, nil
}

func (pr *postRepo) DeletePost(ctx context.Context, name string) (*emptypb.Empty, error) {
	return nil, nil
}
