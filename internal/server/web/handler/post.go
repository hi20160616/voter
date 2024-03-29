package handler

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/configs"
	"github.com/hi20160616/voter/internal/server/web/render"
	"github.com/hi20160616/voter/internal/service"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

type PostReport struct {
	VoteId                 string
	Vote                   *pb.Vote
	A, B, C, D, E, F, G, H int
	PercentA, PercentB     float32
	PercentC, PercentD     float32
	PercentE, PercentF     float32
	PercentG, PercentH     float32
	TxtFields              []string
}

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func IsAdminIp(ip string, cfg *configs.Config) bool {
	for _, e := range cfg.Manager.Admin {
		if e == ip {
			return true
		}
	}
	return false
}

func IsLeaderIp(ip string, cfg *configs.Config) bool {
	for _, e := range cfg.Manager.Leader {
		if e == ip {
			return true
		}
	}
	return false
}

func newPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
	p.Title = "New Post"
	render.Derive(w, "newpost", p)
}

func listPostsHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) && !IsLeaderIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
	}

	ds, err := ps.ListPosts(context.Background(), &pb.ListPostsRequest{})
	if err != nil {
		log.Println(err)
	}
	p.Data = ds.Posts
	p.Data = &struct {
		Posts []*pb.Post
	}{
		Posts: ds.Posts,
	}
	p.Title = "Posts"
	render.Derive(w, "posts", p)
}

func getPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	id := r.URL.Query().Get("id")
	p.ClientIP = RemoteIp(r)

	// Prejudge ip and post is not exist, otherwise, return warning page.
	ips, err := service.NewIpPostService()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}

	// Get posts by ip
	ip := RemoteIp(r)
	posts, err := ips.ListIpPosts(context.Background(), &pb.ListIpPostsRequest{
		Parent: "ip/" + ip + "/ip_posts",
	})
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}

	// Prejudge if the ip voted?
	voted := false
	for _, e := range posts.IpPosts {
		x, err := strconv.Atoi(id)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
		if e.PostId == int32(x) {
			// Voted, redirect to post_report
			voted = true
		}
	}

	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	post, err := ps.GetPost(context.Background(), &pb.GetPostRequest{
		Name: "posts/" + id})
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}

	vs, err := service.NewVoteService()
	if err != nil {
		log.Println(err)
	}
	votes, err := vs.ListVotes(context.Background(), &pb.ListVotesRequest{
		Parent: "pid/" + id + "/votes"})
	if err != nil {
		log.Println(err)
	}

	if voted {
		http.Redirect(w, r, "/posts/report/v?id="+id, 302)
	} else {
		p.Data = struct {
			Post  *pb.Post
			Votes []*pb.Vote
		}{
			Post:  post,
			Votes: votes.Votes,
		}
		p.Title = post.Title
		render.Derive(w, "post", p) // template name: post
	}
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
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
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
		http.Redirect(w, r, fmt.Sprintf("/posts/v?id=%d", post.PostId), 302)
	} else {
		// update a post
		pid, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}
		post, err := ps.UpdatePost(context.Background(), &pb.UpdatePostRequest{
			Post: &pb.Post{
				PostId:   int32(pid),
				Title:    title,
				IsClosed: int32(isClosed),
				Detail:   detail,
			},
		})
		// get vids as fVids from form
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		sv := r.Form["SelectedVotes"]
		fVids := []int{}
		for _, e := range sv {
			vid, err := strconv.Atoi(e)
			if err != nil {
				log.Println(err)
			}
			fVids = append(fVids, vid)
		}
		// get vids as dbVids from database
		pvs, err := service.NewPostVoteService()
		if err != nil {
			log.Println(err)
		}
		pbResponse, err := pvs.ListVidsByPid(context.Background(),
			&pb.ListVidsByPidRequest{Name: "post_votes/" + id + "/list"})
		if err != nil {
			log.Println(err)
		}

		dbVids := []int{}
		for _, e := range pbResponse.Vids {
			dbVids = append(dbVids, int(e))
		}

		// IntSliceDiff return []int in set1 but not in set2
		IntSliceDiff := func(set1, set2 []int) (ret []int) {
			for _, s1 := range set1 {
				i := 0
				for _, s2 := range set2 {
					if s1 == s2 {
						i++
						continue
					}
				}
				if i == 0 {
					ret = append(ret, s1)
				}
			}

			return
		}
		// Need add vids in form but not in database
		addVids := IntSliceDiff(fVids, dbVids)
		// Need del vids in database but not in form
		delVids := IntSliceDiff(dbVids, fVids)

		for _, e := range addVids {
			_, err = pvs.CreatePostVote(context.Background(),
				&pb.CreatePostVoteRequest{
					PostVote: &pb.PostVote{
						PostId: post.PostId,
						VoteId: int32(e),
					}})
		}

		for _, e := range delVids {
			post_vote, err := pvs.GetByPidVid(context.Background(),
				&pb.GetByPidVidRequest{
					Name: fmt.Sprintf("post_votes/%d/%d/id",
						post.PostId, e)})
			if err != nil {
				log.Println(err)
			}
			_, err = pvs.DeletePostVote(context.Background(),
				&pb.DeletePostVoteRequest{
					Name: fmt.Sprintf("post_votes/%d/delete",
						post_vote.PostVoteId)})
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, fmt.Sprintf("/posts/v?id=%d", post.PostId), 302)
	}
}

func editPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
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
	pvs, err := service.NewPostVoteService()
	lsVidByPid, err := pvs.ListVidsByPid(context.Background(),
		&pb.ListVidsByPidRequest{Name: "post_votes/" + id + "/list"})
	if err != nil {
		log.Println(err)
	}
	p.Data = struct {
		Post     *pb.Post
		Votes    []*pb.Vote
		PostVids []int32
	}{
		Post:     post,
		Votes:    votes.Votes,
		PostVids: lsVidByPid.Vids,
	}
	render.Derive(w, "editpost", p)
}

func votePostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	pid := r.URL.Query().Get("id")
	postid, err := strconv.Atoi(pid)
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p.ClientIP = RemoteIp(r)
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	// get vids in the post for options collect.
	pvs, err := service.NewPostVoteService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	postVotes, err := pvs.ListPostVotes(context.Background(),
		&pb.ListPostVotesRequest{Parent: "pid/" + pid + "/post_votes"})
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	vids := []int32{}
	for _, e := range postVotes.PostVotes {
		vids = append(vids, e.VoteId)
	}

	// collect votes
	ipvotes := &pb.IpVotes{}
	for _, e := range vids {
		sv := r.Form["vote"+strconv.Itoa(int(e))]
		optsArr := []byte{'0', '0', '0', '0', '0', '0', '0', '0'}
		for _, e := range sv {
			switch e {
			case "A":
				optsArr[0] = '1'
			case "B":
				optsArr[1] = '1'
			case "C":
				optsArr[2] = '1'
			case "D":
				optsArr[3] = '1'
			case "E":
				optsArr[4] = '1'
			case "F":
				optsArr[5] = '1'
			case "G":
				optsArr[6] = '1'
			case "H":
				optsArr[7] = '1'
			}
		}
		iv := &pb.IpVote{
			Ip:       RemoteIp(r),
			VoteId:   e,
			Opts:     bytes.NewBuffer(optsArr).String(),
			TxtField: r.FormValue("TxtField"),
			PostId:   int32(postid),
		}
		ipvotes.IpVotes = append(ipvotes.IpVotes, iv)
	}

	// insert or update ip_vote
	ivs, err := service.NewIpVoteService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, e := range ipvotes.IpVotes {
		// prejudge if exist of ip, post_id and vote_id at data/ip_vote.go
		// insert row or update if exist
		_, err := ivs.CreateIpVote(context.Background(), &pb.CreateIpVoteRequest{
			IpVote: e})
		if err != nil {
			log.Println(err)
		}
	}

	// insert or update ip_posts while vote success
	ips, err := service.NewIpPostService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = ips.CreateIpPost(context.Background(), &pb.CreateIpPostRequest{
		IpPost: &pb.IpPost{
			Ip:     RemoteIp(r),
			PostId: int32(postid),
		},
	})
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/posts/report/v?id="+pid, 302)
}

func postReportHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	pid := r.URL.Query().Get("id")
	p.ClientIP = RemoteIp(r)
	// get post info for page display
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	post, err := ps.GetPost(context.Background(), &pb.GetPostRequest{
		Name: "posts/" + pid})
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	vs, err := service.NewVoteService()
	votes, err := vs.ListVotes(context.Background(), &pb.ListVotesRequest{
		Parent: "pid/" + pid + "/votes"})
	if err != nil {
		log.Println("postReportHandler: vs.ListVotes: ", err)
	}
	prs := []*PostReport{}
	for _, e := range votes.Votes {
		pr := &PostReport{
			VoteId: strconv.Itoa(int(e.VoteId)),
			Vote:   e,
		}
		prs = append(prs, pr)
	}

	ivs, err := service.NewIpVoteService()
	if err != nil {
		log.Println("postReportHandler: NewIpVoteService: ", err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, e := range prs {
		ipvotes, err := ivs.ListIpVotes(context.Background(),
			&pb.ListIpVotesRequest{Parent: fmt.Sprintf(
				"post_id/%s/vote_id/%s/ip_votes",
				pid, e.VoteId),
			})
		if err != nil {
			log.Println("postReportHandler: ivs.ListIpVotes: ", err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if len(ipvotes.IpVotes) == 0 {
			log.Println("postReportHandler: ivs.ListIpVotes err: "+
				"len(ipvotes.IpVotes) == 0: ", err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// e.TxtFields = append(e.TxtFields, ipvotes.IpVotes)
		for _, e1 := range ipvotes.IpVotes {
			if e1.Opts[0] == '1' {
				e.A++
			}
			if e1.Opts[1] == '1' {
				e.B++
			}
			if e1.Opts[2] == '1' {
				e.C++
			}
			if e1.Opts[3] == '1' {
				e.D++
			}
			if e1.Opts[4] == '1' {
				e.E++
			}
			if e1.Opts[5] == '1' {
				e.F++
			}
			if e1.Opts[6] == '1' {
				e.G++
			}
			if e1.Opts[7] == '1' {
				e.H++
			}
			if strings.TrimSpace(e1.TxtField) != "" {
				e.TxtFields = append(e.TxtFields, e1.TxtField)
			}
		}
		dividend := float32(e.A + e.B + e.C + e.D + e.E + e.F + e.G + e.H)
		e.PercentA = float32(e.A) * 100 / dividend
		e.PercentB = float32(e.B) * 100 / dividend
		e.PercentC = float32(e.C) * 100 / dividend
		e.PercentD = float32(e.D) * 100 / dividend
		e.PercentE = float32(e.E) * 100 / dividend
		e.PercentF = float32(e.F) * 100 / dividend
		e.PercentG = float32(e.G) * 100 / dividend
		e.PercentH = float32(e.H) * 100 / dividend
	}
	p.Data = struct {
		Post        *pb.Post
		PostReports []*PostReport
	}{
		Post:        post,
		PostReports: prs,
	}
	p.Title = "Post Report"
	render.Derive(w, "post_report", p)
}

func delPostHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// # 0. prejudge ip is allowed
	p.ClientIP = RemoteIp(r)
	if !IsAdminIp(p.ClientIP, p.Cfg) {
		p.Title = "404"
		render.Derive(w, "404", p)
		return
	}
	pid := r.URL.Query().Get("id")

	// # 1. Delete post
	ps, err := service.NewPostService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = ps.DeletePost(context.Background(), &pb.DeletePostRequest{
		Name: "posts/" + pid + "/delete",
	})
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// # 2. Delete from ip_posts by pid
	ips, err := service.NewIpPostService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	a, err := ips.ListIpPosts(context.Background(), &pb.ListIpPostsRequest{
		Parent: fmt.Sprintf("ip_posts/%s", pid),
	})
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, e := range a.IpPosts {
		_, err = ips.DeleteIpPost(context.Background(), &pb.DeleteIpPostRequest{
			Name: fmt.Sprintf("ip_posts/%d/delete", e.IpPostId),
		})
		if err != nil {
			log.Println(err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// # 3. Gether vids by pid from post_votes
	pvs, err := service.NewPostVoteService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	postvotes, err := pvs.ListPostVotes(context.Background(), &pb.ListPostVotesRequest{
		Parent: "pid/" + pid + "/post_votes",
	})
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	vids := []string{}
	for _, e := range postvotes.PostVotes {
		vids = append(vids, strconv.Itoa(int(e.VoteId)))
	}
	ivs, err := service.NewIpVoteService()
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// # 4. Delete from post_votes
	// ## 4.1. by pid
	y, err := pvs.ListPostVotes(context.Background(), &pb.ListPostVotesRequest{
		Parent: fmt.Sprintf("pid/%s/post_votes", pid),
	})
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, e := range y.PostVotes {
		_, err = pvs.DeletePostVote(context.Background(), &pb.DeletePostVoteRequest{
			Name: fmt.Sprintf("post_votes/%d/delete", e.PostVoteId),
		})
		if err != nil {
			log.Println(err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	// ## 4.2. by vids
	for _, e := range vids {
		z, err := pvs.ListPostVotes(context.Background(), &pb.ListPostVotesRequest{
			Parent: fmt.Sprintf("vid/%s/post_votes", e),
		})
		if err != nil {
			log.Println(err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		for _, e := range z.PostVotes {
			_, err := pvs.DeletePostVote(context.Background(),
				&pb.DeletePostVoteRequest{
					Name: fmt.Sprintf("post_votes/%d/delete",
						e.PostVoteId),
				})
			if err != nil {
				log.Println(err)
				// http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
	// # 5. Delete from ip_votes by vids
	for _, e := range vids {
		x, err := ivs.ListIpVotes(context.Background(), &pb.ListIpVotesRequest{
			Parent: "vote_id/" + e + "/ip_votes",
		})
		for _, e := range x.IpVotes {
			_, err = ivs.DeleteIpVote(context.Background(),
				&pb.DeleteIpVoteRequest{
					Name: fmt.Sprintf("ip_votes/%d/delete", e.IpVoteId),
				})
			if err != nil {
				log.Println(err)
				// http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		}
	}
	// # 6. Delete from votes by vids
	for _, e := range vids {
		vs, err := service.NewVoteService()
		if err != nil {
			log.Println(err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = vs.DeleteVote(context.Background(), &pb.DeleteVoteRequest{
			Name: fmt.Sprintf("votes/%s/delete", e),
		})
		if err != nil {
			log.Println(err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	http.Redirect(w, r, "/posts/", 302)
}
