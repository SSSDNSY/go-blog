#coding=utf-8
import re
import urllib

def getHtml(url):
    page = urllib.urlopen(url)
    html = page.read()
    return html

def getPage(html):
    reg = r'随笔-([0-9]+)'
    pageCount = re.findall(reg,html)
    return pageCount[0]

def getArticleUrl(html):
    reg = r'(http://www.cnblogs.com/humi/p/[0-9]+.html)'
    articleUrl = re.findall(reg,html)
    return articleUrl

def downloadPage(urlList):
    x = 0
    for article in urlList:
        urllib.urlretrieve(article,'%s.html' % x)
        x+=1

article = []
htmlStr = getHtml("http://www.cnblogs.com/humi/default.html")
pageCount = getPage(htmlStr)
page = int(pageCount)/40+1
for i in range(1,page+1):
    html = getHtml("http://www.cnblogs.com/humi/default.html?page="+str(i))
    articleUrl = getArticleUrl(html)
    article = article.__add__(articleUrl)

article = list(set(article))
downloadPage(article)