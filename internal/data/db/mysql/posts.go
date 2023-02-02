package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Post struct {
	Id, IsOpen             int
	Title, Detail          string
	CreateTime, UpdateTime time.Time
}

type Posts struct {
	Collection []*Post
}

type PostQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertPost(ctx context.Context, post *Post) error {
	q := `INSERT INTO posts(title, is_open, detail) VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE title=?, is_open=?, detail=?`
	aq := &PostQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, post.Title, post.IsOpen, post.Detail,
		post.Title, post.IsOpen, post.Detail)
	return errors.WithMessage(err, "mariadb: posts: Insert error")
}

func (dc *DatabaseClient) UpdatePost(ctx context.Context, post *Post) error {
	q := `UPDATE posts SET title=?, is_open=?, detail=?  WHERE id=?`
	uq := &PostQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query, post.Title, post.IsOpen, post.Detail, post.Id)
	return err
}

// DeletePost2 is true delete from database instead of DeletePost just update the row
func (dc *DatabaseClient) DeletePost(ctx context.Context, id int) error {
	q := `DELETE FROM posts WHERE id=?`
	aq := &PostQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryPost() *PostQuery {
	return &PostQuery{db: dc.db, query: `SELECT * FROM posts`}
}

// All will display all rows
func (aq *PostQuery) All(ctx context.Context) (*Posts, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := aq.db.Query("SELECT * FROM posts WHERE title like ?", "%%test%%")
	rows, err := aq.db.Query(aq.query, aq.args...)
	if err != nil {
		return nil, err
	}
	return mkPost(rows)
}

func (aq *PostQuery) First(ctx context.Context) (*Post, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// cs: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
// the last `or` or `and` in clause will cut off after prepareQuery().
// so, every clause need `or` or `and` for last element.
func (aq *PostQuery) Where(cs ...[4]string) *PostQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *PostQuery) Order(condition string) *PostQuery {
	aq.order = condition
	return aq
}

func (aq *PostQuery) Limit(limit int) *PostQuery {
	aq.limit = &limit
	return aq
}

func (aq *PostQuery) Offset(offset int) *PostQuery {
	aq.offset = &offset
	return aq
}

func (aq *PostQuery) prepareQuery(ctx context.Context) error {
	if aq.clauses != nil {
		aq.query += " WHERE "
		for i, c := range aq.clauses {
			// TODO: 2nd clause cannot be tied together automaticly
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(aq.clauses)-1 {
				aq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				aq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			aq.args = append(aq.args, c[2])
		}
	}
	if aq.order != "" {
		aq.query += " ORDER BY ?"
		aq.args = append(aq.args, aq.order)
	}
	if aq.limit != nil {
		aq.query += " LIMIT ?"
		a := strconv.Itoa(*aq.limit)
		aq.args = append(aq.args, a)
	}
	if aq.offset != nil {
		aq.query += ", ?"
		a := strconv.Itoa(*aq.offset)
		aq.args = append(aq.args, a)
	}
	return nil
}

func mkPost(rows *sql.Rows) (*Posts, error) {
	var title, detail sql.NullString
	var create_time, update_time sql.NullTime
	var id, is_open int
	var posts = &Posts{}
	for rows.Next() {
		if err := rows.Scan(&id, &title, &is_open, &detail, &create_time, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkPost rows.Scan error")
		}
		posts.Collection = append(posts.Collection, &Post{
			Id:         id,
			Title:      title.String,
			Detail:     detail.String,
			IsOpen:     is_open,
			CreateTime: create_time.Time,
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkPost error")
	}
	return posts, nil
}
