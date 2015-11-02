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
    "path/filepath"
    "io/ioutil"
    "flag"
    "os"
    "runtime"
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


func (this *HousePageProcesser) Finish() {

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

    fmt.Printf("%v\n", parser.ToJSON(true))
}

func (this *HousePageProcesser) SaveData(outputDir string) {
    for _, p := range this.parser {
        outputDir = filepath.Join(outputDir, p.Name())
        err := os.MkdirAll(outputDir, 0755)
        if err == nil {
            log.Printf("os.MkdirAll <%v> OK", outputDir)
        } else {
            log.Printf("os.MkdirAll <%v> failed: %s", outputDir, err.Error())
        }
        path := filepath.Join(outputDir, htmlparser.LastDay() + ".json")
        err = ioutil.WriteFile(path, []byte(p.ToJSON(false)), 0755)
        if err == nil {
            log.Printf("WriteFile to <%v> OK", path)
        } else {
            log.Printf("writer JSON data to <%v> failed: %s", path, err.Error())
        }

        //TODO only insert to beijing table
        //dbutil.InsertToMySQL(p.ToJSON(false))
    }
}

var defaultOutput = "/var/mysql_backups"
var defaultLogFile = "/var/mysql_backups/crawling.log"
func init() {
    if runtime.GOOS == "windows" {
        defaultOutput = "e:/1"
        defaultLogFile = "e:/1/crawling.log"
    }
}

func main() {
    var output *string = flag.String("output", defaultOutput, "The dir where to store the crawled output data ")
    var logfile *string = flag.String("logfile", defaultLogFile, "The log file")
    flag.Parse()

    logf, err := os.OpenFile(*logfile, os.O_APPEND|os.O_CREATE, os.ModeAppend)
    if err != nil {
        log.Printf("os.OpenFile <%s> failed : %v", *logfile, err.Error())
    }
    log.SetOutput(logf)
    log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

    err = os.MkdirAll(*output, 0755)
    if err != nil {
        log.Printf("mkdir <%s> failed : %v", *output, err.Error())
        os.Exit(-1)
    }

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
    hpp.SaveData(*output)
}