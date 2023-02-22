package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	v1 "github.com/hi20160616/voter/api/voter/v1"
)

var vs = func() *VoteService {
	vs, err := NewVoteService()
	if err != nil {
		log.Fatal(err)
	}
	return vs
}()

func TestCreateVote(t *testing.T) {

	a, err := vs.CreateVote(context.Background(), &v1.CreateVoteRequest{
		Vote: &v1.Vote{
			Title:   "Richael",
			IsRadio: 1,
			A:       "a. create vote detail",
			B:       "b. create vote detail",
			C:       "c. create vote detail",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestListVotes(t *testing.T) {
	vs, err := vs.ListVotes(context.Background(), &v1.ListVotesRequest{})
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range vs.Votes {
		fmt.Println(a)
	}
}

func TestGetVote(t *testing.T) {
	id := "1"
	u, err := vs.GetVote(context.Background(), &v1.GetVoteRequest{Name: "votes/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("vote Title: ", u.Title)
}

func TestSearchVotes(t *testing.T) {
	name := "votes/vote1/search"
	votes, err := vs.SearchVotes(context.Background(), &v1.SearchVotesRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range votes.Votes {
		fmt.Println(v)
	}
}

func TestUpdateVote(t *testing.T) {
	u, err := vs.UpdateVote(context.Background(), &v1.UpdateVoteRequest{
		Vote: &v1.Vote{
			VoteId:  1,
			Title:   "Updated vote title",
			IsRadio: 0,
			A:       "a. Updated vote detail",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func TestDeleteVote(t *testing.T) {
	id := "1"
	name := "votes/" + id + "/delete"
	if _, err := vs.DeleteVote(context.Background(),
		&v1.DeleteVoteRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	_, err := vs.GetVote(context.Background(), &v1.GetVoteRequest{Name: "votes/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}
