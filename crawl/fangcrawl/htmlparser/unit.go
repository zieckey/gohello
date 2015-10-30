package htmlparser

import (
    "github.com/hu17889/go_spider/core/common/page"
    "time"
    "fmt"
)

type HouseHTMLParser interface {
    Parse(p *page.Page) error
    ToJSON() string
    Name() string
}

type HouseORGParser struct {
    CityName string // 城市名
    URL string // 市房管局网页名称
}

// LastDay 返回昨天的日期字符串，例如"2015-10-27"
func LastDay() string {
    t := time.Now().Add(-24*time.Hour)
    s := fmt.Sprintf("%v-%v-%v", t.Year(), int(t.Month()), t.Day())
    return s
}

// LastDay 返回上个月的日期字符串，例如"2015-10"
func LastMonth() string {
    t := time.Now()
    y := t.Year()
    m := int(t.Month()) - 1
    if m == 0 {
        m = 12
        y = y - 1
    }
    s := fmt.Sprintf("%v-%02d", y, m)
    return s
}