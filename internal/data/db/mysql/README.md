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
| user_ip          | int unsigned | YES  | UNI | NULL              |                                               |
| avatar_url  | varchar(255) | YES  |     | NULL              |                                               |
| phone       | varchar(11)  | YES  |     | NULL              |                                               |
| state       | tinyint(1)   | YES  |     | NULL              |                                               |
| deleted     | tinyint(1)   | YES  |     | 0                 |                                               |
| create_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| update_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
12 rows in set (0.00 sec)
```

### Votes
```
mysql> CREATE TABLE IF NOT EXISTS votes (
  id int(10) NOT NULL AUTO_INCREMENT,
  title VARCHAR(255),
  is_radio TINYINT(1) COMMENT 'is radio: 0=no, 1=yes(default)' DEFAULT 1,
  detail varchar(255),
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
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
| Field       | Type         | Null | Key | Default           | Extra                                         |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
| id          | int          | NO   | PRI | NULL              | auto_increment                                |
| title       | varchar(255) | YES  |     | NULL              |                                               |
| is_radio    | tinyint(1)   | YES  |     | NULL              |                                               |
| detail      | mediumtext   | YES  |     | NULL              |                                               |
| create_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| update_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
6 rows in set (0.00 sec)
```

### Posts
```
mysql> CREATE TABLE IF NOT EXISTS posts (
  id int(10) NOT NULL AUTO_INCREMENT,
  title VARCHAR(255),
  is_open TINYINT(1) COMMENT 'is open: 0=no, 1=yes(default)' DEFAULT 1,
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
| is_open     | tinyint(1)   | YES  |     | 1                 |                                               |
| detail      | varchar(255) | YES  |     | NULL              |                                               |
| create_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| update_time | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+-------------+--------------+------+-----+-------------------+-----------------------------------------------+
6 rows in set (0.00 sec)
```
