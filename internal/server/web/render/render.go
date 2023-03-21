package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/hi20160616/voter/configs"
	"github.com/yuin/goldmark"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Page struct {
	Title string
	Data  interface{}
	Cfg   *configs.Config
}

var tmpl = template.New("")

func init() {
	tmpl.Funcs(template.FuncMap{
		"summary":      summary,
		"smartTime":    smartTime,
		"smartDate":    smartDate,
		"markdown":     markdown,
		"unescapeHTML": unescapeHTML,
		"plusOne":      plusOne,
		"isClosed":     isClosed,
		"if1Checked":   if1Checked,
	})
	// templates = template.Must(templates.ParseFS(tmpl.FS, "default/*.html"))
	pattern := filepath.Join("templates", "default", "*.html")
	tmpl = template.Must(tmpl.ParseGlob(pattern))
}

func Derive(w http.ResponseWriter, tmplName string, p *Page) {
	if err := tmpl.ExecuteTemplate(w, tmplName+".html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("err template: %s.html\n\terror: %#v", tmplName, err)
	}
}

func summary(des string) string {
	dRune := []rune(des)
	if len(dRune) <= 300 {
		return des
	}
	return string(dRune[:300])
}

func parseWithZone(t time.Time) time.Time {
	loc := time.FixedZone("UTC", 8*60*60)
	return t.In(loc)

}

func smartTime(t *timestamppb.Timestamp, site string) string {
	if site == "cna" ||
		site == "dw" ||
		site == "kabar" ||
		site == "ucpnz" ||
		site == "kyodonews" {
		return t.AsTime().Format("15:04")
	}
	return parseWithZone(t.AsTime()).Format("15:04")
}

func smartDate(t *timestamppb.Timestamp, site string) string {
	if site == "cna" ||
		site == "dw" ||
		site == "kabar" ||
		site == "ucpnz" ||
		site == "kyodonews" {
		return t.AsTime().Format("01.02")
	}
	return parseWithZone(t.AsTime()).Format("01.02")
}

func markdown(in string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(in), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}

func plusOne(x int) int {
	return x + 1
}

func isClosed(x int32) string {
	if x == 1 {
		return "**Closed!**"
	}
	return "Is Open!"
}

func if1Checked(x int32) string {
	if x == 1 {
		return "checked"
	}
	return ""
}
