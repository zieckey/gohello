package dbutil


import (
    "github.com/bmizerany/assert"
    "testing"
    "log"
)

func TestGetSecondHandHouseDayData(t *testing.T) {
    log.SetFlags(log.Ldate|log.Lshortfile|log.Ltime)
    json := `
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
    `
    year, mon, day, num, err := getSecondHandHouseDayData(json)
    assert.Equal(t, year, 2015)
    assert.Equal(t, mon, 10)
    assert.Equal(t, day, 27)
    assert.Equal(t, num, 872)
    assert.Equal(t, err, nil)

    year, mon, num, err = getSecondHandHouseMonthData(json)
    assert.Equal(t, year, 2015)
    assert.Equal(t, mon, 9)
    assert.Equal(t, num, 17888)
    assert.Equal(t, err, nil)
}