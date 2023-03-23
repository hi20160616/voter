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
	Id, IsClosed           int
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

func (dc *DatabaseClient) InsertPost(ctx context.Context, post *Post) (int64, error) {
	q := `INSERT INTO posts(title, is_closed, detail) VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE title=?, is_closed=?, detail=?`
	pq := &PostQuery{db: dc.db, query: q}
	result, err := pq.db.Exec(pq.query, post.Title, post.IsClosed, post.Detail,
		post.Title, post.IsClosed, post.Detail)
	if err != nil {
		return 0, errors.WithMessage(err, "mariadb: posts: Insert error")
	}
	return result.LastInsertId()
}

func (dc *DatabaseClient) UpdatePost(ctx context.Context, post *Post) error {
	q := `UPDATE posts SET title=?, is_closed=?, detail=?  WHERE id=?`
	uq := &PostQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query, post.Title, post.IsClosed, post.Detail, post.Id)
	return err
}

// DeletePost2 is true delete from database instead of DeletePost just update the row
func (dc *DatabaseClient) DeletePost(ctx context.Context, id int) error {
	q := `DELETE FROM posts WHERE id=?`
	pq := &PostQuery{db: dc.db, query: q}
	_, err := pq.db.Exec(pq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryPost() *PostQuery {
	return &PostQuery{db: dc.db, query: `SELECT * FROM posts`}
}

// All will display all rows
func (pq *PostQuery) All(ctx context.Context) (*Posts, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := pq.db.Query("SELECT * FROM posts WHERE title like ?", "%%test%%")
	rows, err := pq.db.Query(pq.query, pq.args...)
	if err != nil {
		return nil, err
	}
	return mkPost(rows)
}

func (pq *PostQuery) First(ctx context.Context) (*Post, error) {
	nodes, err := pq.Limit(1).All(ctx)
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
func (pq *PostQuery) Where(cs ...[4]string) *PostQuery {
	pq.clauses = append(pq.clauses, cs...)
	return pq
}

func (pq *PostQuery) Order(condition string) *PostQuery {
	pq.order = condition
	return pq
}

func (pq *PostQuery) Limit(limit int) *PostQuery {
	pq.limit = &limit
	return pq
}

func (pq *PostQuery) Offset(offset int) *PostQuery {
	pq.offset = &offset
	return pq
}

func (pq *PostQuery) prepareQuery(ctx context.Context) error {
	if pq.clauses != nil {
		pq.query += " WHERE "
		for i, c := range pq.clauses {
			// TODO: 2nd clause cannot be tied together automaticly
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(pq.clauses)-1 {
				pq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				pq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			pq.args = append(pq.args, c[2])
		}
	}
	if pq.order != "" {
		pq.query += " ORDER BY ?"
		pq.args = append(pq.args, pq.order)
	}
	if pq.limit != nil {
		pq.query += " LIMIT ?"
		a := strconv.Itoa(*pq.limit)
		pq.args = append(pq.args, a)
	}
	if pq.offset != nil {
		pq.query += ", ?"
		a := strconv.Itoa(*pq.offset)
		pq.args = append(pq.args, a)
	}
	return nil
}

func mkPost(rows *sql.Rows) (*Posts, error) {
	var title, detail sql.NullString
	var create_time, update_time sql.NullTime
	var id, is_closed int
	var posts = &Posts{}
	for rows.Next() {
		if err := rows.Scan(&id, &title, &is_closed, &detail, &create_time, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkPost rows.Scan error")
		}
		posts.Collection = append(posts.Collection, &Post{
			Id:         id,
			Title:      title.String,
			Detail:     detail.String,
			IsClosed:   is_closed,
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
