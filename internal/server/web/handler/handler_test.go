package handler

import (
	"testing"

	"github.com/hi20160616/voter/configs"
)

var cfg = configs.NewConfig("hfcms")

func TestListPostsHandler(t *testing.T) {
	// req := httptest.NewRequest("GET", "http://localhost/list", nil)
	// w := httptest.NewRecorder()
	// p := &render.Page{Cfg: cfg, Title: "test list"}
	// // listPostsHandler(w, req, p)
	// resp := w.Result()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	//         t.Fatal(err)
	// }
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header.Get("Content-Type"))
	// fmt.Println(string(body))
}
