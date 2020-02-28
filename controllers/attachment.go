package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
	"io"
	"io/ioutil"
	"net/url"
	"os"
)

type AttachmentController struct {
	beego.Controller
}

const ATTACHMENT_PATH = "attachment"

func (this *AttachmentController) Get() {
	this.Data["IsFile"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	//文件下载操作
	if len(filePath) > 1 && filePath != "file" {
		if nil != err {
			beego.Error(err)
			return
		}
		f, err := os.Open(filePath)
		if nil != err {
			beego.Error(err)
			return
		}
		defer f.Close()
		_, err = io.Copy(this.Ctx.ResponseWriter, f)
		if nil != err {
			beego.Error(err)
			return
		}
		//文件展示操作
	} else {
		this.TplName = "file_list.html"

		//方法一
		//filepath.Walk(ATTACHMENT_PATH, func(path string, info os.FileInfo, err error) error {
		//	fmt.Println(path, info, "\n")
		//	return nil
		//})
		//方法二
		//for i, f := range rangeFiles(ATTACHMENT_PATH) {
		//	fmt.Printf("file id %d , %#v\n", i, f)
		//}
		this.Data["Files"] = rangeFiles(ATTACHMENT_PATH)
	}
}

func rangeFiles(dir string) []models.Files {
	fs, err := ioutil.ReadDir(dir)
	if err != nil || fs == nil {
		beego.Error(err)
		return nil
	}
	s1 := make([]models.Files, len(fs))
	for i, f := range fs {
		if f.IsDir() {
			s1[i].Name = f.Name() + "/"
			s1[i].Id = i
			s1[i].Time = f.ModTime()
			s1[i].Size = f.Size() / 1024
		} else {
			s1[i].Name = f.Name()
			s1[i].Id = i
			s1[i].Time = f.ModTime()
			s1[i].Size = f.Size() / 1024
		}
	}
	return s1
}
