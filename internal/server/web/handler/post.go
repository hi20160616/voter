package handler

import (
	"context"
	"log"
	"net/http"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/internal/server/web/render"
	"github.com/hi20160616/voter/internal/service"
)

func newPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.Title = "New Post"
	render.Derive(w, "newpost", p)
}

func listPostsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
	}

	ds, err := ps.ListPosts(context.Background(), &pb.ListPostsRequest{})
	if err != nil {
		log.Println(err)
	}
	p.Data = ds.Posts
	p.Title = "Posts"
	render.Derive(w, "posts", p)
}

func getPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
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

// func searchPostsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
//         kws := r.URL.Query().Get("v")
//         kws = strings.ReplaceAll(kws, " ", ",")
//         as, err := service.SearchPosts(context.Background(), &pb.SearchPostsRequest{Name: "posts/" + kws + "/search"}, p.Cfg)
//         if err != nil {
//                 http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//         p.Data = as
//         render.Derive(w, "search", p)
// }

func savePostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	id := r.URL.Query().Get("id")
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
	}

	isClosed := 0
	if v := r.FormValue("IsClosed"); v == "1" {
		isClosed = 1
	}
	title := r.FormValue("PostTitle")
	detail := r.FormValue("PostDetail")
	if id == "" {
		// if is create a post
		post, err := ps.CreatePost(context.Background(), &pb.CreatePostRequest{
			Post: &pb.Post{
				Title:    title,
				IsClosed: int32(isClosed),
				Detail:   detail,
			},
		})
		if err != nil {
			log.Println(err)
		}
		p.Data = post
		render.Derive(w, "post", p)
	} else {
		// edit a post
		post, err := ps.GetPost(context.Background(), &pb.GetPostRequest{
			Name: "posts/" + id,
		})
		if err != nil {
			log.Println(err)
		}
		post, err = ps.UpdatePost(context.Background(), &pb.UpdatePostRequest{
			Post: &pb.Post{
				PostId:   post.PostId,
				Title:    title,
				IsClosed: int32(isClosed),
				Detail:   detail,
			},
		})
		p.Data = post
		render.Derive(w, "post", p)
	}
}

func editPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
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
	p.Data = post
	render.Derive(w, "editpost", p)
}
