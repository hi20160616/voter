package handler

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/configs"
	"github.com/hi20160616/voter/internal/server/web/render"
	"github.com/hi20160616/voter/internal/service"
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
	mux.HandleFunc("/s/", serveResource) // `/s` means source
	mux.HandleFunc("/posts/", makeHandler(listPostsHandler, cfg))
	mux.HandleFunc("/posts/v", makeHandler(getPostHandler, cfg))
	mux.HandleFunc("/posts/new", makeHandler(newPostHandler, cfg))
	mux.HandleFunc("/votes/", makeHandler(listVotesHandler, cfg))
	mux.HandleFunc("/votes/v", makeHandler(getVoteHandler, cfg))
	mux.HandleFunc("/votes/new", makeHandler(newVoteHandler, cfg))
	mux.HandleFunc("/votes/save", makeHandler(saveVoteHandler, cfg))
	// mux.HandleFunc("/posts/s", makeHandler(searchPostsHandler, cfg))
	return mux
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "templates" + req.URL.Path[2:] // filter `/s`
	var contentType string
	switch path != "" {
	case strings.HasSuffix(path, ".css"):
		contentType = "text/css"
	case strings.HasSuffix(path, ".js"):
		contentType = "text/javascript"
	case strings.HasSuffix(path, ".png"):
		contentType = "image/png"
	case strings.HasSuffix(path, ".gif"):
		contentType = "image/gif"
	case strings.HasSuffix(path, ".jpg"):
		contentType = "image/jpg"
	default:
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
	}

	ds, err := ps.ListPosts(context.Background(), &pb.ListPostsRequest{})
	if err != nil {
		log.Println(err)
	}
	render.Derive(w, "home", &render.Page{Title: "Home", Data: ds.Posts})
}
