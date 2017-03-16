package main

import (
	"os"
	"io/ioutil"
	"strings"
	"path/filepath"
	"log"
	"errors"
	"bytes"
	"strconv"
)

/*
	chart文件中，对数据哪一行进行处理：取平均值，然后追加到该行的末尾
 */
func main() {
	log.SetFlags(log.Lshortfile|log.Ltime)

	files, _ := LookupChartFiles(".")
	for _, file := range files {
		log.Printf("process %v", file)
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("read file failed %v : %v", file, err.Error())
			continue
		}

		new_buf := make([]byte, 0)
		out := bytes.NewBuffer(new_buf)
		lines := strings.Split(string(buf), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Data|") {
				line = strings.TrimSpace(line)
				line = strings.TrimRight(line, ",")
				rr := strings.Split(line, "=") // split "Data|muduo-1000=1283,1278,1271,1267,1284,1235,1236,1238,1248,1262,"
				nn := strings.Split(rr[1], ",")
				sum := 0
				count := 0
				for _, n := range nn {
					n = strings.TrimSpace(n)
					c, err := strconv.Atoi(n)
					if err != nil {
						log.Printf("ERROR process line [%v] n=[%v] failed : %v", line, n, err.Error())
						os.Exit(1)
					}
					sum += c
					count++
				}
				out.Write([]byte(line))
				out.Write([]byte(","))
				out.Write([]byte(strconv.Itoa(sum / count)))
			} else if strings.Contains(line, "XAxisNumbers") {
				line = strings.TrimSpace(line)
				out.Write([]byte(line))
				out.Write([]byte(","))
				out.Write([]byte("666"))
			} else if strings.Contains(line, "SubTitle") {
				line = strings.TrimSpace(line)
				line = strings.Replace(line, "number.", "number except the last one which is the average value of the privious 10 values.", -1)
				out.Write([]byte(line))
			} else {
				out.Write([]byte(line))
			}
			out.Write([]byte("\n"))
		}


		new_file := "new-" + file
		err = ioutil.WriteFile(new_file, out.Bytes(), 0644)
		if err != nil {
			log.Printf("ERROR WriteFile failed %v", new_file)
		}
	}
}


func LookupChartFiles(dir string) ([]string, error) {
	var files []string = make([]string, 0, 5)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if ok, err := filepath.Match("*.chart", f.Name()); err != nil {
			return err
		} else if ok {
			log.Printf("Add chart file %v", path)
			files = append(files, path)
		}
		return nil
	})

	if len(files) == 0 {
		return files, errors.New("Not found any *.chart files")
	}

	return files, err
}
