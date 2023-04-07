# MySQL

> Reference: https://hub.docker.com/_/mysql

```
docker run -p 3306:3306 --name=voter -e MYSQL_ROOT_PASSWORD='rootpassword' -d mysql/mysql-server:latest
docker exec -it voter mysql -uroot -prootpassword
```

## SQL
### Database
```
mysql> drop database voter;
mysql> create database voter;
mysql> create user 'voter_db_user'@'%' identified by 'voter_db_pwd';
mysql> grant all privileges on voter.* to 'voter_db_user'@'%';
mysql> flush privileges;
mysql> show databases;
mysql> use voter;
```

### Users
```
mysql> CREATE TABLE IF NOT EXISTS users (
  id int(10) NOT NULL AUTO_INCREMENT,
  username VARCHAR(255),
  password VARCHAR(255),
  realname VARCHAR(255),
  nickname VARCHAR(255),
  user_ip INT(4) UNSIGNED UNIQUE,
  avatar_url VARCHAR(255),
  phone VARCHAR(11),
  state TINYINT(1) DEFAULT 0 COMMENT 'user state: 0=normal, 1=disable',
  deleted TINYINT(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

mysql> show tables;
+-----------------+
| Tables_in_voter |
+-----------------+
| users           |
+-----------------+

mysql> describe users;
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
| Field       | Type         | Null | Key | Default           | Extra                                         |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
| id          | int          | NO   | PRI | NULL              | auto_increment                                |
| username    | varchar(255) | YES  |     | NULL              |                                               |
| password    | varchar(255) | YES  |     | NULL              |                                               |
| realname    | varchar(255) | YES  |     | NULL              |                                               |
| nickname    | varchar(255) | YES  |     | NULL              |                                               |
| user_ip     | int unsigned | YES  | UNI | NULL              |                                               |
| avatar_url  | varchar(255) | YES  |     | NULL              |                                               |
| phone       | varchar(11)  | YES  |     | NULL              |                                               |
| state       | tinyint(1)   | YES  |     | 0                 |                                               |
| deleted     | tinyint(1)   | YES  |     | 0                 |                                               |
| create_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| update_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
12 rows in set (0.01 sec)
```

### Votes
```
mysql> CREATE TABLE IF NOT EXISTS votes (
  id int(10) NOT NULL AUTO_INCREMENT,
  title VARCHAR(255),
  is_radio TINYINT(1) DEFAULT 0 COMMENT 'is radio: 1=no, 0=yes(default)',
  a varchar(255),
  b varchar(255),
  c varchar(255),
  d varchar(255),
  e varchar(255),
  f varchar(255),
  g varchar(255),
  h varchar(255),
  has_txt_field TINYINT(1) DEFAULT 0 COMMENT 'has txt field: 1=yes, 0=no(default)',
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

mysql> show tables;
+-----------------+
| Tables_in_voter |
+-----------------+
| users           |
| votes           |
+-----------------+
2 rows in set (0.01 sec)

mysql> desc votes;
+---------------+--------------+------+-----+-------------------+-----------------------------------------------+
| Field         | Type         | Null | Key | Default           | Extra                                         |
+---------------+--------------+------+-----+-------------------+-----------------------------------------------+
| id            | int          | NO   | PRI | NULL              | auto_increment                                |
| title         | varchar(255) | YES  |     | NULL              |                                               |
| is_radio      | tinyint(1)   | YES  |     | 0                 |                                               |
| a             | varchar(255) | YES  |     | NULL              |                                               |
| b             | varchar(255) | YES  |     | NULL              |                                               |
| c             | varchar(255) | YES  |     | NULL              |                                               |
| d             | varchar(255) | YES  |     | NULL              |                                               |
| e             | varchar(255) | YES  |     | NULL              |                                               |
| f             | varchar(255) | YES  |     | NULL              |                                               |
| g             | varchar(255) | YES  |     | NULL              |                                               |
| h             | varchar(255) | YES  |     | NULL              |                                               |
| has_txt_field | tinyint(1)   | YES  |     | 0                 |                                               |
| create_time   | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| update_time   | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+---------------+--------------+------+-----+-------------------+-----------------------------------------------+
14 rows in set (0.01 sec)
```

### Posts
```
mysql> CREATE TABLE IF NOT EXISTS posts (
  id int(10) NOT NULL AUTO_INCREMENT,
  title VARCHAR(255),
  is_closed TINYINT(1) DEFAULT 0 COMMENT 'is closed: 0=no(default), 1=yes',
  detail VARCHAR(255),
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

mysql> show tables;
+-----------------+
| Tables_in_voter |
+-----------------+
| posts           |
| users           |
| votes           |
+-----------------+
3 rows in set (0.00 sec)

mysql> desc posts;
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
| Field       | Type         | Null | Key | Default           | Extra                                         |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
| id          | int          | NO   | PRI | NULL              | auto_increment                                |
| title       | varchar(255) | YES  |     | NULL              |                                               |
| is_closed   | tinyint(1)   | YES  |     | 0                 |                                               |
| detail      | varchar(255) | YES  |     | NULL              |                                               |
| create_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| update_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
6 rows in set (0.01 sec)
```

### PostVotes
```
mysql> CREATE TABLE IF NOT EXISTS post_votes (
  id int(10) NOT NULL AUTO_INCREMENT,
  post_id int(10),
  vote_id int(10),
  PRIMARY KEY (id)
);

mysql> desc post_votes;
+---------+------+------+-----+---------+----------------+
| Field   | Type | Null | Key | Default | Extra          |
+---------+------+------+-----+---------+----------------+
| id      | int  | NO   | PRI | NULL    | auto_increment |
| post_id | int  | YES  |     | NULL    |                |
| vote_id | int  | YES  |     | NULL    |                |
+---------+------+------+-----+---------+----------------+
3 rows in set (0.00 sec)
```

### IpVotes
```
mysql> CREATE TABLE IF NOT EXISTS ip_votes (
  id int(10) NOT NULL AUTO_INCREMENT,
  ip int(4) UNSIGNED,
  vote_id int(10),
  opts VARCHAR(8),
  txt_field VARCHAR(255),
  PRIMARY KEY (id)
);

mysql> desc ip_votes;
+-----------+--------------+------+-----+---------+----------------+
| Field     | Type         | Null | Key | Default | Extra          |
+-----------+--------------+------+-----+---------+----------------+
| id        | int          | NO   | PRI | NULL    | auto_increment |
| ip        | int          | YES  |     | NULL    |                |
| vote_id   | int          | YES  |     | NULL    |                |
| opts      | varchar(8)   | YES  |     | NULL    |                |
| txt_field | varchar(255) | YES  |     | NULL    |                |
+-----------+--------------+------+-----+---------+----------------+
5 rows in set (0.00 sec)
```

### IpPosts
```
mysql> CREATE TABLE IF NOT EXISTS ip_posts (
  id int(10) NOT NULL AUTO_INCREMENT,
  ip int(4) UNSIGNED,
  post_id int(10),
  PRIMARY KEY (id)
);
```
