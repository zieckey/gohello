package main

import (
	"fmt"
	"github.com/opesun/goquery"
	"strings"
)

var example = `<!DOCTYPE html>
<html>
	<head>
		<title>
		the title of the page
		</title>
	</head>
<body>
	<div class=hey custom_attr="wow"><h2>Title here</h2></div>
	<span><h2>Yoyoyo</h2></span>
	<div id="x">
		<span>
			content<a href="xxx"><div><li>1st div content</li></div></a>
		</span>
	</div>
	<div class="yo hey">
		<a href="xyz"><div class="cow sheep bunny"><h8>h8 content</h8></div></a>
	</div>
</body>
</html>
`

func main() {
	x, _ := goquery.Parse(strings.NewReader(example))
	x.Find("a div").Print()
	fmt.Println("---")
	x.Find("a div.cow").Print()
	
	fmt.Println("0 ++++++++++++++++++")
	fmt.Println(strings.Trim(x.Find("html head title").Html(), "\r\n\t ")) // 得到title,输出：the title of the page
	fmt.Println(x.Find("body div").Attr("custom_attr")) // 得到某个标签的属性值，输出：wow
	fmt.Println(x.Find("body div").Attr("id")) // 得到某个标签的属性值
	fmt.Println("1 ++++++++++++++++++")
	fmt.Println(x.Find("body div.yo").Html()) // 输出：<a href="xyz"><div class="cow sheep bunny"><h8>h8 content</h8></div></a>
	fmt.Println("11 ++++++++++++++++++")
	fmt.Println(x.Find("body div.yo").Text()) // 输出：h8 content
	fmt.Println("111 ++++++++++++++++++")
	fmt.Println(x.Find("body div").Html()) // ?????????????? 输出：
	fmt.Println("2 ++++++++++++++++++")
	fmt.Println(x.Find("a div li").Html()) // 输出：1st div content
	fmt.Println(x.Find("span a div li").Html())// 输出：1st div content
	fmt.Println(x.Find("div span a div li").Html())// 输出：1st div content
	fmt.Println(x.Find("body div span a div li").Html())// 输出：1st div content
	fmt.Println("3 ++++++++++++++++++")
	fmt.Println(x.Find("a div h8").Html()) // 输出：h8 content
	fmt.Println("4 ++++++++++++++++++")
	fmt.Println(x.Find("div").HasClass("yo"))// 输出：true
	fmt.Println("5 ++++++++++++++++++")
	fmt.Println(x.Find("").Attrs("href")) // 输出：[xxx xyz]
	fmt.Println("6 ++++++++++++++++++")
	fmt.Println(x.Find("body span h2").Html())// 输出：Yoyoyo
}