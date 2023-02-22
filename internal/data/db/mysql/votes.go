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

type Vote struct {
	Id, IsRadio, HasTxtField      int
	Title, A, B, C, D, E, F, G, H string
	CreateTime, UpdateTime        time.Time
}

type Votes struct {
	Collection []*Vote
}

type VoteQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertVote(ctx context.Context, vote *Vote) error {
	q := `
INSERT INTO votes(title, is_radio, a, b, c, d, e, f, g, h, has_txt_field)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE title=?, is_radio=?, a=?, b=?, c=?, d=?, e=?, f=?,
g=?, h=?, has_txt_field=?`
	aq := &VoteQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(
		aq.query, vote.Title, vote.IsRadio, vote.A, vote.B, vote.C,
		vote.D, vote.E, vote.F, vote.G, vote.H, vote.HasTxtField,
		vote.Title, vote.IsRadio, vote.A, vote.B, vote.C,
		vote.D, vote.E, vote.F, vote.G, vote.H, vote.HasTxtField)
	return errors.WithMessage(err, "mariadb: votes: Insert error")
}

func (dc *DatabaseClient) UpdateVote(ctx context.Context, vote *Vote) error {
	q := `UPDATE votes SET title=?, is_radio=?, a=?, b=?, c=?, d=?, e=?, f=?,
	g=?, h=?, has_txt_field=? WHERE id=?`
	uq := &VoteQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query, vote.Title, vote.IsRadio, vote.A, vote.B,
		vote.C, vote.D, vote.E, vote.F, vote.G, vote.H, vote.HasTxtField,
		vote.Id)
	return err
}

// DeleteVote2 is true delete from database instead of DeleteVote just update the row
func (dc *DatabaseClient) DeleteVote(ctx context.Context, id int) error {
	q := `DELETE FROM votes WHERE id=?`
	aq := &VoteQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryVote() *VoteQuery {
	return &VoteQuery{db: dc.db, query: `SELECT * FROM votes`}
}

// All will display all rows
func (aq *VoteQuery) All(ctx context.Context) (*Votes, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := aq.db.Query("SELECT * FROM votes WHERE title like ?", "%%test%%")
	rows, err := aq.db.Query(aq.query, aq.args...)
	if err != nil {
		return nil, err
	}
	return mkVote(rows)
}

func (aq *VoteQuery) First(ctx context.Context) (*Vote, error) {
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
func (aq *VoteQuery) Where(cs ...[4]string) *VoteQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *VoteQuery) Order(condition string) *VoteQuery {
	aq.order = condition
	return aq
}

func (aq *VoteQuery) Limit(limit int) *VoteQuery {
	aq.limit = &limit
	return aq
}

func (aq *VoteQuery) Offset(offset int) *VoteQuery {
	aq.offset = &offset
	return aq
}

func (aq *VoteQuery) prepareQuery(ctx context.Context) error {
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

func mkVote(rows *sql.Rows) (*Votes, error) {
	var title, a, b, c, d, e, f, g, h sql.NullString
	var create_time, update_time sql.NullTime
	var id, is_radio, has_txt_field int
	var votes = &Votes{}
	for rows.Next() {
		if err := rows.Scan(&id, &title, &is_radio, &a, &b, &c, &d, &e, &f,
			&g, &h, &has_txt_field, &create_time, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkVote rows.Scan error")
		}
		votes.Collection = append(votes.Collection, &Vote{
			Id:          id,
			Title:       title.String,
			IsRadio:     is_radio,
			A:           a.String,
			B:           b.String,
			C:           c.String,
			D:           d.String,
			E:           e.String,
			F:           f.String,
			G:           g.String,
			H:           h.String,
			HasTxtField: has_txt_field,
			CreateTime:  create_time.Time,
			UpdateTime:  update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkVote error")
	}
	return votes, nil
}
