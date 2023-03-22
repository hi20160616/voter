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

var _ biz.PostVoteRepo = new(postVoteRepo)

type postVoteRepo struct {
	data *Data
	log  *log.Logger
}

func NewPostVoteRepo(data *Data, logger *log.Logger) biz.PostVoteRepo {
	return &postVoteRepo{
		data: data,
		log:  log.Default(),
	}
}

func (pvr *postVoteRepo) ListPostVotes(ctx context.Context, parent string) (*biz.PostVotes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	pvs := &mysql.PostVotes{}
	bizpvs := &biz.PostVotes{}
	var err error

	re := regexp.MustCompile(`^(categories|tags)/(.+)/post_votes$`)
	x := re.FindStringSubmatch(parent)
	if len(x) != 3 {
		pvs, err = pvr.data.DBClient.DatabaseClient.QueryPostVote().All(ctx)
	}
	if err != nil {
		return nil, err
	}

	for _, p := range pvs.Collection {
		bizpvs.Collection = append(bizpvs.Collection, &biz.PostVote{
			PostVoteId: p.Id,
			PostId:     p.PostId,
			VoteId:     p.VoteId,
		})
	}
	return bizpvs, nil
}

func (pvr *postVoteRepo) GetPostVote(ctx context.Context, name string) (*biz.PostVote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^post_votes/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New(
			"GetPostVote: name: " + name + " cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id}
	p, err := pvr.data.DBClient.DatabaseClient.QueryPostVote().Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.PostVote{
		PostVoteId: p.Id,
		PostId:     p.PostId,
		VoteId:     p.VoteId,
	}, nil
}

func (pvr *postVoteRepo) GetByPidVid(ctx context.Context, name string) (*biz.PostVote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^post_votes/([\d.]+)/([\d.]+)/id$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 3 {
		return nil, errors.New(
			"GetByPidVid: name: " + name + " cannot match regex express")
	}
	pid, vid := x[1], x[2]
	clases := [][4]string{
		{"post_id", "=", pid, "and"},
		{"vote_id", "=", vid, "and"}}
	pv, err := pvr.data.DBClient.DatabaseClient.QueryPostVote().Where(clases...).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.PostVote{
		PostVoteId: pv.Id,
		PostId:     pv.PostId,
		VoteId:     pv.VoteId,
	}, nil
}

func (pvr *postVoteRepo) ListVidsByPid(ctx context.Context, name string) (*biz.PidVids, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^post_votes/([\d.]+)/list$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New(
			"ListVidsByPid: name: " + name + " cannot match regex express")
	}
	pid := x[1]
	clause := [4]string{"post_id", "=", pid}
	pvs, err := pvr.data.DBClient.DatabaseClient.QueryPostVote().
		Where(clause).All(ctx)
	if err != nil {
		return nil, err
	}
	iPid, err := strconv.Atoi(pid)
	if err != nil {
		log.Println(err)
	}
	bizps := &biz.PidVids{Pid: iPid}
	for _, p := range pvs.Collection {
		bizps.Vids = append(bizps.Vids, p.VoteId)
	}
	return bizps, nil
}

func (pvr *postVoteRepo) CreatePostVote(ctx context.Context, postVote *biz.PostVote) (*biz.PostVote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	clauses := [][4]string{
		{"post_id", "=", strconv.Itoa(postVote.PostId), "and"},
		{"vote_id", "=", strconv.Itoa(postVote.VoteId), "and"},
	}
	_, err := pvr.data.DBClient.DatabaseClient.QueryPostVote().
		Where(clauses...).First(ctx)
	if err != nil && errors.Is(err, mysql.ErrNotFound) {
		if err := pvr.data.DBClient.DatabaseClient.
			InsertPostVote(ctx, &mysql.PostVote{
				PostId: postVote.PostId,
				VoteId: postVote.VoteId,
			}); err != nil {
			return nil, err
		} else {
			return postVote, nil
		}
	}
	if err != nil {
		return nil, err
	}
	// err is nil means post_id and vote_id query back, it is exist.
	return nil, ErrRowExist
}

func (pvr *postVoteRepo) UpdatePostVote(ctx context.Context, postVote *biz.PostVote) (*biz.PostVote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbPostVote, err := pvr.data.DBClient.DatabaseClient.QueryPostVote().
		Where([4]string{"id", "=", strconv.Itoa(postVote.PostVoteId), "and"}).
		First(ctx)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(
			"postVoteRepo: UpdatePostVote: query postVote by id error: %v: %v",
			postVote, err))
	}
	if &postVote.PostId != nil {
		dbPostVote.PostId = postVote.PostId
	}
	if &postVote.VoteId != nil {
		dbPostVote.VoteId = postVote.VoteId
	}
	if err := pvr.data.DBClient.DatabaseClient.UpdatePostVote(ctx,
		dbPostVote); err != nil {
		return nil, err
	}
	return &biz.PostVote{
		PostVoteId: dbPostVote.Id,
		PostId:     dbPostVote.PostId,
		VoteId:     dbPostVote.VoteId,
	}, nil
}

func (pvr *postVoteRepo) DeletePostVote(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^post_votes/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("postVoteRepo: DeletePostVote: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("postVoteRepo: DeletePostVote: postVote id should be integer only")
	}
	return &emptypb.Empty{}, pvr.data.DBClient.DatabaseClient.DeletePostVote(ctx, id)
}
