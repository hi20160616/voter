package handler

import (
	"net/http"

	"github.com/hi20160616/voter/internal/server/web/render"
)

func newVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "New Vote"
	render.Derive(w, "newvote", p)
}

func saveVoteHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "Save Vote"
	render.Derive(w, "savevote", p)
}
