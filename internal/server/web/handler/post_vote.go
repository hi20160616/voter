package handler

import (
	"context"
	"log"
	"net/http"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/internal/server/web/render"
	"github.com/hi20160616/voter/internal/service"
)

func newPostVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "New PostVote"
	render.Derive(w, "newpostvote", p)
}

func listPostVotesHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	pvs, err := service.NewPostVoteService()
	if err != nil {
		log.Println(err)
	}

	ds, err := pvs.ListPostVotes(context.Background(), &pb.ListPostVotesRequest{})
	if err != nil {
		log.Println(err)
	}
	p.Data = ds.PostVotes
	p.Title = "Post votes"
	render.Derive(w, "posts", p)
}
