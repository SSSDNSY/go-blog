package util

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

const nucUrl = "https://jingyan.baidu.com/user/nuc"
const referer = "https://jingyan.baidu.com/"

/**
构造百度请求，爬取所需数据
*/
func GetPerson(baiduId string) (string, error) {
	//构造请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", nucUrl, nil)

	if err != nil {
		logs.Error("httpUtil err : ", err)
		return "", err
	}

	//设置请求头
	req.Header.Set("Referer", referer)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "jingyan.baidu.com")
	req.Host = req.Header.Get("Host")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Content-Type", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Set("Cookie", baiduId)
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("GetPerson() httpUtil err : ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("httpUtil err : ", err)
		return "", err
	}
	return string(body), err
}

func GetPostExp(baiduId string, pn string) (htmlStr string) {
	url := "https://jingyan.baidu.com/user/nucpage/content?tab=exp&expType=published&pn=" + pn
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		logs.Error("", err)
		return ""
	}
	req.Header.Add("Host", " jingyan.baidu.com")
	req.Header.Add("Connection", " keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", " 1")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Add("Sec-Fetch-User", " ?1")
	req.Header.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Sec-Fetch-Site", " same-origin")
	req.Header.Add("Sec-Fetch-Mode", " nested-navigate")
	req.Header.Add("Referer", " https://jingyan.baidu.com/user/nucpage/content?tab=exp&expType=published&pn=20")
	//req.Header.Add("Accept-Encoding", " gzip, deflate, br")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9")
	req.Header.Add("Cookie", baiduId)
	res, err := client.Do(req)
	if err != nil {
		logs.Error("GetPostExp() client.Do(req) err : ", err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error("ioutil.ReadAll err : ", err)
		return ""
	}
	//fmt.Println(string(body))
	return string(body)
}

func GetReward(pn string) (htmlStr string) {
	//url := "https://jingyan.baidu.com/patch?tab=highquality&pn=" + pn
	url := "https://jingyan.baidu.com/patch"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logs.Error("", err)
		return ""

	}

	req.PostForm.Add("tab", "highquality")
	req.PostForm.Add("pn", pn)
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Add("Host", "jingyan.baidu.com")
	req.Header.Add("Connection", " keep-alive")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9")
	req.Header.Add("Connection", " keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")

	res, err := client.Do(req)
	if err != nil {
		logs.Error("GetReward() client.Do(req) err : ", err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error("ioutil.ReadAll err : ", err)
		return ""
	}
	return string(body)

}
