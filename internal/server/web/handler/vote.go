package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/internal/server/web/render"
	"github.com/hi20160616/voter/internal/service"
)

func newVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
	p.Title = "New Vote"
	render.Derive(w, "newvote", p)
}

func saveVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
	id := r.URL.Query().Get("id")
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
	if id == "" {
		v, err := vs.CreateVote(context.Background(), &pb.CreateVoteRequest{
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
		p.Data = struct{ Vote *pb.Vote }{Vote: v}
		p.Title = "Continute add vote"
		render.Derive(w, "newvote", p)
	} else {
		// update a vote
		vid, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}
		v, err := vs.UpdateVote(context.Background(), &pb.UpdateVoteRequest{
			Vote: &pb.Vote{
				VoteId:      int32(vid),
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
		p.Data = struct{ Vote *pb.Vote }{Vote: v}
		p.Title = "Updated vote done."
		render.Derive(w, "vote", p)
	}

}

func listVotesHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
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
	p.Data = struct{ Vote *pb.Vote }{Vote: vote}
	p.Title = vote.Title
	render.Derive(w, "vote", p) // template name: vote
}

func editVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
	id := r.URL.Query().Get("id")
	vs, err := service.NewVoteService()
	if err != nil {
		log.Println(err)
	}

	vote, err := vs.GetVote(context.Background(), &pb.GetVoteRequest{
		Name: "votes/" + id,
	})
	if err != nil {
		log.Println(err)
	}
	p.Data = struct{ Vote *pb.Vote }{Vote: vote}
	render.Derive(w, "editvote", p)
}

func delVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
	id := r.URL.Query().Get("id")
	vs, err := service.NewVoteService()
	if err != nil {
		log.Println(err)
	}

	_, err = vs.DeleteVote(context.Background(), &pb.DeleteVoteRequest{
		Name: "votes/" + id + "/delete",
	})
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
