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
	aq := &PostVoteQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, postVote.PostId, postVote.VoteId,
		postVote.PostId, postVote.VoteId)
	return errors.WithMessage(err, "mariadb: postVotes: Insert error")
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
	aq := &PostVoteQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryPostVote() *PostVoteQuery {
	return &PostVoteQuery{db: dc.db, query: `SELECT * FROM post_votes`}
}

// All will display all rows
func (aq *PostVoteQuery) All(ctx context.Context) (*PostVotes, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := aq.db.Query("SELECT * FROM postVotes WHERE title like ?", "%%test%%")
	rows, err := aq.db.Query(aq.query, aq.args...)
	if err != nil {
		return nil, err
	}
	return mkPostVote(rows)
}

func (aq *PostVoteQuery) First(ctx context.Context) (*PostVote, error) {
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
func (aq *PostVoteQuery) Where(cs ...[4]string) *PostVoteQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *PostVoteQuery) Order(condition string) *PostVoteQuery {
	aq.order = condition
	return aq
}

func (aq *PostVoteQuery) Limit(limit int) *PostVoteQuery {
	aq.limit = &limit
	return aq
}

func (aq *PostVoteQuery) Offset(offset int) *PostVoteQuery {
	aq.offset = &offset
	return aq
}

func (aq *PostVoteQuery) prepareQuery(ctx context.Context) error {
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
