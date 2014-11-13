package pkg

import (
	"fmt"
	"io"
	"os"
)

func TestFileReadAndWrite() {

	path := "TestFileReadAndWrite.exe"
	tty, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0)
	defer tty.Close()
	if err != nil {
		fmt.Printf("open file failed [%v] error=[%v]\n", path, err)
		return
	}
	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	s := "abc"
	b := []byte(s)
	w.Write(b)
	tty.Sync()

	//TODO 读取文件失败，为什么？？？
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Printf("read file failed [%v] error=[%v] readn=%d\n", path, err, n)
		return
	}

	s = string(buf)
	fmt.Printf("readn=%d s.len=%d content:[%s]\n", n, len(s), s)
}
