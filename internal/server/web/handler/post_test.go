package handler

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/hi20160616/voter/internal/server/web/render"
)

func TestGetPostHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/pid/6/votes", nil)
	w := httptest.NewRecorder()
	p := &render.Page{Cfg: cfg, Title: "test list"}
	getPostHandler(w, req, p)
	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
