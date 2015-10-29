//
package main

import (
    "fmt"
    "github.com/hu17889/go_spider/core/common/page"
    "github.com/hu17889/go_spider/core/common/request"
    "github.com/hu17889/go_spider/core/spider"
    "./htmlparser"
    "log"
    "net/http"
    "runtime"
    "path/filepath"
    "io/ioutil"
)

type HousePageProcesser struct {
    parser map[string]htmlparser.HouseHTMLParser
}

func NewHousePageProcesser() *HousePageProcesser {
    m := &HousePageProcesser{}
    m.parser = make(map[string]htmlparser.HouseHTMLParser)
    m.parser["http://www.bjjs.gov.cn/tabid/2167/default.aspx"] = htmlparser.NewBeijingFangORGHTMLParser()
    return m
}

func (this *HousePageProcesser) Process(p *page.Page) {
    if !p.IsSucc() {
        println(p.Errormsg())
        return
    }

    //fmt.Printf("url=%v\n", p.GetRequest().GetUrl())
    parser, ok := this.parser[p.GetRequest().GetUrl()]
    if !ok {
        log.Printf("Cannot find HTML parser for url : %v", p.GetRequest().GetUrl())
        return
    }

    err := parser.Parse(p)
    if err != nil {
        log.Printf("parse url : %v failed", p.GetRequest().GetUrl())
        return
    }

    fmt.Printf("%v\n", parser.ToJSON())

//    jForwardDeliveryHhousing := simplejson.New() // 期房
//    jReadyHouse := simplejson.New() // 现房(新房)
//    jSecondHandHouse := simplejson.New()    // 存量房（二手房）
//    jday := simplejson.New()
//    jmonth := simplejson.New()
//
//    //TODO 检查 html parsing 的返回值是否正确，如果不正确报警
//    // 期房
//    t := query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_totalCount']").Text()
//    jForwardDeliveryHhousing.Set("当前可售期房总套数", t)
//    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount']").Text()
//    jForwardDeliveryHhousing.Set("当前可售住宅套数", t)
//
//    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_totalCount4']").Text()
//    jForwardDeliveryHhousing.Set("期房网上签约总套数", t)
//    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount4']").Text()
//    jForwardDeliveryHhousing.Set("期房网上签约住宅套数", t)
//
//    jday.Set("期房", jForwardDeliveryHhousing)
//
//    // 存量房
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount']").Text()
//    jSecondHandHouse.Set("当前可售存量房总套数", t)
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount']").Text()
//    jSecondHandHouse.Set("当前可售存量房住宅套数", t)
//
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount4']").Text()
//    jSecondHandHouse.Set("存量房网上签约总套数", t)
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount4']").Text()
//    jSecondHandHouse.Set("存量房网上签约住宅套数", t)
//
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount2']").Text()
//    jSecondHandHouse.Set("新发布房源总套数", t)
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount2']").Text()
//    jSecondHandHouse.Set("新发布房源住宅套数", t)
//
//    jday.Set("存量房", jSecondHandHouse)
//
//    // 现房
//    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount6']").Text()
//    jReadyHouse.Set("当前可售现房住宅套数", t)
//
//    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_totalCount8']").Text()
//    jReadyHouse.Set("现房网上签约总套数", t)
//    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount8']").Text()
//    jReadyHouse.Set("现房网上签约住宅套数", t)
//
//    jday.Set("现房", jReadyHouse)
//
//
//    // 月批准预售许可证
//    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount2']").Text()
//    jmonth.Set("月批准预售许可证的住宅套数", t)
//
//    // 月度存量房网上签约
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount3']").Text()
//    jmonth.Set("月度存量房网上签约总套数", t)
//    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount3']").Text()
//    jmonth.Set("月度存量房网上签约住宅套数", t)
//
//    jresult := simplejson.New()
//    jresult.Set(LastDay(), jday)
//    jresult.Set(LastMonth(), jmonth)
//
//    jbuf, _ := jresult.EncodePretty()
//    fmt.Println(string(jbuf))
}

//func main() {
//    sp := spider.NewSpider(NewMyPageProcesser(), "TaskName")
//
//    urls := []string{
//        "http://www.bjjs.gov.cn/tabid/2167/default.aspx",
//    }
//
//    var reqs []*request.Request
//    for _, url := range urls {
//        req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
//        req.GetHeader().Set("", "")
//        reqs = append(reqs, req)
//    }
//    sp.SetThreadnum(2).GetAllByRequest(reqs)
//}
//// LastDay 返回昨天的日期字符串，例如"2015-10-27"
//func LastDay() string {
//    t := time.Now().Add(-24*time.Hour)
//    s := fmt.Sprintf("%v-%v-%v", t.Year(), int(t.Month()), t.Day())
//    return s
//}
//
//// LastDay 返回上个月的日期字符串，例如"2015-10"
//func LastMonth() string {
//    t := time.Now()
//    y := t.Year()
//    m := int(t.Month()) - 1
//    if m == 0 {
//        m = 12
//        y = y - 1
//    }
//    s := fmt.Sprintf("%v-%02d", y, m)
//    return s
//}

func (this *HousePageProcesser) SaveData() {
    for _, p := range this.parser {
        path := filepath.Join(DataPath, p.Name(), htmlparser.LastDay() + ".json")
        err := ioutil.WriteFile(path, []byte(p.ToJSON()), 0755)
        if err = nil {
            log.Printf("WriteFile to <%v> OK", path)
        } else {
            log.Printf("writer JSON data to <%v> OK", path)
        }
    }
}

var DataPath = "/home/weizili/fang"

func init() {
    if runtime.GOOS == "windows" {
        DataPath = "e:/360yunpan/fangstat"
    }
}

func main() {
    hpp := NewHousePageProcesser()
    sp := spider.NewSpider(hpp, "TaskName")

    var reqs []*request.Request
    hh := make(http.Header)
    hh["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36"}
    for url, _ := range hpp.parser {
        req := request.NewRequest(url, "html", "", "GET", "", hh, nil, nil, nil)
        reqs = append(reqs, req)
    }
    sp.SetThreadnum(2).GetAllByRequest(reqs)
    hpp.SaveData()
}