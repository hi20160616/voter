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

type IpVote struct {
	Id, VoteId         int
	Ip, Opts, TxtField string
}

type IpVotes struct {
	Collection []*IpVote
}

type IpVoteQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertIpVote(ctx context.Context, ipVote *IpVote) (int64, error) {
	q := `INSERT INTO ip_votes(ip, vote_id, opts, txt_field) VALUES (INET_ATON(?), ?, ?, ?)
		ON DUPLICATE KEY UPDATE ip=INET_ATON(?), vote_id=?, opts=?, txt_field=?`
	ivq := &IpVoteQuery{db: dc.db, query: q}
	x, err := ivq.db.Exec(
		ivq.query, ipVote.Ip, ipVote.VoteId, ipVote.Opts, ipVote.TxtField,
		ipVote.Ip, ipVote.VoteId, ipVote.Opts, ipVote.TxtField)
	if err != nil {
		return 0, errors.WithMessage(err, "mysql: ipVotes: Insert error")
	}
	return x.LastInsertId()
}

func (dc *DatabaseClient) UpdateIpVote(ctx context.Context, ipVote *IpVote) error {
	q := `UPDATE ip_votes SET ip=INET_ATON(?), vote_id=?, opts=?, txt_field=? WHERE id=?`
	uq := &IpVoteQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query, ipVote.Ip, ipVote.VoteId, ipVote.Opts,
		ipVote.TxtField, ipVote.Id)
	return err
}

// DeleteIpVote2 is true delete from database instead of DeleteIpVote just update the row
func (dc *DatabaseClient) DeleteIpVote(ctx context.Context, id int) error {
	q := `DELETE FROM ip_votes WHERE id=?`
	ivq := &IpVoteQuery{db: dc.db, query: q}
	_, err := ivq.db.Exec(ivq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryIpVote() *IpVoteQuery {
	return &IpVoteQuery{db: dc.db, query: `SELECT * FROM ip_votes`}
}

// All will display all rows
func (ivq *IpVoteQuery) All(ctx context.Context) (*IpVotes, error) {
	if err := ivq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := ivq.db.Query("SELECT * FROM ipVotes WHERE title like ?", "%%test%%")
	rows, err := ivq.db.Query(ivq.query, ivq.args...)
	if err != nil {
		return nil, err
	}
	return mkIpVote(rows)
}

func (ivq *IpVoteQuery) First(ctx context.Context) (*IpVote, error) {
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
func (ivq *IpVoteQuery) Where(cs ...[4]string) *IpVoteQuery {
	ivq.clauses = append(ivq.clauses, cs...)
	return ivq
}

func (ivq *IpVoteQuery) Order(condition string) *IpVoteQuery {
	ivq.order = condition
	return ivq
}

func (ivq *IpVoteQuery) Limit(limit int) *IpVoteQuery {
	ivq.limit = &limit
	return ivq
}

func (ivq *IpVoteQuery) Offset(offset int) *IpVoteQuery {
	ivq.offset = &offset
	return ivq
}

func (ivq *IpVoteQuery) prepareQuery(ctx context.Context) error {
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

func mkIpVote(rows *sql.Rows) (*IpVotes, error) {
	var id, vote_id int
	var ip, opts, txtField sql.NullString
	var ipVotes = &IpVotes{}
	for rows.Next() {
		if err := rows.Scan(&id, &ip, &vote_id, &opts, &txtField); err != nil {
			return nil, errors.WithMessage(err, "mkIpVote rows.Scan error")
		}
		ipVotes.Collection = append(ipVotes.Collection, &IpVote{
			Id: id, Ip: ip.String, VoteId: vote_id, Opts: opts.String,
			TxtField: txtField.String})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkIpVote error")
	}
	return ipVotes, nil
}
