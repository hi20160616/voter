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
