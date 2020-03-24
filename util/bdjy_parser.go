package util

import (
	"github.com/PuerkitoBio/goquery"
	"go-blog/models"
	"log"
	"strconv"
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
func ParseExPublished(htmlTxt string, bdid string) map[string]string {
	totalView := 0
	totalVote := 0
	totalFavo := 0

	result := make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))
	if err != nil {
		log.Fatal("goquery.NewDocumentFromReader err :", err)
		return result
	}
	//totalView
	str := doc.Find(".padding8").Get(2).Attr[1].Val
	totalPn, _ := strconv.Atoi(strings.Split(str, "=")[3])

	for i := 0; i <= totalPn; i += 20 {
		htmstr := GetExpNum(bdid, strconv.Itoa(i))
		tmp := GetOnePageTotal(htmstr)
		v, _ := strconv.Atoi(tmp["v"])
		t, _ := strconv.Atoi(tmp["t"])
		f, _ := strconv.Atoi(tmp["f"])
		totalView += v
		totalVote += t
		totalFavo += f
	}

	//赋值
	result["totalPn"] = strings.Split(str, "=")[3]
	result["totalVote"] = strconv.Itoa(totalVote)
	result["totalFavo"] = strconv.Itoa(totalFavo)
	result["totalView"] = strconv.Itoa(totalView)
	return result
}

/**
* 得到一页的统计值
 */
func GetOnePageTotal(htmlTxt string) map[string]string {
	var view = 0
	var favo = 0
	var vote = 0

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))
	doc.Find("p.f-meta").Each(func(i int, s *goquery.Selection) {
		_v, _ := strconv.Atoi(s.Find("span.view-count").Text())
		_t, _ := strconv.Atoi(s.Find("span.vote-count").Text())
		_f, _ := strconv.Atoi(s.Find("span.favo-count").Text())
		view += _v
		vote += _t
		favo += _f
	})
	rtnMap := make(map[string]string)
	rtnMap["v"] = strconv.Itoa(view)
	rtnMap["t"] = strconv.Itoa(vote)
	rtnMap["f"] = strconv.Itoa(favo)

	return rtnMap
}
