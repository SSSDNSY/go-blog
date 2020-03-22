package util

import (
	"github.com/PuerkitoBio/goquery"
	"go-blog/models"
	"log"
	"strings"
)

/**
解析百度经验个人页面
*/
func ParseBDParam(htmlTxt string) (*models.BDParam, error) {

	bd := &models.BDParam{}
	otherMap := make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))
	if err != nil {
		log.Fatal("goquery.NewDocumentFromReader err :", err)
		return bd, err
	}
	doc.Find("a.uname").Each(func(i int, s *goquery.Selection) {
		bd.Uname = s.Text()
	})
	doc.Find("span.active-value").Each(func(i int, s *goquery.Selection) {
		bd.Active = s.Text()
	})
	doc.Find("span.cash-value").Each(func(i int, s *goquery.Selection) {
		bd.Cash = s.Text()
	})
	doc.Find("span.interact-value").Each(func(i int, s *goquery.Selection) {
		bd.Interact = s.Text()
	})
	doc.Find("span.exp-num").Each(func(i int, s *goquery.Selection) {
		bd.Expnum = s.Text()
	})
	doc.Find("span.fans-num").Each(func(i int, s *goquery.Selection) {
		bd.Fans = s.Text()
	})
	doc.Find("span.quality-value").Each(func(i int, s *goquery.Selection) {
		bd.Intro = s.Text()
	})
	doc.Find("div.tooltip-header").Each(func(i int, s *goquery.Selection) {
		bd.Level = s.Text()
	})
	doc.Find("span.origin-value").Each(func(i int, s *goquery.Selection) {
		bd.Origin = s.Text()
	})
	doc.Find("span.huixiang-value").Each(func(i int, s *goquery.Selection) {
		bd.Returns = s.Text()
	})
	doc.Find("span.quality-value").Each(func(i int, s *goquery.Selection) {
		bd.Quality = s.Text()
	})
	doc.Find("span.wealth-value").Each(func(i int, s *goquery.Selection) {
		bd.Wealth = s.Text()
	})
	doc.Find("img[alt=" + bd.Uname + "]").Each(func(i int, s *goquery.Selection) {
		otherMap["avatorUrl"], _ = s.Attr("src")
	})
	bd.Others = otherMap
	return bd, err
}

/**
解析百度经验提交页面
*/
func ParseExPublished(htmlTxt string) map[string]string {
	//totalView:=0
	//totalVote
	//totalPn=0

	result := make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))
	if err != nil {
		log.Fatal("goquery.NewDocumentFromReader err :", err)
		return result
	}
	//totalView
	str := doc.Find(".padding8").Get(2).Attr[1].Val

	//赋值
	result["totalView"] = strings.Split(str, "=")[3]
	return result
}
