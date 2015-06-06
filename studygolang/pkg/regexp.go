package pkg

import (
"regexp"
	"fmt"
)

func TestRegexp() {
	// 比较两个字符串相等
	p1 := "abc"
	s := "abc"
	
	regx, err := regexp.Compile(p1)
	b := regx.MatchString(s)
	fmt.Printf("pattern=[%s] s=[%s] regx=%v err=%v\n", p1, s, b, err)
}