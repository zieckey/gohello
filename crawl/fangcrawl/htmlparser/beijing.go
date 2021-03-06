package htmlparser

import (
    "github.com/hu17889/go_spider/core/common/page"
    "github.com/bitly/go-simplejson"
    "strconv"
    "errors"
    "log"
)


type HouseORGParser struct {
    CityName string // 城市名
    URL string // 市房管局网页名称
}

type BeijingFangORGHTMLParser struct {
    HouseORGParser
    HouseCount HouseCount
}

func NewBeijingFangORGHTMLParser() *BeijingFangORGHTMLParser {
    n := &BeijingFangORGHTMLParser{}
    n.CityName = "beijing"
    n.URL = "http://www.bjjs.gov.cn/tabid/2167/default.aspx"
    return n
}

func (this* BeijingFangORGHTMLParser) Name() string {
    return this.CityName
}

func (this* BeijingFangORGHTMLParser) Parse(p *page.Page) error {
    if !p.IsSucc() {
        println(p.Errormsg())
        return errors.New(p.Errormsg())
    }

    query := p.GetHtmlParser()

    var err error

    // 期房
    t := query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_totalCount']").Text()
    if this.HouseCount.ForwardDeliveryHouse.TotalSellingCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 当前可售期房总套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount']").Text()
    if this.HouseCount.ForwardDeliveryHouse.ResidenceSellingCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 当前可售住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_totalCount4']").Text()
    if this.HouseCount.ForwardDeliveryHouse.TotalSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 期房网上签约总套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount4']").Text()
    if this.HouseCount.ForwardDeliveryHouse.ResidenceSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 期房网上签约住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }


    // 存量房
    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount']").Text()
    if this.HouseCount.SecondHandHouse.TotalSellingCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 当前可售存量房总套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount']").Text()
    if this.HouseCount.SecondHandHouse.ResidenceSellingCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 当前可售存量房住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount4']").Text()
    if this.HouseCount.SecondHandHouse.TotalSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 存量房网上签约总套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount4']").Text()
    if this.HouseCount.SecondHandHouse.ResidenceSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 存量房网上签约住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount2']").Text()
    if this.HouseCount.SecondHandHouse.NewPublishedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 新发布房源总套数 错误：%v\n", this.CityName, err.Error())
        return err
    }
    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount2']").Text()
    if this.HouseCount.SecondHandHouse.NewPublishedResidenceCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 新发布房源住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    // 现房
    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount6']").Text()
    if this.HouseCount.ReadyHouse.ResidenceSellingCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 当前可售现房住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_totalCount8']").Text()
    if this.HouseCount.ReadyHouse.TotalSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 现房网上签约总套数 错误：%v\n", this.CityName, err.Error())
        return err
    }
    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount8']").Text()
    if this.HouseCount.ReadyHouse.ResidenceSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 现房网上签约住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }


    // 月批准预售许可证
    t = query.Find("div[class='cont_top_wrap'] td[id='ess_ctr5115_ContentPane'] span[id='ess_ctr5115_FDCJY_HouseTransactionStatist_residenceCount2']").Text()
    if this.HouseCount.MonthHouse.NewPublishedResidenceCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 月批准预售许可证的住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    // 月度存量房网上签约
    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_totalCount3']").Text()
    if this.HouseCount.MonthHouse.TotalSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 月度存量房网上签约总套数 错误：%v\n", this.CityName, err.Error())
        return err
    }
    t = query.Find("td[id='ess_ctr5112_ContentPane'] div[id='ess_ctr5112_ModuleContent'] span[id='ess_ctr5112_FDCJY_SignOnlineStatistics_residenceCount3']").Text()
    if this.HouseCount.MonthHouse.ResidenceSignedCount, err = strconv.Atoi(t); err != nil {
        log.Printf("%v 月度存量房网上签约住宅套数 错误：%v\n", this.CityName, err.Error())
        return err
    }

    return nil
}

func (this* BeijingFangORGHTMLParser) GetHouseCount() HouseCount {
    return this.HouseCount
}

func (this* BeijingFangORGHTMLParser) ToJSON(readable bool) string {

    jForwardDeliveryHhousing := simplejson.New() // 期房
    jReadyHouse := simplejson.New() // 现房(新房)
    jSecondHandHouse := simplejson.New()    // 存量房（二手房）
    jday := simplejson.New()
    jmonth := simplejson.New()

    // 期房
    jForwardDeliveryHhousing.Set("当前可售期房总套数", this.HouseCount.ForwardDeliveryHouse.TotalSellingCount)
    jForwardDeliveryHhousing.Set("当前可售住宅套数", this.HouseCount.ForwardDeliveryHouse.ResidenceSellingCount)
    jForwardDeliveryHhousing.Set("期房网上签约总套数", this.HouseCount.ForwardDeliveryHouse.TotalSignedCount)
    jForwardDeliveryHhousing.Set("期房网上签约住宅套数", this.HouseCount.ForwardDeliveryHouse.ResidenceSignedCount)
    jday.Set("期房", jForwardDeliveryHhousing)

    // 存量房
    jSecondHandHouse.Set("当前可售存量房总套数", this.HouseCount.SecondHandHouse.TotalSellingCount)
    jSecondHandHouse.Set("当前可售存量房住宅套数", this.HouseCount.SecondHandHouse.ResidenceSellingCount)
    jSecondHandHouse.Set("存量房网上签约总套数", this.HouseCount.SecondHandHouse.TotalSignedCount)
    jSecondHandHouse.Set("存量房网上签约住宅套数", this.HouseCount.SecondHandHouse.ResidenceSignedCount)
    jSecondHandHouse.Set("新发布房源总套数", this.HouseCount.SecondHandHouse.NewPublishedCount)
    jSecondHandHouse.Set("新发布房源住宅套数", this.HouseCount.SecondHandHouse.NewPublishedResidenceCount)
    jday.Set("存量房", jSecondHandHouse)

    // 现房
    jReadyHouse.Set("当前可售现房住宅套数", this.HouseCount.ReadyHouse.ResidenceSellingCount)
    jReadyHouse.Set("现房网上签约总套数", this.HouseCount.ReadyHouse.TotalSignedCount)
    jReadyHouse.Set("现房网上签约住宅套数", this.HouseCount.ReadyHouse.ResidenceSignedCount)
    jday.Set("现房", jReadyHouse)

    // 月批准预售许可证
    jmonth.Set("月批准预售许可证的住宅套数", this.HouseCount.MonthHouse.NewPublishedResidenceCount)
    jmonth.Set("月度存量房网上签约总套数", this.HouseCount.MonthHouse.TotalSignedCount)
    jmonth.Set("月度存量房网上签约住宅套数", this.HouseCount.MonthHouse.ResidenceSignedCount)

    jResult := simplejson.New()
    jResult.Set(LastDay(), jday)
    jResult.Set(LastMonth(), jmonth)

    if readable {
        buf, _ := jResult.EncodePretty()
        return  string(buf)
    } else {
        buf, _ := jResult.Encode()
        return  string(buf)
    }
}
