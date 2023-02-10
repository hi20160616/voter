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

type User struct {
	Id, State, Deleted                                               int
	Username, Password, Realname, Nickname, AvatarUrl, Phone, UserIp string
	CreateTime, UpdateTime                                           time.Time
}

type Users struct {
	Collection []*User
}

type UserQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertUser(ctx context.Context, user *User) error {
	q := `INSERT INTO users(
		username, password, realname, nickname, avatar_url, phone, user_ip
		) VALUES (?, ?, ?, ?, ?, ?, INET_ATON(?))
		ON DUPLICATE KEY UPDATE
		username=?, password=?, realname=?, nickname=?,
		avatar_url=?, phone=?, user_ip=INET_ATON(?)`
	aq := &UserQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query,
		user.Username, user.Password, user.Realname, user.Nickname,
		user.AvatarUrl, user.Phone, user.UserIp,
		user.Username, user.Password, user.Realname, user.Nickname,
		user.AvatarUrl, user.Phone, user.UserIp,
	)
	return errors.WithMessage(err, "mysql: users: Insert error")
}

func (dc *DatabaseClient) UpdateUser(ctx context.Context, user *User) error {
	q := `UPDATE users SET
		password=?, realname=?, nickname=?,
		avatar_url=?, phone=?, user_ip=INET_ATON(?), state=?
		WHERE id=?`
	uq := &UserQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query,
		user.Password, user.Realname, user.Nickname,
		user.AvatarUrl, user.Phone, user.UserIp, user.State,
		user.Id)
	return err
}

// DeleteUser2 is true delete from database instead of DeleteUser just update the row
func (dc *DatabaseClient) DeleteUser2(ctx context.Context, id int) error {
	q := `DELETE FROM users WHERE id=?`
	aq := &UserQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser is soft delete, not delete from database,
// but update deleted field to 1
// DeleteUser is cooperate with All(ctx), that just return
// all rows except deleted is 1
func (dc *DatabaseClient) DeleteUser(ctx context.Context, id int) error {
	q := `UPDATE users SET deleted=? WHERE id=?`
	aq := &UserQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, 1, id)
	return err
}

func (dc *DatabaseClient) UndeleteUser(ctx context.Context, id int) error {
	q := `UPDATE users SET deleted=? WHERE id=?`
	aq := &UserQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, 0, id)
	return err
}

func (dc *DatabaseClient) QueryUser() *UserQuery {
	return &UserQuery{db: dc.db,
		query: `SELECT
		id, username, password, realname, nickname, avatar_url, phone,
		INET_NTOA(user_ip), state, deleted, create_time, update_time
		FROM users`}
}

// All2 will display all rows even if deleted field value is 1
func (aq *UserQuery) All2(ctx context.Context) (*Users, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := aq.db.Query("SELECT * FROM users WHERE title like ?", "%%test%%")
	rows, err := aq.db.Query(aq.query, aq.args...)
	if err != nil {
		return nil, err
	}
	return mkUser(rows)
}

// All will display all lines that the deleted field value is 0
func (aq *UserQuery) All(ctx context.Context) (*Users, error) {
	return aq.Where([4]string{"deleted", "=", "0"}).All2(ctx)
}

func (aq *UserQuery) First(ctx context.Context) (*User, error) {
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
func (aq *UserQuery) Where(cs ...[4]string) *UserQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *UserQuery) Order(condition string) *UserQuery {
	aq.order = condition
	return aq
}

func (aq *UserQuery) Limit(limit int) *UserQuery {
	aq.limit = &limit
	return aq
}

func (aq *UserQuery) Offset(offset int) *UserQuery {
	aq.offset = &offset
	return aq
}

func (aq *UserQuery) prepareQuery(ctx context.Context) error {
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

func mkUser(rows *sql.Rows) (*Users, error) {
	var username, password, realname, nickname, avatar_url,
		phone, user_ip sql.NullString
	var create_time, update_time sql.NullTime
	var id, state, deleted int
	var users = &Users{}
	for rows.Next() {
		if err := rows.Scan(&id, &username, &password, &realname, &nickname,
			&avatar_url, &phone, &user_ip, &state, &deleted,
			&create_time, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkUser rows.Scan error")
		}
		users.Collection = append(users.Collection, &User{
			Id:         id,
			Username:   username.String,
			Password:   password.String,
			Realname:   realname.String,
			Nickname:   nickname.String,
			AvatarUrl:  avatar_url.String,
			Phone:      phone.String,
			UserIp:     user_ip.String,
			State:      state,
			Deleted:    deleted,
			CreateTime: create_time.Time,
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkUser error")
	}
	return users, nil
}
