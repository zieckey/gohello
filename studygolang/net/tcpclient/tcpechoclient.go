package main

import (
	"flag"
	"net"
	"strings"
)

var serverIpPort = flag.String("h", "127.0.0.1:2007", "The tcp server ip address and listening port")
var closeWrite = flag.Bool("c", false, "close write after send all message")

/*
假如我是一个client，你是一个server/proxy
我的代码逻辑如下：
connect(); send(); shutdown(WRITE); recv(); close();
proxy会在client调用shutdown的时候epoll_wait返回，recv会返回0， errno是eof。这时候你会怎么处理？

服务器正确的处理逻辑请参考libevent-1.4的代码, libevent-1.4.14b-stable/http.c:evhttp_connection_done
*/
func main() {
	flag.Parse()
	serverAddr, err := net.ResolveTCPAddr("tcp", *serverIpPort)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		return
	}
	defer conn.Close()

	reply := make([]byte, 1024*128)
	sendn := 1024*127
	message := strings.Repeat("a", sendn)
	_, err = conn.Write([]byte(message))
	if err != nil {
		println("Write to server failed:", err.Error())
		return
	}
	if *closeWrite {
		conn.CloseWrite()
	}
	sum := 0
	for {
		n, err := conn.Read(reply)
		if err != nil {
			println("read from server failed:", err.Error())
			break
		}
		println("read ", n, " bytes from server")
		sum += n
		if sum == sendn {
			println("read all bytes done!")
			break
		}
	}
}
