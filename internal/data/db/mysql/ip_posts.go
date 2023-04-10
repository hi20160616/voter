package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type IpPost struct {
	Id, PostId int
	Ip         string
}

type IpPosts struct {
	Collection []*IpPost
}

type IpPostQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertIpPost(ctx context.Context, ipPost *IpPost) (int64, error) {
	q := `INSERT INTO ip_posts(ip, post_id) VALUES (INET_ATON(?), ?)
		ON DUPLICATE KEY UPDATE ip=INET_ATON(?), post_id=?`
	ivq := &IpPostQuery{db: dc.db, query: q}
	x, err := ivq.db.Exec(
		ivq.query, ipPost.Ip, ipPost.PostId,
		ipPost.Ip, ipPost.PostId)
	if err != nil {
		return 0, errors.WithMessage(err, "mysql: ipPosts: Insert error")
	}
	return x.LastInsertId()
}

func (dc *DatabaseClient) UpdateIpPost(ctx context.Context, ipPost *IpPost) error {
	q := `UPDATE ip_posts SET ip=INET_ATON(?), post_id=? WHERE id=?`
	uq := &IpPostQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query, ipPost.Ip, ipPost.PostId, ipPost.Id)
	return err
}

func (dc *DatabaseClient) DeleteIpPost(ctx context.Context, id int) error {
	q := `DELETE FROM ip_posts WHERE id=?`
	ivq := &IpPostQuery{db: dc.db, query: q}
	_, err := ivq.db.Exec(ivq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryIpPost() *IpPostQuery {
	return &IpPostQuery{db: dc.db, query: `SELECT * FROM ip_posts`}
}

// All will display all rows
func (ivq *IpPostQuery) All(ctx context.Context) (*IpPosts, error) {
	if err := ivq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := ivq.db.Query("SELECT * FROM ipPosts WHERE title like ?", "%%test%%")
	rows, err := ivq.db.Query(ivq.query, ivq.args...)
	if err != nil {
		return nil, err
	}
	return mkIpPost(rows)
}

func (ivq *IpPostQuery) First(ctx context.Context) (*IpPost, error) {
	nodes, err := ivq.Limit(1).All(ctx)
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
func (ivq *IpPostQuery) Where(cs ...[4]string) *IpPostQuery {
	ivq.clauses = append(ivq.clauses, cs...)
	return ivq
}

func (ivq *IpPostQuery) Order(condition string) *IpPostQuery {
	ivq.order = condition
	return ivq
}

func (ivq *IpPostQuery) Limit(limit int) *IpPostQuery {
	ivq.limit = &limit
	return ivq
}

func (ivq *IpPostQuery) Offset(offset int) *IpPostQuery {
	ivq.offset = &offset
	return ivq
}

func (ivq *IpPostQuery) prepareQuery(ctx context.Context) error {
	if ivq.clauses != nil {
		ivq.query += " WHERE "
		for i, c := range ivq.clauses {
			// TODO: 2nd clause cannot be tied together automaticly
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(ivq.clauses)-1 {
				ivq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				ivq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			ivq.args = append(ivq.args, c[2])
		}
	}
	if ivq.order != "" {
		ivq.query += " ORDER BY ?"
		ivq.args = append(ivq.args, ivq.order)
	}
	if ivq.limit != nil {
		ivq.query += " LIMIT ?"
		a := strconv.Itoa(*ivq.limit)
		ivq.args = append(ivq.args, a)
	}
	if ivq.offset != nil {
		ivq.query += ", ?"
		a := strconv.Itoa(*ivq.offset)
		ivq.args = append(ivq.args, a)
	}
	return nil
}

func mkIpPost(rows *sql.Rows) (*IpPosts, error) {
	var id, post_id int
	var ip sql.NullString
	var ipPosts = &IpPosts{}
	for rows.Next() {
		if err := rows.Scan(&id, &ip, &post_id); err != nil {
			return nil, errors.WithMessage(err, "mkIpPost rows.Scan error")
		}
		ipPosts.Collection = append(ipPosts.Collection, &IpPost{
			Id: id, Ip: ip.String, PostId: post_id})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkIpPost error")
	}
	return ipPosts, nil
}
