package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
)

func Query(r *http.Request) string {
	id := r.FormValue("id")
	query := r.FormValue("query")

	//参数合法性检查

	//具体的业务逻辑，查询数据库/NoSQL等数据引擎，然后做逻辑计算，然后合并结果
	//这里简单抽象，直接返回欢迎语
	result := fmt.Sprintf("hello, %v", id)

	// 记录一条查询日志，用于离线统计和分析
	log.Printf("<id=%v><query=%v><result=%v><ip=%v>", id, query, result, r.RemoteAddr)

	return result
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	result := Query(r)
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/q", Handler)
	hostname, _ := os.Hostname()
	log.Printf("start http://%s:8091/q", hostname)
	log.Fatal(http.ListenAndServe(":8091", nil))
}

