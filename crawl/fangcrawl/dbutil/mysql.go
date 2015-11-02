package dbutil

import (

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/bitly/go-simplejson"
    "fmt"
    "time"
    "log"
    "runtime"
)



func InsertToMySQL(jd string) error {
    dburl := "stat:123456@tcp(127.0.0.1:3306)/statistic"
    if runtime.GOOS == "windows" {
        dburl = "root:@tcp(127.0.0.1:3306)/statistic"
    }
    db, err := sql.Open("mysql", dburl)
    if err != nil {
        panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

    {
        stmt, err := db.Prepare("insert into stat_esf_bj_netregist_by_day(year, mon, day, num) values(?,?,?,?)");
        if err != nil {
            log.Println(err.Error());
            return err;
        }
        defer stmt.Close();
        year, mon, day, num, err := getSecondHandHouseDayData(jd)
        if err != nil {
            return err
        }
        if result, err := stmt.Exec(year, mon, day, num); err == nil {
            if _, err := result.LastInsertId(); err == nil {
                log.Printf("insert year=%d mon=%d day=%d num=%d OK", year, mon, day, num)
            }
        }

        // test
        getSecondHandHouseMonthData(jd)
    }

    // 每个月第一天执行，将上一个月的统计数据计算出来
    if time.Now().Day() == 1 {
        stmt, err := db.Prepare("insert into stat_esf_bj_netregist_by_mon(year, mon, num) values(?,?,?)");
        if err != nil {
            log.Println(err.Error());
            return err;
        }
        defer stmt.Close();
        year, mon, num, err := getSecondHandHouseMonthData(jd)
        if err != nil {
            return err
        }
        if result, err := stmt.Exec(year, mon, num); err == nil {
            if _, err := result.LastInsertId(); err == nil {
                log.Printf("insert year=%d mon=%d num=%d OK", year, mon, num)
            }
        }
    }

    return nil
}

/*
{
  "2015-09": {
    "月度存量房网上签约住宅套数": 17888,
    "月度存量房网上签约总套数": 20041,
    "月批准预售许可证的住宅套数": 8196
  },
  "2015-10-27": {
    "存量房": {
      "存量房网上签约住宅套数": 872,
      "存量房网上签约总套数": 976,
      "当前可售存量房住宅套数": 34229,
      "当前可售存量房总套数": 36325,
      "新发布房源住宅套数": 770,
      "新发布房源总套数": 834
    },
    "期房": {
      "当前可售住宅套数": 43289,
      "当前可售期房总套数": 108015,
      "期房网上签约住宅套数": 209,
      "期房网上签约总套数": 290
    },
    "现房": {
      "当前可售现房住宅套数": 756164,
      "现房网上签约住宅套数": 48,
      "现房网上签约总套数": 229
    }
  }
}
*/
func getSecondHandHouseDayData(jd string) (year, mon, day, num int, err error) {
    json, err := simplejson.NewJson([]byte(jd))
    if err != nil {
        return 0,0,0,0, err
    }
    m, err := json.Map()
    if err != nil {
        return 0,0,0,0, err
    }
    log.Printf("len(map)=%d", len(m))
    var ymd string
    for k, _ := range m {
        log.Printf("key=%v", k)
        if len(k) == len("2015-10-27") {
            ymd = k
            break
        }
    }
    fmt.Sscanf(ymd, "%d-%d-%d", &year, &mon, &day)
    num = json.GetPath(ymd, "存量房", "存量房网上签约住宅套数").MustInt(-1)
    log.Printf("year=%d mon=%d day=%d num=%d", year, mon, day, num)
    return year, mon, day, num, nil
}

func getSecondHandHouseMonthData(jd string) (year, mon, num int, err error) {
    json, err := simplejson.NewJson([]byte(jd))
    if err != nil {
        return 0,0,0, err
    }
    m, err := json.Map()
    if err != nil {
        return 0,0,0, err
    }
    var ymd string
    for k, _ := range m {
        if len(k) == len("2015-10") {
            ymd = k
            break
        }
    }
    fmt.Sscanf(ymd, "%d-%d", &year, &mon)
    num = json.GetPath(ymd, "月度存量房网上签约住宅套数").MustInt(-1)
    log.Printf("year=%d mon=%d num=%d", year, mon, num)
    return year, mon, num, nil
}