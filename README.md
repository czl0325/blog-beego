# blog-beego
beego开发个人博客项目，视频位于https://pan.baidu.com/s/1gultXms_P6urYaZputz1Ig，密码: p715，视频略旧，2015年的，我使用最新的beego，对工程做了改进，加入自己的理解。


#### 知识点

* 博客分类增删改查，文章增删改查，评论增删改查，文章与评论的一对多关系

#### 表结构设计

```go
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
	Labels          string
	Attachment      string
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime      time.Time `orm:"auto_now;type(datetime)"`
	Views           int64     `orm:"default(0)"`
	Author          string
	ReplyTime       time.Time  `orm:"null"`
	ReplyCount      int64      `orm:"default(0)"`
	ReplyLastUserId int64      `orm:"null"`
	Comments        []*Comment `orm:"reverse(many)"`
}

type Comment struct {
	Id         int64
	Tid        int64
	Topic      *Topic `orm:"rel(fk)"`
	Content    string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}
```