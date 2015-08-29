package main

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/PuerkitoBio/fetchbot"
)

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(handler))
	f.CrawlDelay = 20 * time.Second
	f.UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1541.0 Safari/537.36"
	queue := f.Start()
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.SendStringGet("http://360.cn", "http://www.so.com", "http://hao.360.cn")
	queue.Close()
}

func handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())
}
