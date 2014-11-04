package net

import (
	"net"
	"fmt"
	"strconv"
)


////////////////////////////////////////////////////////
//
//错误检查
//
////////////////////////////////////////////////////////
func checkError(err error,info string) (res bool) {
	
	if(err != nil){
		fmt.Println(info+"  " + err.Error())
		return false
	}
	return true
}

////////////////////////////////////////////////////////
//
//服务器端接收数据线程
//参数：
//		数据连接 conn
//
////////////////////////////////////////////////////////
func Handler(conn net.Conn){
	
	fmt.Println("connection is up, from ...", conn.RemoteAddr().String())
	
	buf := make([]byte,1024)
	for{
		lenght, err := conn.Read(buf)
		if(checkError(err,"Connection")==false){
			conn.Close()
			break
		}
		if lenght > 0{
			buf[lenght]=0
		}
		fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		conn.Write(buf)
	}

}

func TcpServer(port int) {
	service:= string(":") + strconv.Itoa(port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  	if err != nil {
  		fmt.Errorf("%v", err.Error)
  		return
  	}
  	
  	listener, err := net.ListenTCP("tcp4", tcpAddr)
  	for {
  		fmt.Printf("Listening at %d ...", port)
  		conn, err := listener.Accept()
  		checkError(err, "AcceptCheck")
  		fmt.Println("Accepting ...") 		
  		//启动一个新线程
		go Handler(conn)
  	}
	
}
