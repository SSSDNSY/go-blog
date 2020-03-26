package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	_DB_NAME        = "data/app1.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id         int64
	Title      string
	Created    time.Time `orm:"index"`
	Views      int64     `orm:"index"`
	TopicTime  time.Time
	TopicCount int64
	TopicLast  time.Time
}

type Topic struct {
	Id             int64
	Uid            int64
	Title          string
	Category       string
	Labels         string
	Content        string `orm:"size(5000)"`
	Attachment     string
	Created        time.Time `orm:"index"`
	Updated        time.Time `orm:"index"`
	Views          int64     `orm:"index"`
	Author         string
	ReplyTime      time.Time `orm:"index"`
	ReplyCount     int64
	ReplyLastUsrId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

type Files struct {
	Id   int
	Name string
	Size int64
	Time time.Time
}

type BDParam struct {
	Uname  string
	Level  string
	Intro  string
	Expnum string
	Fans   string

	Returns  string
	Quality  string
	Interact string
	Cash     string
	Wealth   string
	Active   string
	Origin   string

	Timing string

	Others map[string]string
}

type Nav struct {
	TotalNum   int //总条数
	CurrentNum int //当前页
	pageNum    int //第几页
	PageSize   int //每页数目
}

/***********************************************************************************************************
TODO===================================================================================================
***********************************************************************************************************/

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

//
func AddCateGory(name string) error {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	cate := &Category{Title: name, Created: time.Now().Local(), TopicTime: time.Now().Local(), TopicLast: time.Now().Local()}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
func GetAllCateGories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func AddTopic(title, content, category, label, attachmentFileName string) error {
	//处理标签 $beego#$bee#
	label = "$" + strings.Join(
		strings.Split(strings.Trim(label, " "), " "), "#$") + "#"
	o := orm.NewOrm()
	topic := &Topic{
		Title:          title,
		Category:       category,
		Labels:         label,
		Content:        content,
		Attachment:     attachmentFileName,
		Created:        time.Now().Local(),
		Updated:        time.Now().Local(),
		Views:          0,
		Author:         "",
		ReplyTime:      time.Now().Local(),
		ReplyCount:     0,
		ReplyLastUsrId: 0,
	}
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		//topics,err=GetAllTopic(cate,false)
		//cate.TopicCount = len(topics)
		_, err = o.Update(cate)
		if err != nil {
			return err
		}
	}
	return err
}

func GetAllTopic(cate string, label string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate) //根据文章分类过滤
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#") //根据文章标签过滤  contains
		}
		_, err = qs.OrderBy("-updated").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetTopic(tid string) (*Topic, error) {

	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if nil != err {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	//update之后 保存到数据库之后的修改为页面格式 空格分隔
	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "$", " ", -1), "#", "", -1)
	return topic, err
}
func EditTopic(tid, title, content, category, label, attachmentFileName string) error {
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{
		Id: tidNum,
	}
	var oldCate, oldAttachmentFileName string
	if o.Read(topic) == nil {

		oldCate = topic.Category
		oldAttachmentFileName = topic.Attachment

		topic.Title = title
		topic.Category = category
		topic.Attachment = attachmentFileName
		topic.Content = content
		topic.Labels = label
		topic.Created = time.Now().Local()
		o.Update(topic)
	}
	//更新分类的文章数
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if nil == err {
			cate.TopicCount--
			o.Update(cate)
		}
	}
	//删除旧的附件
	if len(attachmentFileName) > 0 {
		err = os.Remove(path.Join("attachment", oldAttachmentFileName))
		if err != nil {
			beego.Error(err)
		}
	}
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if nil == err {
		cate.TopicCount++
		o.Update(cate)
	}
	return nil
}

func DelTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if nil != err {
		return err
	}
	var oldCate string
	o := orm.NewOrm()
	topic := &Topic{
		Id: tidNum,
	}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if nil == err {
			cate.TopicCount--
			o.Update(cate)
		}
	}
	return err
}

func AddReply(tid, nikeName, content string) error {
	tidNu, err := strconv.ParseInt(tid, 10, 64)
	if nil != err {
		return err
	}
	reply := &Comment{
		Id:      0,
		Tid:     tidNu,
		Name:    nikeName,
		Content: content,
		Created: time.Now().Local(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}
	topic := &Topic{
		Id: tidNu,
	}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now().Local()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNu, err := strconv.ParseInt(tid, 10, 64)
	if nil != err {
		return nil, err
	}
	replies = make([]*Comment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNu).OrderBy("-created").All(&replies)
	return replies, err
}

func DelReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if nil != err {
		return err
	}

	var tidNum int64
	reply := &Comment{
		Id: ridNum,
	}
	o := orm.NewOrm()
	if o.Read(reply) == nil {
		tidNum = reply.Id
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}
	replies := make([]*Comment, 0)
	qs := o.QueryTable("Comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return nil
	}
	topic := &Topic{
		Id: tidNum,
	}
	if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].Created
		topic.ReplyCount = int64(len(replies))
		_, err = o.Update(topic)
	}

	return err
}

func transId(id string) {

}
