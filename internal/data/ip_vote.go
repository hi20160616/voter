package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"time"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ biz.IpVoteRepo = new(ipVoteRepo)

type ipVoteRepo struct {
	data *Data
	log  *log.Logger
}

func NewIpVoteRepo(data *Data, logger *log.Logger) biz.IpVoteRepo {
	return &ipVoteRepo{
		data: data,
		log:  log.Default(),
	}
}

func (ivr *ipVoteRepo) ListIpVotes(ctx context.Context, parent string) (*biz.IpVotes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	ivs := &mysql.IpVotes{}
	bizivs := &biz.IpVotes{}
	var err error

	re := regexp.MustCompile(`^(vote_id)/(.+)/ip_votes$`)
	x := re.FindStringSubmatch(parent)
	if len(x) != 3 {
		ivs, err = ivr.data.DBClient.DatabaseClient.QueryIpVote().All(ctx)
	} else {
		ivs, err = ivr.data.DBClient.DatabaseClient.QueryIpVote().
			Where([4]string{"vote_id", "=", x[2]}).All(ctx)
	}

	if err != nil {
		return nil, err
	}

	for _, p := range ivs.Collection {
		bizivs.Collection = append(bizivs.Collection, &biz.IpVote{
			IpVoteId: p.Id,
			Ip:       p.Ip,
			VoteId:   p.VoteId,
			Opts:     p.Opts,
			TxtField: p.TxtField,
		})
	}
	return bizivs, nil
}

func (ivr *ipVoteRepo) GetIpVote(ctx context.Context, name string) (*biz.IpVote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^ip_votes/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New(
			"GetIpVote: name: " + name + " cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id}
	p, err := ivr.data.DBClient.DatabaseClient.QueryIpVote().Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.IpVote{
		IpVoteId: p.Id,
		Ip:       p.Ip,
		VoteId:   p.VoteId,
		Opts:     p.Opts,
		TxtField: p.TxtField,
	}, nil
}

// inet_aton convert ip,
// If the given input IPv4 address is a.b.c.d then
// return string of value a×2563+ b×2562+ c×2561 + d
func inet_aton(ip string) string {
	var ip_int uint32
	ip_byte := net.ParseIP(ip).To4()
	for i := 0; i < len(ip_byte); i++ {
		ip_int |= uint32(ip_byte[i])
		if i < 3 {
			ip_int <<= 8
		}
	}
	return strconv.Itoa(int(ip_int))
}

// CreateIpVote insert a row to database if ip and vote not exist, otherwise, update this row
func (ivr *ipVoteRepo) CreateIpVote(ctx context.Context, ipVote *biz.IpVote) (*biz.IpVote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	ip := inet_aton(ipVote.Ip)
	clauses := [][4]string{
		{"ip", "=", ip, "and"},
		{"vote_id", "=", strconv.Itoa(ipVote.VoteId), "and"},
	}

	_, err := ivr.data.DBClient.DatabaseClient.QueryIpVote().
		Where(clauses...).First(ctx)
	if err != nil {
		if errors.Is(err, mysql.ErrNotFound) {
			ipvoteid, err := ivr.data.DBClient.DatabaseClient.
				InsertIpVote(ctx, &mysql.IpVote{
					Ip:       ipVote.Ip,
					VoteId:   ipVote.VoteId,
					Opts:     ipVote.Opts,
					TxtField: ipVote.TxtField,
				})
			if err != nil {
				return nil, err
			} else {
				ipVote.IpVoteId = int(ipvoteid)
				return ipVote, nil
			}
		} else {
			return nil, err
		}
	}
	// err is nil means ip_id and vote_id query back, it is exist.
	// update the row
	return ivr.UpdateIpVote(ctx, ipVote)
}

func (ivr *ipVoteRepo) UpdateIpVote(ctx context.Context, ipVote *biz.IpVote) (*biz.IpVote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbIpVote := &mysql.IpVote{}
	var err error
	if ipVote.IpVoteId == 0 {
		ip := inet_aton(ipVote.Ip)
		dbIpVote, err = ivr.data.DBClient.DatabaseClient.QueryIpVote().Where(
			[][4]string{
				{"ip", "=", ip, "and"},
				{"vote_id", "=", strconv.Itoa(ipVote.VoteId), "and"},
			}...).First(ctx)
		if err != nil {
			return nil, errors.Join(fmt.Errorf(
				"ipVoteRepo: UpdateIpVote: query ipVote by ip"+
					"and vote_id error: %v: %v", ipVote, err))
		}
	} else {
		dbIpVote, err = ivr.data.DBClient.DatabaseClient.QueryIpVote().Where(
			[4]string{"id", "=", strconv.Itoa(ipVote.IpVoteId), "and"},
		).First(ctx)
		if err != nil {
			return nil, errors.Join(fmt.Errorf(
				"ipVoteRepo: UpdateIpVote: query ipVote by id error: %v: %v",
				ipVote, err))
		}
	}
	if &ipVote.Ip != nil {
		dbIpVote.Ip = ipVote.Ip
	}
	if &ipVote.VoteId != nil {
		dbIpVote.VoteId = ipVote.VoteId
	}
	if &ipVote.Opts != nil {
		dbIpVote.Opts = ipVote.Opts
	}
	if &ipVote.TxtField != nil {
		dbIpVote.TxtField = ipVote.TxtField
	}
	if err := ivr.data.DBClient.DatabaseClient.UpdateIpVote(ctx,
		dbIpVote); err != nil {
		return nil, err
	}
	return &biz.IpVote{
		IpVoteId: dbIpVote.Id,
		Ip:       dbIpVote.Ip,
		VoteId:   dbIpVote.VoteId,
		Opts:     dbIpVote.Opts,
		TxtField: dbIpVote.TxtField,
	}, nil
}

func (ivr *ipVoteRepo) DeleteIpVote(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^ip_votes/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("ipVoteRepo: DeleteIpVote: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("ipVoteRepo: DeleteIpVote: ipVote id should be integer only")
	}
	return &emptypb.Empty{}, ivr.data.DBClient.DatabaseClient.DeleteIpVote(ctx, id)
}
