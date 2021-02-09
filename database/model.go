package database

import "time"

// User 用户表
type User struct {
	ID       uint64     `xorm:"id int(64) pk autoincr"`
	Name     string     `xorm:"varchar(256) notnull"`
	Password string     `xorm:"varchar(64) notnull"`
	Phone    string     `xorm:"varchar(64) unique notnull"`
	Created  time.Time  `xorm:"created"`
	Updated  time.Time  `xorm:"updated"`
	Deleted  *time.Time `xorm:"deleted"`
}
