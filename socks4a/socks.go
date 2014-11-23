package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
)

//import "sync"

const bindAddr = ":2011"
const defaultBufferSize = 4096
const maxConn = 0x10000

type Tunnel struct {
	lconn net.Conn // the connection from local app
	rconn net.Conn // the connection to remote server
	lbuf  *bufio.Reader // refactor : don't use this buffered io
	//closed bool
}

func socks4a(ipBuf []byte) bool {
	if ipBuf[0] == 0 &&
		ipBuf[1] == 0 &&
		ipBuf[2] == 0 &&
		ipBuf[3] != 0 {
		return true
	}

	return false
}

func servRemoteTunnel(t *Tunnel) {
	buf := make([]byte, defaultBufferSize)
	for {
		n, err := t.rconn.Read(buf)
		log.Printf("read %v bytes from remote server %v [%v]", n, t.rconn.RemoteAddr(), string(buf))
		if n == 0 {
			t.lconn.Close()
			t.rconn.Close()
			return
		}
		
		if err != nil {
			log.Printf("Read from remote server %v", err)
			return
		}

		t.lconn.Write(buf[:n])
	}
}

func refuse(t *Tunnel) {
	buf := []byte{0, 0x5b, 0, 0, 0, 0, 0, 0}
	t.lconn.Write(buf)
	t.lconn.Close()
}

func grant(t *Tunnel) {
	buf := []byte{0, 0x5a, 0, 0, 0, 0, 0, 0}
	t.lconn.Write(buf)
}

func servLocalTunnel(lconn net.Conn) {
	t := new(Tunnel)
	t.lconn = lconn
	t.lbuf = bufio.NewReader(lconn)
	buf := make([]byte, defaultBufferSize)
	bufSize := 0
	for {
		n, _ := t.lbuf.Read(buf[bufSize:])
		log.Printf("read %v bytes from local app %v bytes : [%s]", n, t.lconn.RemoteAddr(), string(buf))
		if n == 0 {
			t.lconn.Close()
			return
		}

		if t.rconn != nil {
			log.Printf("write %v bytes to remote server %v", n, t.rconn.RemoteAddr())
			t.rconn.Write(buf[:n+bufSize])
			bufSize = 0
			continue
		}

		// connecting to remote server
		if n+bufSize > 8 {
			end := bytes.IndexByte(buf[8:], byte(0))
			if end < -1 {
				log.Printf("cannot find '\\0', need to read more data")
				continue // need to read more data
			}
			ver := buf[0]
			cmd := buf[1]
			port := int(buf[2]<<8) + int(buf[3])
			ip := net.IPv4(buf[4], buf[5], buf[6], buf[7])
			log.Printf("ver=%v cmd=%v remote addr %v:%v\n", ver, cmd, ip, port)
			if socks4a(buf[4:]) {
				// TODO get the remote ip
			}

			a := &net.TCPAddr{IP: ip, Port: port}
			c, err2 := net.DialTCP("tcp", nil, a)
			if err2 != nil {
				log.Printf("DialTCP", err2)
				refuse(t)
				return
			}
			log.Printf("has connected to remote %s", a)
			t.rconn = c
			go servRemoteTunnel(t)
			bufSize = 0
			grant(t)
		}
	}
}

func start(addr string) {
	a, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	l, err2 := net.ListenTCP("tcp", a)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("socks server bind %s", a)
	for {
		c, err3 := l.Accept()
		if err3 != nil {
			log.Println(err)
			continue
		}
		log.Printf("a connection came from %v", c.RemoteAddr().String())
		go servLocalTunnel(c)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	start(bindAddr)
}
