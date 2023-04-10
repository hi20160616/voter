package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ biz.IpPostRepo = new(ipPostRepo)

type ipPostRepo struct {
	data *Data
	log  *log.Logger
}

func NewIpPostRepo(data *Data, logger *log.Logger) biz.IpPostRepo {
	return &ipPostRepo{
		data: data,
		log:  log.Default(),
	}
}

func (ipr *ipPostRepo) ListIpPosts(ctx context.Context, parent string) (*biz.IpPosts, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	ips := &mysql.IpPosts{}
	bizips := &biz.IpPosts{}
	var err error

	re := regexp.MustCompile(`^(ip)/(.+)/ip_posts$`)
	x := re.FindStringSubmatch(parent)
	if len(x) != 3 {
		ips, err = ipr.data.DBClient.DatabaseClient.QueryIpPost().All(ctx)
	} else {
		ip := inet_aton(x[2])
		ips, err = ipr.data.DBClient.DatabaseClient.QueryIpPost().
			Where([4]string{"ip", "=", ip}).All(ctx)
	}

	if err != nil {
		return nil, err
	}

	for _, p := range ips.Collection {
		bizips.Collection = append(bizips.Collection, &biz.IpPost{
			IpPostId: p.Id,
			Ip:       p.Ip,
			PostId:   p.PostId,
		})
	}
	return bizips, nil
}

func (ipr *ipPostRepo) GetIpPost(ctx context.Context, name string) (*biz.IpPost, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^ip_posts/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New(
			"GetIpPost: name: " + name + " cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id}
	p, err := ipr.data.DBClient.DatabaseClient.QueryIpPost().Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.IpPost{
		IpPostId: p.Id,
		Ip:       p.Ip,
		PostId:   p.PostId,
	}, nil
}

// CreateIpPost insert a row to database if ip and post not exist, otherwise, update this row
func (ipr *ipPostRepo) CreateIpPost(ctx context.Context, ipPost *biz.IpPost) (*biz.IpPost, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	ip := inet_aton(ipPost.Ip)
	clauses := [][4]string{
		{"ip", "=", ip, "and"},
		{"post_id", "=", strconv.Itoa(ipPost.PostId), "and"},
	}

	_, err := ipr.data.DBClient.DatabaseClient.QueryIpPost().
		Where(clauses...).First(ctx)
	if err != nil {
		if errors.Is(err, mysql.ErrNotFound) {
			ippostid, err := ipr.data.DBClient.DatabaseClient.
				InsertIpPost(ctx, &mysql.IpPost{
					Ip:     ipPost.Ip,
					PostId: ipPost.PostId,
				})
			if err != nil {
				return nil, err
			} else {
				ipPost.IpPostId = int(ippostid)
				return ipPost, nil
			}
		} else {
			return nil, err
		}
	}
	// err is nil means ip_id and post_id query back, it is exist.
	// update the row
	return ipr.UpdateIpPost(ctx, ipPost)
}

func (ipr *ipPostRepo) UpdateIpPost(ctx context.Context, ipPost *biz.IpPost) (*biz.IpPost, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbIpPost := &mysql.IpPost{}
	var err error
	if ipPost.IpPostId == 0 {
		ip := inet_aton(ipPost.Ip)
		dbIpPost, err = ipr.data.DBClient.DatabaseClient.QueryIpPost().Where(
			[][4]string{
				{"ip", "=", ip, "and"},
				{"post_id", "=", strconv.Itoa(ipPost.PostId), "and"},
			}...).First(ctx)
		if err != nil {
			return nil, errors.Join(fmt.Errorf(
				"ipPostRepo: UpdateIpPost: query ipPost by ip"+
					"and post_id error: %v: %v", ipPost, err))
		}
	} else {
		dbIpPost, err = ipr.data.DBClient.DatabaseClient.QueryIpPost().Where(
			[4]string{"id", "=", strconv.Itoa(ipPost.IpPostId), "and"},
		).First(ctx)
		if err != nil {
			return nil, errors.Join(fmt.Errorf(
				"ipPostRepo: UpdateIpPost: query ipPost by id error: %v: %v",
				ipPost, err))
		}
	}
	if &ipPost.Ip != nil {
		dbIpPost.Ip = ipPost.Ip
	}
	if &ipPost.PostId != nil {
		dbIpPost.PostId = ipPost.PostId
	}
	if err := ipr.data.DBClient.DatabaseClient.UpdateIpPost(ctx,
		dbIpPost); err != nil {
		return nil, err
	}
	return &biz.IpPost{
		IpPostId: dbIpPost.Id,
		Ip:       dbIpPost.Ip,
		PostId:   dbIpPost.PostId,
	}, nil
}

func (ipr *ipPostRepo) DeleteIpPost(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^ip_posts/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("ipPostRepo: DeleteIpPost: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("ipPostRepo: DeleteIpPost: ipPost id should be integer only")
	}
	return &emptypb.Empty{}, ipr.data.DBClient.DatabaseClient.DeleteIpPost(ctx, id)
}
