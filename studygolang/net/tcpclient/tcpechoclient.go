package main

import (
	"flag"
	"net"
	"strings"
	"log"
	"time"
)

var serverIpPort = flag.String("h", "127.0.0.1:2007", "The tcp server ip address and listening port")
var closeWrite = flag.Bool("c", false, "close write after send all message")
var messageLen = flag.Int("l", 1024*127, "length of the message sending to server")

/*
假如我是一个client，你是一个server/proxy
我的代码逻辑如下：
connect(); send(); shutdown(WRITE); recv(); close();
proxy会在client调用shutdown的时候epoll_wait返回，recv会返回0， errno是eof。这时候你会怎么处理？

服务器正确的处理逻辑请参考libevent-1.4的代码, libevent-1.4.14b-stable/http.c:evhttp_connection_done
*/
func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	flag.Parse()
	serverAddr, err := net.ResolveTCPAddr("tcp", *serverIpPort)
	if err != nil {
		log.Println("ResolveTCPAddr failed:", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Println("Dial failed:", err.Error())
		return
	}
	defer conn.Close()

	reply := make([]byte, 1024*128)
	message := strings.Repeat("a", *messageLen)
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("Write to server failed:", err.Error())
		return
	}
	if *closeWrite {
		conn.CloseWrite()
	}
	sum := 0
	for {
		n, err := conn.Read(reply)
		if err != nil {
			log.Println("read from server failed:", err.Error())
			break
		}
		log.Println("read ", n, " bytes from server")
		sum += n
		if sum == *messageLen {
			log.Println("read all bytes done!")
			break
		}
	}
	
	time.Sleep(5*time.Second)
}
