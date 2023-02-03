package handler

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/configs"
	"github.com/hi20160616/voter/internal/server/render"
	"github.com/hi20160616/voter/internal/service"
	tmpl "github.com/hi20160616/voter/templates"
)

var validPath = regexp.MustCompile("^/(posts|votes|users|search)/(.*?)$")

// makeHandler invoke fn after path valided, and arrange args from url to object: `&render.Page{}`
func makeHandler(fn func(http.ResponseWriter, *http.Request, *render.Page), cfg *configs.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
		}
		fn(w, r, &render.Page{Cfg: cfg})
	}
}

// GetHandler is a handler merger and a router for mutipl handler
func GetHandler(cfg *configs.Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		homeHandler(w, req)
	})
	mux.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.FS(tmpl.FS))))
	mux.HandleFunc("/posts/", makeHandler(listPostsHandler, cfg))
	mux.HandleFunc("/posts/v", makeHandler(getPostHandler, cfg))
	mux.HandleFunc("/posts/s", makeHandler(searchPostsHandler, cfg))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render.Derive(w, "home", &render.Page{Title: "Home", Data: "need to be done"})
}

func listPostsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	ds, err := service.ListPosts(context.Background(), &pb.ListPostsRequest{}, p.Cfg)
	if err != nil {
		log.Println(err)
	}
	p.Data = ds.Posts
	p.Title = "Posts"
	render.Derive(w, "posts", p)
}

func getPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	id := r.URL.Query().Get("id")
	a, err := service.GetPost(context.Background(),
		&pb.GetPostRequest{Name: "posts/" + id}, p.Cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p.Data = a
	p.Title = a.Title
	render.Derive(w, "post", p) // template name: post
}

func searchPostsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	kws := r.URL.Query().Get("v")
	kws = strings.ReplaceAll(kws, " ", ",")
	as, err := service.SearchPosts(context.Background(), &pb.SearchPostsRequest{Name: "posts/" + kws + "/search"}, p.Cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p.Data = as
	render.Derive(w, "search", p)
}
