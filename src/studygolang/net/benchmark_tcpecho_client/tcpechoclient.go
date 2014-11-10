package main

import (
	"flag"
)

var localIp = flag.String("localIp", "0.0.0.0", "The local ip address")
var localIpCount = flag.Int("localIpCount", 1, "The total count of local ip")
var connPerIp = flag.Int("connPerIp", 1, "The concurrenc count of every local ip")
var serverIp = flag.String("serverIp", "127.0.0.1", "The tcp server ip address")
var serverPort = flag.Int("serverPort", 2007, "The tcp server listening port")
var messageLen = flag.Int("messageLen", 26, "The length of the message sending to server")

func main() {
	
}


