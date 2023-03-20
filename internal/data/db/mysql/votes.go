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
	vq := &VoteQuery{db: dc.db, query: q}
	_, err := vq.db.Exec(
		vq.query, vote.Title, vote.IsRadio, vote.A, vote.B, vote.C,
		vote.D, vote.E, vote.F, vote.G, vote.H, vote.HasTxtField,
		vote.Title, vote.IsRadio, vote.A, vote.B, vote.C,
		vote.D, vote.E, vote.F, vote.G, vote.H, vote.HasTxtField)
	return errors.WithMessage(err, "mariadb: votes: Insert error")
}

func (dc *DatabaseClient) UpdateVote(ctx context.Context, vote *Vote) error {
	q := `UPDATE votes SET title=?, is_radio=?, a=?, b=?, c=?, d=?, e=?, f=?,
	g=?, h=?, has_txt_field=? WHERE id=?`
	vq := &VoteQuery{db: dc.db, query: q}
	_, err := vq.db.Exec(vq.query, vote.Title, vote.IsRadio, vote.A, vote.B,
		vote.C, vote.D, vote.E, vote.F, vote.G, vote.H, vote.HasTxtField,
		vote.Id)
	return err
}

// DeleteVote2 is true delete from database instead of DeleteVote just update the row
func (dc *DatabaseClient) DeleteVote(ctx context.Context, id int) error {
	q := `DELETE FROM votes WHERE id=?`
	vq := &VoteQuery{db: dc.db, query: q}
	_, err := vq.db.Exec(vq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryVote() *VoteQuery {
	return &VoteQuery{db: dc.db, query: `SELECT * FROM votes`}
}

// All will display all rows
func (vq *VoteQuery) All(ctx context.Context) (*Votes, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := vq.db.Query("SELECT * FROM votes WHERE title like ?", "%%test%%")
	rows, err := vq.db.Query(vq.query, vq.args...)
	if err != nil {
		return nil, err
	}
	return mkVote(rows)
}

func (vq *VoteQuery) First(ctx context.Context) (*Vote, error) {
	nodes, err := vq.Limit(1).All(ctx)
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
func (vq *VoteQuery) Where(cs ...[4]string) *VoteQuery {
	vq.clauses = append(vq.clauses, cs...)
	return vq
}

func (vq *VoteQuery) Order(condition string) *VoteQuery {
	vq.order = condition
	return vq
}

func (vq *VoteQuery) Limit(limit int) *VoteQuery {
	vq.limit = &limit
	return vq
}

func (vq *VoteQuery) Offset(offset int) *VoteQuery {
	vq.offset = &offset
	return vq
}

func (vq *VoteQuery) prepareQuery(ctx context.Context) error {
	if vq.clauses != nil {
		vq.query += " WHERE "
		for i, c := range vq.clauses {
			// TODO: 2nd clause cannot be tied together automaticly
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(vq.clauses)-1 {
				vq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				vq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			vq.args = append(vq.args, c[2])
		}
	}
	if vq.order != "" {
		vq.query += " ORDER BY ?"
		vq.args = append(vq.args, vq.order)
	}
	if vq.limit != nil {
		vq.query += " LIMIT ?"
		a := strconv.Itoa(*vq.limit)
		vq.args = append(vq.args, a)
	}
	if vq.offset != nil {
		vq.query += ", ?"
		a := strconv.Itoa(*vq.offset)
		vq.args = append(vq.args, a)
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
