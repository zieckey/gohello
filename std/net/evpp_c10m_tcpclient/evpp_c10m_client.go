package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var localIpPort = flag.String("localIpPort", "0.0.0.0:0", "The local ip address and port to bind")
var localIpCount = flag.Int("localIpCount", 1, "The total count of local ip")
var connPerIp = flag.Int("connPerIp", 1, "The concurrent connection count for every local ip")
var serverIpPort = flag.String("serverIpPort", "127.0.0.1:2007", "The tcp server ip address and listening port")
var messageLen = flag.Int("messageLen", 26, "The length of the message sending to server")
var sleepIntervalMs = flag.Int("sleepIntervalMs", 1000, "The sleeping interval time between message sending on one connection")

func main() {
	flag.Parse()
	for ipIndex := 0; ipIndex < *localIpCount; ipIndex++ {
		for i := 0; i < *connPerIp; i++ {
			lipp := *localIpPort
			if lipp != "0.0.0.0:0" {
				lipp = calcIpPort(lipp, ipIndex, 0)
			}
			go connect(*serverIpPort, lipp, *messageLen, *sleepIntervalMs)
		}
	}
	ch := make(chan int)
	<-ch
}

func connect(serverIpPort string, localIpPort string, messageLen int, sleepIntervalMs int) {
	serverAddr, err := net.ResolveTCPAddr("tcp", serverIpPort)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	var localAddr *net.TCPAddr
	if localIpPort != "0.0.0.0:0" {
		localAddr, err = net.ResolveTCPAddr("tcp", localIpPort)
	}

	conn, err := net.DialTCP("tcp", localAddr, serverAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		return
	}

	defer conn.Close()
	fmt.Printf("Connected to %v OK from %v\n", serverAddr, localAddr)

	reply := make([]byte, 1024*128)
	message := []byte(strings.Repeat("a", messageLen))
	lenBuff := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuff, uint32(messageLen))
	for {
		_, err = conn.Write(lenBuff)
		_, err = conn.Write(message)
		if err != nil {
			println("Write to server failed:", err.Error())
			return
		}

		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			return
		}

		time.Sleep(time.Duration(sleepIntervalMs) * time.Millisecond)
	}
}

// 根据index计算当前ip的下一个IP
// 例如输入 "192.168.0.150:80", ipIndex=2 ===> "192.168.0.152:80"
func calcIpPort(ipPort string, ipIndex int, portIndex int) string {
	spp := strings.Split(ipPort, ":")
	a, _ := strconv.Atoi(string(spp[1]))
	port := strconv.Itoa(a + portIndex)
	dotip := strings.Split(spp[0], ".")
	a, _ = strconv.Atoi(string(dotip[3]))
	next := strconv.Itoa(a + ipIndex)
	r := dotip[0] + "." + dotip[1] + "." + dotip[2] + "." + next + ":" + port
	return r
}
