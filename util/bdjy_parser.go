package util

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"go-blog/models"
	"strconv"
	"strings"
)

//缓存
var CacheExp = GetIns()

/**
解析百度经验个人页面
*/
func ParsePerson(htmlTxt string) (*models.BDParam, error) {

	bd := &models.BDParam{}
	otherMap := make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))
	if err != nil {
		logs.Error("goquery.NewDocumentFromReader err :", err)
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
解析已发布的百度经验
*/
func ParseExPublished(bdid string) map[string]map[string]map[string]string {
	//totalView := 0
	//totalVote := 0
	//totalFavo := 0
	var totalPn int
	result := make(map[string]map[string]map[string]string, 1000)

	//第一次获得
	htmlTxt := GetPostExp(bdid, "0")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))
	if err != nil {
		logs.Error("goquery.NewDocumentFromReader err :", err)
		return result
	}

	//totalView
	str, ok := doc.Find("a.padding8").Last().Attr("href")
	if ok {
		totalPn, _ = strconv.Atoi(strings.Split(str, "=")[3])
	} else {
		return result
	}

	//爬取所有经验模型数据
	for i := 20; totalPn > 20 && i <= totalPn; i += 20 {
		htmstr := GetPostExp(bdid, strconv.Itoa(i))
		tmp := GetOnePageTotal(htmstr)
		result["exp"+strconv.Itoa(i)] = tmp
	}

	//赋值
	//result["totalPn"] = strings.Split(str, "=")[3]
	//result["totalVote"] = strconv.Itoa(totalVote)
	//result["totalFavo"] = strconv.Itoa(totalFavo)
	//result["totalView"] = strconv.Itoa(totalView)

	return result
}

func ParRewardExp() map[int]string {
	htmlTxt := GetReward("0")

	var totalPn int
	var index int
	index = 0
	maps := make(map[int]string)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))
	if err != nil {
		logs.Error(" ParRewardExp goquery .NewDocumentFromReader err :", err)
	}
	doc.Find("div.li-par").Each(func(i int, s *goquery.Selection) {
		maps[index] = s.Text()
		index++
	})

	str, ok := doc.Find("a.padding8").Last().Attr("href")
	if ok {
		totalPn, _ = strconv.Atoi(strings.Split(str, "=")[8])

	}
	for idx := 15; idx < totalPn; idx += 15 {
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(GetReward(strconv.Itoa(idx))))
		doc.Find("div.li-par").Each(func(i int, s *goquery.Selection) {
			maps[index] = s.Text()
			index++
		})
	}

	//fmt.Println(maps)
	return maps
}

/**
* 得到一页的统计值
 */
func GetOnePageTotal(htmlTxt string) map[string]map[string]string {

	rtnMap := make(map[string]map[string]string)
	var idx = 1
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(htmlTxt))

	//获取单篇title信息
	doc.Find("p.f-titl").Each(func(i int, s *goquery.Selection) {
		title := make(map[string]string, 20)

		idx++
		et := s.Find("a.f14").Text()
		title["et"] = et

		ehref, ok := s.Find("a.f14").Attr("href")
		if ok {
			title["ehref"] = ehref
		}

		eid, ok := s.Find("a.op").Attr("eid")
		if ok {
			title["eid"] = eid
		}

		edate := s.Find("span.f-date").Text()
		title["edate"] = edate

		ehq, ok := s.Find("span.icon-high-quality").Attr("title")
		if ok {
			title["ehq"] = ehq
		}
		rtnMap["tit"+strconv.Itoa(idx)] = title
	})
	idx = 1
	//获取单篇meta信息
	doc.Find("p.f-meta").Each(func(i int, s *goquery.Selection) {
		meta := make(map[string]string, 20)
		idx++
		eview := s.Find("span.view-count").Text()
		evote := s.Find("span.vote-count").Text()
		efavo := s.Find("span.favo-count").Text()
		efvote := s.Find("span.f-vote").Text()
		meta["eview"] = eview
		meta["evote"] = evote
		meta["vfavo"] = efavo
		meta["vfavo"] = efvote
		rtnMap["mata"+strconv.Itoa(idx)] = meta

	})

	return rtnMap
}
