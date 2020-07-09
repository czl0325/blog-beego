package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
)

const DB_NAME = "data/blog.db"
const SQLITE3_DRIVER = "sqlite3"

type Category struct {
	Id              int64
	Title           string    `orm:"unique;index"`
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	Views           int64
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime      time.Time `orm:"auto_now;type(datetime)"`
	Views           int64
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB()  {
	_, err := os.Stat(DB_NAME)
	if err != nil || os.IsNotExist(err) {
		os.MkdirAll(path.Dir(DB_NAME), os.ModePerm)
		os.Create(DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", SQLITE3_DRIVER, DB_NAME, 10)
}