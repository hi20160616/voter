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
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^posts/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id}
	p, err := pr.data.DBClient.DatabaseClient.QueryPost().
		Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Post{
		PostId:     p.Id,
		Title:      p.Title,
		IsOpen:     p.IsOpen,
		Detail:     p.Detail,
		UpdateTime: p.UpdateTime,
		CreateTime: p.CreateTime,
	}, nil
}

func (pr *postRepo) SearchPosts(ctx context.Context, name string) (*biz.Posts, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^posts/(.+)/search$`)
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
			[4]string{"title", "like", kw, "or"},
			[4]string{"detail", "like", kw, "or"},
		)
	}
	ps, err := pr.data.DBClient.DatabaseClient.QueryPost().
		Where(cs...).All(ctx)
	if err != nil {
		return nil, err
	}
	bizps := &biz.Posts{Collection: []*biz.Post{}}
	for _, p := range ps.Collection {
		bizps.Collection = append(bizps.Collection, &biz.Post{
			PostId:     p.Id,
			Title:      p.Title,
			IsOpen:     p.IsOpen,
			Detail:     p.Detail,
			UpdateTime: p.UpdateTime,
			CreateTime: p.CreateTime,
		})
	}
	return bizps, nil
}

func (pr *postRepo) CreatePost(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := pr.data.DBClient.DatabaseClient.
		InsertPost(ctx, &mysql.Post{
			Title:  post.Title,
			IsOpen: post.IsOpen,
			Detail: post.Detail,
		}); err != nil {
		return nil, err
	}
	return post, nil
}

func (pr *postRepo) UpdatePost(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbPost, err := pr.data.DBClient.DatabaseClient.QueryPost().
		Where([4]string{"id", "=", strconv.Itoa(post.PostId), "and"}).
		First(ctx)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(
			"postRepo: UpdatePost: query post by id error: %v: %v",
			post, err))
	}
	if &post.Title != nil {
		dbPost.Title = post.Title
	}
	if &post.IsOpen != nil {
		dbPost.IsOpen = post.IsOpen
	}
	if &post.Detail != nil {
		dbPost.Detail = post.Detail
	}
	if err := pr.data.DBClient.DatabaseClient.UpdatePost(ctx,
		dbPost); err != nil {
		return nil, err
	}
	return &biz.Post{
		PostId:     dbPost.Id,
		Title:      dbPost.Title,
		IsOpen:     dbPost.IsOpen,
		Detail:     dbPost.Detail,
		CreateTime: dbPost.CreateTime,
		UpdateTime: dbPost.UpdateTime,
	}, nil
}

func (pr *postRepo) DeletePost(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^posts/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("postRepo: DeletePost: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("postRepo: DeletePost: post id should be integer only")
	}
	return &emptypb.Empty{}, pr.data.DBClient.DatabaseClient.DeletePost(ctx, id)
}
