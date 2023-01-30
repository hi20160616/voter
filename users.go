package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id, state, deleted                        int
	ip, realname, nickname, avatar_url, phone string
	create_time, update_time                  time.Time
)

func rebuild() error {
	qCreateUsers := `
	DROP TABLE users;
	CREATE TABLE IF NOT EXISTS users (
	  id int(10) NOT NULL AUTO_INCREMENT,
	  ip INT(4) UNSIGNED UNIQUE,
	  realname VARCHAR(255),
	  nickname VARCHAR(255),
	  avatar_url VARCHAR(255),
	  phone VARCHAR(11),
	  state TINYINT(1) DEFAULT 0 COMMENT 'user state: 0=normal, 1=disable',
	  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
	  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	  PRIMARY KEY (id)
	);`

	if _, err := DB.Exec(qCreateUsers); err != nil {
		return err
	}
	return nil
}

func insert(ip, realname, nickname, avatar_url, phone string) (int64, error) {
	result, err := DB.Exec(`
		INSERT INTO users (ip, realname, nickname, avatar_url, phone)
		VALUES (INET_ATON(?), ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		ip=INET_ATON(?), realname=?, nickname=?, avatar_url=?, phone=?`,
		ip, realname, nickname, avatar_url, phone)
	if err != nil {
		return 0, err
	}
	userID, err := result.LastInsertId()
	return userID, err
}

func queryUser(id int) error {
	query := `SELECT * FROM users WHERE id = ?`
	return DB.QueryRow(query, 1).Scan(&id, &ip, &realname, &nickname, &avatar_url,
		&phone, &state, &deleted, &create_time, &update_time)
}
