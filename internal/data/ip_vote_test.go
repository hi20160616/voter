package data

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
)

var repo3 = func() biz.IpVoteRepo {
	dc, err := mysql.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewIpVoteRepo(&Data{DBClient: dc}, log.Default())
}()

func TestGetIpVote(t *testing.T) {
	x, err := repo3.GetIpVote(context.Background(), "ip_votes/3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Ip: ", x.Ip)
	fmt.Println("IpVoteId: ", x.IpVoteId)
	fmt.Println("Opts: ", x.Opts)
	fmt.Println("TxtField: ", x.TxtField)
}

func TestCreateIpVote(t *testing.T) {
	tcs := []*biz.IpVote{
		{Ip: "127.0.0.1", VoteId: 1, Opts: "00010000"},
		{Ip: "127.0.0.1", VoteId: 1, Opts: "00100000"},
		{Ip: "127.0.0.1", VoteId: 2, Opts: "01000000"},
		{Ip: "127.0.0.1", VoteId: 3, Opts: "11110000"},
	}

	for _, tc := range tcs {
		x, err := repo3.CreateIpVote(context.Background(), tc)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(x)
	}
}

func TestUpdateIpVote(t *testing.T) {
	x, err := repo3.UpdateIpVote(context.Background(), &biz.IpVote{
		IpVoteId: 1,
		Ip:       "12.12.12.123",
		VoteId:   1,
		Opts:     "01110000",
		TxtField: "Updated IpVote",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Ip: ", x.Ip)
	fmt.Println("IpVoteId: ", x.IpVoteId)
	fmt.Println("VoteId: ", x.VoteId)
	fmt.Println("Opts: ", x.Opts)
	fmt.Println("TxtField: ", x.TxtField)
}

// func TestListVidsByPid(t *testing.T) {
//         vids, err := repo2.ListVidsByPid(context.Background(), "post_votes/1/list")
//         if err != nil {
//                 t.Fatal(err)
//         }
//         fmt.Println("pid:", vids.Pid)
//         for _, e := range vids.Vids {
//                 fmt.Println(e)
//         }
// }
