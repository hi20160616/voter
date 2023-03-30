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

type PostVote struct {
	Id, PostId, VoteId int
}

type PostVotes struct {
	Collection []*PostVote
}

type PostVoteQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertPostVote(ctx context.Context, postVote *PostVote) error {
	q := `INSERT INTO post_votes(post_id, vote_id) VALUES (?, ?)
		ON DUPLICATE KEY UPDATE post_id=?, vote_id=?`
	pvq := &PostVoteQuery{db: dc.db, query: q}
	_, err := pvq.db.Exec(pvq.query, postVote.PostId, postVote.VoteId,
		postVote.PostId, postVote.VoteId)
	return errors.WithMessage(err, "mysql: postVotes: Insert error")
}

func (dc *DatabaseClient) UpdatePostVote(ctx context.Context, postVote *PostVote) error {
	q := `UPDATE post_votes SET post_id=?, vote_id=? WHERE id=?`
	uq := &PostVoteQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query, postVote.PostId, postVote.VoteId, postVote.Id)
	return err
}

// DeletePostVote2 is true delete from database instead of DeletePostVote just update the row
func (dc *DatabaseClient) DeletePostVote(ctx context.Context, id int) error {
	q := `DELETE FROM post_votes WHERE id=?`
	pvq := &PostVoteQuery{db: dc.db, query: q}
	_, err := pvq.db.Exec(pvq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryPostVote() *PostVoteQuery {
	return &PostVoteQuery{db: dc.db, query: `SELECT * FROM post_votes`}
}

// All will display all rows
func (pvq *PostVoteQuery) All(ctx context.Context) (*PostVotes, error) {
	if err := pvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := pvq.db.Query("SELECT * FROM postVotes WHERE title like ?", "%%test%%")
	rows, err := pvq.db.Query(pvq.query, pvq.args...)
	if err != nil {
		return nil, err
	}
	return mkPostVote(rows)
}

func (pvq *PostVoteQuery) First(ctx context.Context) (*PostVote, error) {
	nodes, err := pvq.Limit(1).All(ctx)
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
func (pvq *PostVoteQuery) Where(cs ...[4]string) *PostVoteQuery {
	pvq.clauses = append(pvq.clauses, cs...)
	return pvq
}

func (pvq *PostVoteQuery) Order(condition string) *PostVoteQuery {
	pvq.order = condition
	return pvq
}

func (pvq *PostVoteQuery) Limit(limit int) *PostVoteQuery {
	pvq.limit = &limit
	return pvq
}

func (pvq *PostVoteQuery) Offset(offset int) *PostVoteQuery {
	pvq.offset = &offset
	return pvq
}

func (pvq *PostVoteQuery) prepareQuery(ctx context.Context) error {
	if pvq.clauses != nil {
		pvq.query += " WHERE "
		for i, c := range pvq.clauses {
			// TODO: 2nd clause cannot be tied together automaticly
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(pvq.clauses)-1 {
				pvq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				pvq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			pvq.args = append(pvq.args, c[2])
		}
	}
	if pvq.order != "" {
		pvq.query += " ORDER BY ?"
		pvq.args = append(pvq.args, pvq.order)
	}
	if pvq.limit != nil {
		pvq.query += " LIMIT ?"
		a := strconv.Itoa(*pvq.limit)
		pvq.args = append(pvq.args, a)
	}
	if pvq.offset != nil {
		pvq.query += ", ?"
		a := strconv.Itoa(*pvq.offset)
		pvq.args = append(pvq.args, a)
	}
	return nil
}

func mkPostVote(rows *sql.Rows) (*PostVotes, error) {
	var id, post_id, vote_id int
	var postVotes = &PostVotes{}
	for rows.Next() {
		if err := rows.Scan(&id, &post_id, &vote_id); err != nil {
			return nil, errors.WithMessage(err, "mkPostVote rows.Scan error")
		}
		postVotes.Collection = append(postVotes.Collection, &PostVote{
			Id: id, PostId: post_id, VoteId: vote_id})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkPostVote error")
	}
	return postVotes, nil
}
