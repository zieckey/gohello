package main

import (
    "bytes"
    "fmt"
    "net"
    "encoding/base64"
    "time"
    "net/http"
    "io/ioutil"
    "io"
    "os"
    "strconv"

    "github.com/PuerkitoBio/goquery"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
)

type Host struct {
    ip string
    name string
}

type Result struct {
    host Host
    cost []float64 // 成功的请求耗时情况
    successful int // 成功的次数
    fail int // 失败的次数
}

func UDP(v Host, res *Result) {
    res.host = v
    addr, err := net.ResolveUDPAddr("udp", v.ip+":53")
    if err != nil {
        fmt.Println("server address error. It MUST be a format like this hostname:port", err)
        return
    }

    start := time.Now()
    msg := ""
    data := make([]byte, 4096)

    // Create a udp socket and connect to server
    socket, err := net.DialUDP("udp4", nil, addr)
    if err != nil {
        fmt.Printf("connect to udpserver %v failed : %v", addr.String(), err.Error())
        return
    }
    defer socket.Close()
    socket.SetDeadline(start.Add(time.Duration(2) * time.Second))

    // send data to server
    senddata, err := base64.StdEncoding.DecodeString("DAsAAwABtncDaQYAmJSkC7lXtnS4O0yHqozqEKTOf5kID3UNO1h3eJh1CxM=")
    _, err = socket.Write(senddata)
    if err != nil {
        fmt.Println("send data error ", err)
        return
    }

    // recv data from server
    read, remoteAddr, err := socket.ReadFromUDP(data)
    if err != nil {
        if e, ok := err.(*net.OpError); ok && e.Timeout() {
            msg = "timeout"
            res.fail++
        } else {
            fmt.Printf("UDP cost %vs [%v], response data len:%v\n", time.Now().Sub(start).Seconds(), remoteAddr, read)
            res.fail++
        }
    } else {
        res.cost = append(res.cost, time.Now().Sub(start).Seconds())
        res.successful++
    }

    fmt.Printf("UDP cost %.3vs [%v] %v:%v, response data len:%v\n", time.Now().Sub(start).Seconds(), remoteAddr, v.name, msg, read)
}

func HTTP(v Host, res *Result) {
    url := "http://" + v.ip + "/status.html"
    res.host = v
    start := time.Now()
    msg := ""
    data, err := HttpGet(&http.Client{}, url, nil, time.Duration(2) * time.Second)
    if err != nil {
        msg = "timeout"
        res.fail++
    } else {
        res.cost = append(res.cost, time.Now().Sub(start).Seconds())
        res.successful++
    }

    fmt.Printf("HTTP cost %.3vs [%v] %v:%v, response data len:%v [%s]\n", time.Now().Sub(start).Seconds(), v.ip, v.name, msg, len(data), base64.StdEncoding.EncodeToString(data))
}

var udpResults []Result
var httpResults []Result

func main() {
    hosts := []Host{
        {"220.181.131.229", "北京电信"},
    }

    udpResults = make([]Result, len(hosts))
    httpResults = make([]Result, len(hosts))

    loop := 10
    if len(os.Args) == 2 {
        loop, _ = strconv.Atoi(os.Args[1])
    }

    for k := 0; k < loop; k++ {
        fmt.Printf("================ Round %d >>>\n", k)
        for i, v := range hosts {
            UDP(v, &udpResults[i])
        }

        for i, v := range hosts {
            HTTP(v, &httpResults[i])
        }
    }


    f, err := os.OpenFile("result.txt", os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0644)
    if err != nil {
        fmt.Printf("ERROR os.OpenFile(%v) failed : %v", "result.txt", err.Error())
        return
    }
    defer f.Close()
    w := io.MultiWriter(f, os.Stdout)

    w.Write([]byte(fmt.Sprintf("\n\n\n%v %v ============> Summary:\n", time.Now().String(), ip138())))
    for _, r := range udpResults {
        w.Write([]byte(fmt.Sprintf("UDP %v\n", r.String())))
    }

    for _, r := range httpResults {
        w.Write([]byte(fmt.Sprintf("HTTP %v\n", r.String())))
    }
}


func (r *Result) String() string {
    cost := 0.0
    for _, v := range r.cost {
        cost += v
    }
    return fmt.Sprintf("succ:%v fail:%v total:%v rate:%.3v cost:%.3v %v %v", r.successful, r.fail, r.successful+r.fail, float64(r.successful)/float64(r.successful+r.fail), cost/float64(len(r.cost)), r.host.ip, r.host.name)
}


var UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1541.0 Safari/537.36"

// HttpCall makes HTTP method call.
func HttpCall(client *http.Client, method, url string, header http.Header, body io.Reader, timeout time.Duration) (io.ReadCloser, error) {
    req, err := http.NewRequest(method, url, body)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", UserAgent)
    for k, vs := range header {
        req.Header[k] = vs
    }
    client.Timeout = timeout
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode == 200 {
        return resp.Body, nil
    }
    resp.Body.Close()
    if resp.StatusCode == 404 { // 403 can be rate limit error.  || resp.StatusCode == 403 {
        err = fmt.Errorf("resource not found: %s", url)
    } else {
        err = fmt.Errorf("%s %s -> %d", method, url, resp.StatusCode)
    }
    return nil, err
}

// HttpGet gets the specified resource.
// ErrNotFound is returned if the server responds with status 404.
func HttpGet(client *http.Client, url string, header http.Header, timeout time.Duration) ([]byte, error) {
    r, err := HttpCall(client, "GET", url, header, nil, timeout)
    if err != nil {
        return []byte(""), err
    }
    b, err := ioutil.ReadAll(r)
    r.Close()
    return b, err
}

func ip138() string {
    doc, err := goquery.NewDocument("http://1212.ip138.com/ic.asp")
    if err != nil {
        return fmt.Sprintf("HTTP GET ip138.com failed : %v", err.Error())
    }

    info := doc.Find("html body center")
    return gbk2utf8(info.Text())
}

func gbk2utf8(gbk string) string {
    in := bytes.NewReader([]byte(gbk))
    o := transform.NewReader(in, simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(o)
    if e != nil {
        return ""
    }
    return string(d)
}