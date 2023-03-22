package handler

import (
	"context"
	"fmt"
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

func getPostVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	id := r.URL.Query().Get("id")
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
	}

	post, err := ps.GetPost(context.Background(), &pb.GetPostRequest{Name: "posts/" + id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p.Data = post
	p.Title = post.Title
	render.Derive(w, "post", p) // template name: post
}

func savePostVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	id := r.URL.Query().Get("id")
	pvs, err := service.NewPostVoteService()
	if err != nil {
		log.Println(err)
	}

	if id == "" {
		// if is create a postVote
		pv, err := pvs.CreatePostVote(context.Background(), &pb.CreatePostVoteRequest{
			PostVote: &pb.PostVote{},
		})
		if err != nil {
			log.Println(err)
		}
		p.Data = &pb.PostVote{
			PostVoteId: pv.PostVoteId,
			PostId:     pv.PostId,
			VoteId:     pv.VoteId,
		}
		render.Derive(w, "post", p)
	} else {
		// edit a post
		pv, err := pvs.GetPostVote(context.Background(), &pb.GetPostVoteRequest{
			Name: "postvotes/" + id,
		})
		if err != nil {
			log.Println(err)
		}
		pv, err = pvs.UpdatePostVote(context.Background(), &pb.UpdatePostVoteRequest{
			PostVote: &pb.PostVote{
				PostVoteId: pv.PostVoteId,
				PostId:     pv.PostId,
				VoteId:     pv.VoteId,
			},
		})
		// add votes_post
		// get vids from form
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		vids := r.Form["SelectedVotes"]
		fmt.Println("vids: ", vids)

		p.Data = pv
		render.Derive(w, "post", p)
	}
}

func editPostVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	id := r.URL.Query().Get("id")
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
	}

	post, err := ps.GetPost(context.Background(), &pb.GetPostRequest{
		Name: "posts/" + id,
	})
	if err != nil {
		log.Println(err)
	}
	vs, err := service.NewVoteService()
	if err != nil {
		log.Println(err)
	}
	votes, err := vs.ListVotes(context.Background(), &pb.ListVotesRequest{})
	if err != nil {
		log.Println(err)
	}
	p.Data = struct {
		Post  *pb.Post
		Votes []*pb.Vote
	}{
		Post:  post,
		Votes: votes.Votes,
	}
	render.Derive(w, "editpost", p)
}
