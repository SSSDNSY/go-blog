package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go-blog/util"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
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

type ScRec struct {
	//'id', 'user_step', 'user_time', 'user_infos', 'add_time', 'update_time', 'ext_data'
	Id         int64     `orm:"auto"`
	UserStep   int32     `orm:"index"`
	UserTime   float64   `orm:"index"`
	UserInfo   string    `orm:"index"`
	AddTime    time.Time `orm:"index"`
	UpdateTime time.Time
	ExtData    string
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
	util.Init()
	orm.Debug = util.OrmDebug
	orm.RegisterModel(new(Category), new(Topic), new(Comment), new(ScRec))
	orm.RegisterDriver(util.DriveName, orm.DRMySQL)
	orm.RegisterDataBase("default", util.DriveName, util.DbConn, 10)
}

func AddScRecord(userInfo string, userStep int32, userTime float64) error {
	if len(userInfo) <= 0 {
		return errors.New("userInfo 不能为空")
	}
	if userTime <= 0 {
		return errors.New("userTime 必须大于0")
	}
	if userStep <= 0 {
		return errors.New("userStep 必须大于0")
	}
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	cate := &ScRec{
		UserStep:   userStep,
		UserTime:   userTime,
		UserInfo:   userInfo,
		AddTime:    time.Now().Add(8 * time.Hour),
		UpdateTime: time.Now().Add(8 * time.Hour),
		ExtData:    "",
	}
	_, err := o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
func QueryScRecord(limit int) ([]*ScRec, error) {
	if limit <= 0 || limit > 500 {
		return nil, errors.New("limit 入参非法！")
	}
	o := orm.NewOrm()
	sc := make([]*ScRec, 0)
	qs := o.QueryTable("sc_rec").
		Exclude("user_info", "AI player").
		OrderBy("user_step", "user_time", "add_time", "update_time").
		Limit(limit)
	_, err := qs.All(&sc)
	return sc, err
}

func AddCateGory(name string) error {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	cate := &Category{Title: name, Created: time.Now().Add(8 * time.Hour), TopicTime: time.Now().Add(8 * time.Hour), TopicLast: time.Now().Add(8 * time.Hour)}
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
		Created:        time.Now().Add(8 * time.Hour),
		Updated:        time.Now().Add(8 * time.Hour),
		Views:          0,
		Author:         "",
		ReplyTime:      time.Now().Add(8 * time.Hour),
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

func GetPageTopic(pageNum int64, pageSize int64) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	count, err := qs.Count()
	if count%pageSize < pageNum {
		return nil, err
	}
	_, err = qs.OrderBy("-updated").Limit(pageSize, (pageNum-1)*pageSize).All(&topics)
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
		topic.Created = time.Now().Add(8 * time.Hour)
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
		Created: time.Now().Add(8 * time.Hour),
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
		topic.ReplyTime = time.Now().Add(8 * time.Hour)
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

func GetBlogCount() (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Topic")
	count, err := qs.Count()
	if err != nil {
		return count, err
	}
	return count, nil
}
func GetCateCount() (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Category")
	count, err := qs.Count()
	if err != nil {
		return count, err
	}
	return count, nil
}

func transId(id string) {

}
