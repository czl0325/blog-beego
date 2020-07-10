package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const DB_NAME = "data/blog.db"
const SQLITE3_DRIVER = "sqlite3"

type Category struct {
	Id              int64
	Title           string    `orm:"unique;index"`
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	Views           int64     `orm:"default(0)"`
	TopicTime       time.Time `orm:"null"`
	TopicCount      int64     `orm:"null"`
	TopicLastUserId int64     `orm:"null"`
}

type Topic struct {
	Id              int64
	Uid             int64
	Cid             int64
	Cname           string
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime      time.Time `orm:"auto_now;type(datetime)"`
	Views           int64     `orm:"default(0)"`
	Author          string
	ReplyTime       time.Time `orm:"null"`
	ReplyCount      int64     `orm:"default(0)"`
	ReplyLastUserId int64     `orm:"null"`
}

func RegisterDB() {
	_, err := os.Stat(DB_NAME)
	if err != nil || os.IsNotExist(err) {
		os.MkdirAll(path.Dir(DB_NAME), os.ModePerm)
		os.Create(DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", SQLITE3_DRIVER, DB_NAME, 10)
}

func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	categories := make([]*Category, 0)
	_, err := qs.All(&categories)
	return categories, err
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title: name,
	}
	_, err := o.Insert(cate)
	return err
}

func DeleteCategory(id int64) error {
	o := orm.NewOrm()
	cate := &Category{
		Id: id,
	}
	_, err := o.Delete(cate)
	return err
}

func AddTopic(id, title, cid, content string) error {
	o := orm.NewOrm()
	var err error

	_cid, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	category := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", _cid).One(&category)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("查无此分类")
	}

	topic := &Topic{
		Uid:        0,
		Cid:        category.Id,
		Cname:      category.Title,
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if len(id) != 0 {
		topic.Id, _ = strconv.ParseInt(id, 10, 64)
		_, err = o.Update(topic)
	} else {
		_, err = o.Insert(topic)
	}
	return err
}

func GetAllTopic(desc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if desc {
		_, err = qs.OrderBy("-CreateTime").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetTopicById(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func DeleteTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{
		Id: tid,
	}
	_, err = o.Delete(topic)
	return err
}
