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

var _ biz.VoteRepo = new(voteRepo)

type voteRepo struct {
	data *Data
	log  *log.Logger
}

func NewVoteRepo(data *Data, logger *log.Logger) biz.VoteRepo {
	return &voteRepo{
		data: data,
		log:  log.Default(),
	}
}

// parent=pid/*/votes
func (vr *voteRepo) ListVotes(ctx context.Context, parent string) (*biz.Votes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	bizvs := &biz.Votes{}
	vs := &mysql.Votes{}
	var err error
	// reserved cases
	// re := regexp.MustCompile(`^(departments|roles|votegroups|pid)/(.+)/votes$`)
	re := regexp.MustCompile(`^(pid)/(.+)/votes$`)
	x := re.FindStringSubmatch(parent)
	y, err := regexp.MatchString(parent, `^votes$`)
	if err != nil {
		return nil, err
	}
	if len(x) != 3 && y {
		vs, err = vr.data.DBClient.DatabaseClient.QueryVote().All(ctx)
	} else {
		clause := [4]string{}
		clauses := [][4]string{}
		switch x[1] {
		case "pid":
			clause = [4]string{"post_id", "=", x[2], "or"}
			pvs, err := vr.data.DBClient.DatabaseClient.QueryPostVote().
				Where(clause).All(ctx)
			if err != nil {
				return nil, err
			}
			for _, e := range pvs.Collection {
				clauses = append(clauses, [4]string{"id", "=",
					strconv.Itoa(e.VoteId), "or"})
			}
		}
		vs, err = vr.data.DBClient.DatabaseClient.QueryVote().
			Where(clauses...).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, v := range vs.Collection {
		bizvs.Collection = append(bizvs.Collection, &biz.Vote{
			VoteId:      v.Id,
			Title:       v.Title,
			IsRadio:     v.IsRadio,
			A:           v.A,
			B:           v.B,
			C:           v.C,
			D:           v.D,
			E:           v.E,
			F:           v.F,
			G:           v.G,
			H:           v.H,
			HasTxtField: v.HasTxtField,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		})
	}
	return bizvs, nil
}

func (vr *voteRepo) ListVotesByPid(ctx context.Context, parent string) (*biz.Votes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	bizvs := &biz.Votes{}
	vs := &mysql.Votes{}
	var err error
	// reserved cases
	// re := regexp.MustCompile(`^(departments|roles|votegroups|pid)/(.+)/votes$`)
	re := regexp.MustCompile(`^(pid)/(.+)/votes$`)
	x := re.FindStringSubmatch(parent)
	y, err := regexp.MatchString(parent, `^votes$`)
	if err != nil {
		return nil, err
	}
	if len(x) != 3 && y {
		vs, err = vr.data.DBClient.DatabaseClient.QueryVote().All(ctx)
	} else {
		clause := [4]string{}
		switch x[1] {
		// reserved cases
		// case "departments":
		//         clause = [4]string{"department_id", "=", x[2], "and"}
		// case "roles":
		//         clause = [4]string{"role_id", "=", x[2], "and"}
		// case "votegroups":
		//         clause = [4]string{"votegroup_id", "=", x[2], "and"}
		case "pid":
			clause = [4]string{"post_id", "=", x[2], "and"}
		}
		vs, err = vr.data.DBClient.DatabaseClient.QueryVote().
			Where(clause).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, v := range vs.Collection {
		bizvs.Collection = append(bizvs.Collection, &biz.Vote{
			VoteId:      v.Id,
			Title:       v.Title,
			IsRadio:     v.IsRadio,
			A:           v.A,
			B:           v.B,
			C:           v.C,
			D:           v.D,
			E:           v.E,
			F:           v.F,
			G:           v.G,
			H:           v.H,
			HasTxtField: v.HasTxtField,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		})
	}
	return bizvs, nil
}

func (vr *voteRepo) GetVote(ctx context.Context, name string) (*biz.Vote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// name=votes/1
	re := regexp.MustCompile(`^votes/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id, "and"}
	v, err := vr.data.DBClient.DatabaseClient.QueryVote().
		Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Vote{
		VoteId:      v.Id,
		Title:       v.Title,
		IsRadio:     v.IsRadio,
		A:           v.A,
		B:           v.B,
		C:           v.C,
		D:           v.D,
		E:           v.E,
		F:           v.F,
		G:           v.G,
		H:           v.H,
		HasTxtField: v.HasTxtField,
		CreateTime:  v.CreateTime,
		UpdateTime:  v.UpdateTime,
	}, nil
}

func (vr *voteRepo) SearchVotes(ctx context.Context, name string) (*biz.Votes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^votes/(.+)/search$`)
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
			[4]string{"a", "like", kw, "or"},
			[4]string{"b", "like", kw, "or"},
			[4]string{"c", "like", kw, "and"},
			[4]string{"d", "like", kw, "and"},
			[4]string{"e", "like", kw, "and"},
			[4]string{"f", "like", kw, "and"},
			[4]string{"g", "like", kw, "and"},
			[4]string{"h", "like", kw, "and"},
		)
	}
	vs, err := vr.data.DBClient.DatabaseClient.QueryVote().
		Where(cs...).All(ctx)
	if err != nil {
		return nil, err
	}
	bizvs := &biz.Votes{Collection: []*biz.Vote{}}
	for _, v := range vs.Collection {
		bizvs.Collection = append(bizvs.Collection, &biz.Vote{
			VoteId:      v.Id,
			Title:       v.Title,
			IsRadio:     v.IsRadio,
			A:           v.A,
			B:           v.B,
			C:           v.C,
			D:           v.D,
			E:           v.E,
			F:           v.F,
			G:           v.G,
			H:           v.H,
			HasTxtField: v.HasTxtField,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
		})
	}
	return bizvs, nil
}

func (vr *voteRepo) CreateVote(ctx context.Context, vote *biz.Vote) (*biz.Vote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := vr.data.DBClient.DatabaseClient.
		InsertVote(ctx, &mysql.Vote{
			Title:       vote.Title,
			IsRadio:     vote.IsRadio,
			A:           vote.A,
			B:           vote.B,
			C:           vote.C,
			D:           vote.D,
			E:           vote.E,
			F:           vote.F,
			G:           vote.G,
			H:           vote.H,
			HasTxtField: vote.HasTxtField,
			CreateTime:  vote.CreateTime,
			UpdateTime:  vote.UpdateTime,
		}); err != nil {
		return nil, err
	}
	return vote, nil
}

func (vr *voteRepo) UpdateVote(ctx context.Context, vote *biz.Vote) (*biz.Vote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbVote, err := vr.data.DBClient.DatabaseClient.QueryVote().
		Where([4]string{"id", "=", strconv.Itoa(vote.VoteId), "and"}).
		First(ctx)
	if err != nil {
		return nil, errors.Join(fmt.Errorf(
			"voteRepo: UpdateVote: query vote by id error: %v: %v",
			vote, err))
	}
	if &vote.Title != nil {
		dbVote.Title = vote.Title
	}
	if &vote.IsRadio != nil {
		dbVote.IsRadio = vote.IsRadio
	}
	if &vote.A != nil {
		dbVote.A = vote.A
	}
	if &vote.B != nil {
		dbVote.B = vote.B
	}
	if &vote.C != nil {
		dbVote.C = vote.C
	}
	if &vote.D != nil {
		dbVote.D = vote.D
	}
	if &vote.E != nil {
		dbVote.E = vote.E
	}
	if &vote.F != nil {
		dbVote.F = vote.F
	}
	if &vote.G != nil {
		dbVote.G = vote.G
	}
	if &vote.H != nil {
		dbVote.H = vote.H
	}
	if &vote.HasTxtField != nil {
		dbVote.HasTxtField = vote.HasTxtField
	}
	if err := vr.data.DBClient.DatabaseClient.
		UpdateVote(ctx, dbVote); err != nil {
		return nil, err
	}
	return &biz.Vote{
		VoteId:      dbVote.Id,
		Title:       dbVote.Title,
		IsRadio:     dbVote.IsRadio,
		A:           dbVote.A,
		B:           dbVote.B,
		C:           dbVote.C,
		D:           dbVote.D,
		E:           dbVote.E,
		F:           dbVote.F,
		G:           dbVote.G,
		H:           dbVote.H,
		HasTxtField: dbVote.HasTxtField,
		CreateTime:  dbVote.CreateTime,
		UpdateTime:  dbVote.UpdateTime,
	}, nil
}

// DeleteVote is soft delete, that can be undeleted, it just update deleted field to 1
func (vr *voteRepo) DeleteVote(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^votes/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New(
				"voteRepo: DeleteVote: name cannot match regex express: " + name)
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("voteRepo: DeleteVote: vote id should be integer only")
	}
	return &emptypb.Empty{}, vr.data.DBClient.DatabaseClient.DeleteVote(ctx, id)
}
