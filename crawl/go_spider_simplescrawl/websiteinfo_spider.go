//
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	"strings"
	"encoding/json"
	"github.com/bitly/go-simplejson"
)

var result = make(map[string]map[string]string)

type MyPageProcesser struct {  
}

func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()

	// 360.cn
	query.Find("head meta").Each(func(i int, s *goquery.Selection) {
		n, _ := s.Attr("name")
		if strings.ToLower(n) == "description" {
			if content, exist := s.Attr("content"); exist {
				p.AddField("description", content)
			}

		}
	})
	query.Find("head link").Each(func(i int, s *goquery.Selection) {
		n, _ := s.Attr("rel")
		if strings.ToLower(n) == "icon" {
			if href, exist := s.Attr("href"); exist {
				p.AddField("icon", href)
			}
		}
	})

	// 99bill.com
	if _, exist := p.GetPageItems().GetAll()["description"]; !exist {
		t := query.Find("head title").Text()
		if len(t) > 0 {
			p.AddField("description", t)
		}
	}
	query.Find("div[class='logo'] img").Each(func(i int, s *goquery.Selection) {
		n, _ := s.Attr("src")
		p.AddField("icon", n)
	})
	
	result[p.GetRequest().GetUrl()] = p.GetPageItems().GetAll()
}

func main() {
	// spider input:
	//  PageProcesser ;
	//  task name used in Pipeline for record;
	sp := spider.NewSpider(NewMyPageProcesser(), "TaskName")

	urls := []string{
		"http://360.cn",
		"http://99bill.com",
		"http://baidu.com",
		//		"http://bestpay.com.cn",
		//		"http://chinahr.com",
		//		"http://kugou.com",
		//		"http://lakala.com",
		//		"http://letao.com",
	}

	var reqs []*request.Request
	for _, url := range urls {
		req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
		reqs = append(reqs, req)
	}
	pageItemsArr := sp.SetThreadnum(2).GetAllByRequest(reqs)
	//pageItemsArr := sp.SetThreadnum(2).GetAll(urls, "html")
	for _, item := range pageItemsArr {
		url := item.GetRequest().GetUrl()
		
		for name, value := range item.GetAll() {
			fmt.Println(url + "\t\t" + name + "\t:\t" + value)
		}
	}
	
	println("--------------------------------------- Reuslt -------------------------------")
	if jbuf, err := json.Marshal(result); err == nil {
	 	json, err :=simplejson.NewJson(jbuf)
	 	fmt.Println(err)
	 	jbuf, err = json.EncodePretty()
	 	fmt.Println(err)
	 	fmt.Println(string(jbuf))
	} else {
		println("json.Marshal ERROR", err)
	}
}
