package htmlparser

import (
    "github.com/hu17889/go_spider/core/common/page"
    "time"
    "fmt"
)

type HouseHTMLParser interface {
    Parse(p *page.Page) error
    ToJSON(readable bool) string
    Name() string
    GetHouseCount() HouseCount
}

type HouseDayCount struct {
    TotalSellingCount          int // 当前可售房屋套数
    ResidenceSellingCount      int // 当前可售住宅套数
    TotalSignedCount           int // 网上签约总套数
    ResidenceSignedCount       int // 网上签约住宅套数
    NewPublishedCount          int // 新发布房源总套数
    NewPublishedResidenceCount int // 新发布房源住宅套数
}

type HouseMonthCount struct {
    TotalSignedCount           int // 月度网上签约总套数
    ResidenceSignedCount       int // 月度网上签约住宅套数
    NewPublishedResidenceCount int // 月度批准预售许可证的住宅套数
}

type HouseCount struct {
    ForwardDeliveryHouse HouseDayCount   // 期房
    ReadyHouse           HouseDayCount   // 现房
    SecondHandHouse      HouseDayCount   // 存量房(二手房)
    MonthHouse           HouseMonthCount // 上一个月数据
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