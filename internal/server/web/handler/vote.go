package handler

import (
	"context"
	"log"
	"net/http"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/internal/server/web/render"
	"github.com/hi20160616/voter/internal/service"
)

func newVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "New Vote"
	render.Derive(w, "newvote", p)
}

func saveVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	vs, err := service.NewVoteService()
	if err != nil {
		log.Println(err)
	}

	isRadio := 0
	if v := r.FormValue("VoteType"); v == "CheckBox" {
		isRadio = 1
	}
	hasTxtField := 0
	if v := r.FormValue("HasTxtField"); v == "Yes" {
		hasTxtField = 1
	}

	_, err = vs.CreateVote(context.Background(), &pb.CreateVoteRequest{
		Vote: &pb.Vote{
			Title:       r.FormValue("VoteTitle"),
			IsRadio:     int32(isRadio),
			A:           r.FormValue("A"),
			B:           r.FormValue("B"),
			C:           r.FormValue("C"),
			D:           r.FormValue("D"),
			E:           r.FormValue("E"),
			F:           r.FormValue("F"),
			G:           r.FormValue("G"),
			H:           r.FormValue("H"),
			HasTxtField: int32(hasTxtField),
		},
	})
	if err != nil {
		log.Println(err)
	}
	p.Title = "Continute add vote"
	render.Derive(w, "newvote", p)
}

func listVotesHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	vs, err := service.NewVoteService()
	if err != nil {
		log.Println(err)
	}

	ds, err := vs.ListVotes(context.Background(), &pb.ListVotesRequest{})
	if err != nil {
		log.Println(err)
	}
	p.Data = ds.Votes
	p.Title = "Votes"
	render.Derive(w, "votes", p)
}

func getVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	id := r.URL.Query().Get("id")
	ps, err := service.NewVoteService()
	if err != nil {
		log.Println(err)
	}

	vote, err := ps.GetVote(context.Background(), &pb.GetVoteRequest{Name: "votes/" + id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p.Data = vote
	p.Title = vote.Title
	render.Derive(w, "vote", p) // template name: vote
}
