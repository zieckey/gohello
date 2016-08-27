package main

import (
    "fmt"
    "net"
    "os"
    "encoding/base64"
)

func main() {
    hostport := "10.16.28.17:1053"
    if len(os.Args) == 2 {
        hostport = os.Args[1]
    }

    addr, err := net.ResolveUDPAddr("udp", hostport)
    if err != nil {
        fmt.Println("server address error. It MUST be a format like this hostname:port", err)
        return
    }

    // Create a udp socket and connect to server
    socket, err := net.DialUDP("udp4", nil, addr)
    if err != nil {
        fmt.Printf("connect to udpserver %v failed : %v", addr.String(), err.Error())
        return
    }
    defer socket.Close()

    b64 := "DgYkMAEIpoTHxgABAAC9Syk0fi/rgUhyxw7iWPJGOx4MEe3SQa5huZgwPV6VTQTQrRfPeAAkOBptmY9+jAwCmApDgi9Man/JZewdPwRa/OZCY6UvcF66wBoavDwc+Q==" // Tutorial v6, net_method=0
    senddata, err := base64.StdEncoding.DecodeString(b64)
    _, err = socket.Write(senddata)
    if err != nil {
        fmt.Println("send data error ", err)
        return
    }

    // recv data from server
    data := make([]byte, 4096)
    read, remoteAddr, err := socket.ReadFromUDP(data)
    if err != nil {
        fmt.Println("recv data error ", err)
        return
    }

    fmt.Printf("server addr [%v], response data len:%v [%s]\n", remoteAddr, read, string(data[:read]))
}
