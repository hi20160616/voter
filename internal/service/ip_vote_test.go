package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	v1 "github.com/hi20160616/voter/api/voter/v1"
)

var ivs = func() *IpVoteService {
	ivs, err := NewIpVoteService()
	if err != nil {
		log.Fatal(err)
	}
	return ivs
}()

func TestCreateIpVote(t *testing.T) {

	a, err := ivs.CreateIpVote(context.Background(), &v1.CreateIpVoteRequest{
		IpVote: &v1.IpVote{
			Ip:       "127.0.0.1",
			VoteId:   1,
			Opts:     "11000000",
			TxtField: "test",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestListIpVotes(t *testing.T) {
	x, err := ivs.ListIpVotes(context.Background(), &v1.ListIpVotesRequest{})
	y, err := ivs.ListIpVotes(context.Background(), &v1.ListIpVotesRequest{
		Parent: "vote_id/2/ip_votes",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("vote_id=nil:::::::::::::::::::")
	for _, a := range x.IpVotes {
		fmt.Println(a)
	}
	fmt.Println("vote_id=2:::::::::::::::::::")
	for _, a := range y.IpVotes {
		fmt.Println(a)
	}
}

func TestGetIpVote(t *testing.T) {
	id := "3"
	x, err := ivs.GetIpVote(context.Background(), &v1.GetIpVoteRequest{Name: "ip_votes/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Ip: ", x.Ip)
	fmt.Println("IpVoteId: ", x.IpVoteId)
	fmt.Println("Opts: ", x.Opts)
	fmt.Println("TxtField: ", x.TxtField)
}

func TestUpdateIpVote(t *testing.T) {
	a, err := ivs.UpdateIpVote(context.Background(), &v1.UpdateIpVoteRequest{
		IpVote: &v1.IpVote{
			IpVoteId: 1,
			Ip:       "192.168.1.123",
			VoteId:   1,
			Opts:     "10100000",
			TxtField: "Yeah",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestDeleteIpVote(t *testing.T) {
	id := "15"
	name := "ip_votes/" + id + "/delete"
	if _, err := ivs.DeleteIpVote(context.Background(), &v1.DeleteIpVoteRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	_, err := ivs.GetIpVote(context.Background(), &v1.GetIpVoteRequest{Name: "ip_votes/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}
