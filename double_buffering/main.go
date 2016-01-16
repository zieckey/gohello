package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"bytes"
	"strings"
	"io"
)

var blackIDs map[string]int
func LoadBlackIDs(filepath string) error {
	// 加载黑名单列表文件，每行一个
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	r := bytes.NewBuffer(b)
	for {
		id, err := r.ReadString('\n')
		if err == io.EOF || err == nil {
			id = strings.TrimSpace(id)
			if len(id) > 0 {
				blackIDs[id] = 1
			}
		}

		if err != nil {
			break
		}
	}

	return nil
}

func IsBlackID(id string) bool {
	_, exist := blackIDs[id]
	return exist
}

func Query(r *http.Request) (string, error) {
	id := r.FormValue("id")
	query := r.FormValue("query")

	//参数合法性检查

	if IsBlackID(id) {
		return "ERROR", fmt.Errorf("ERROR id")
	}

	//具体的业务逻辑，查询数据库/NoSQL等数据引擎，然后做逻辑计算，然后合并结果
	//这里简单抽象，直接返回欢迎语
	result := fmt.Sprintf("hello, %v", id)

	// 记录一条查询日志，用于离线统计和分析
	log.Printf("<id=%v><query=%v><result=%v><ip=%v>", id, query, result, r.RemoteAddr)

	return result, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	result, err := Query(r)
	if err == nil {
		w.Write([]byte(result))
	} else {
		w.WriteHeader(403)
		w.Write([]byte(result))
	}
}

func main() {
	blackIDs = make(map[string]int)
	if len(os.Args) == 2 {
		err := LoadBlackIDs(os.Args[1])
		panic(err)
	}

	http.HandleFunc("/q", Handler)
	hostname, _ := os.Hostname()
	log.Printf("start http://%s:8091/q", hostname)
	log.Fatal(http.ListenAndServe(":8091", nil))
}

