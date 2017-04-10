package main

import (
    "os"
    "fmt"
    "net"
    "bufio"
    "strings"
)

func send(conn *net.TCPConn, ch chan string) {
    for {
        s := <-ch
        if s == "quit" {
            os.Exit(0)
        }
        conn.Write([]byte(s))
    }
}

func recv(conn *net.TCPConn) {
    for {
        b := make([]byte, 1024)
        n, err := conn.Read(b)
        if err == nil {
            fmt.Printf("%v\n", string(b[0:n]))
        }
    }
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage: %v host port\n", os.Args[0])
        return
    }

    addr, err := net.ResolveTCPAddr("tcp", os.Args[1] + ":" + os.Args[2])
    conn, err := net.DialTCP("tcp", nil, addr)
    if err != nil {
        fmt.Printf("Cannot connect to %v:%v\n", os.Args[1], os.Args[2])
        return
    }

    ch := make(chan string)
    go send(conn, ch)
    go recv(conn)

    r := bufio.NewReader(os.Stdin)
    for {
        s, err := r.ReadString('\n')
        if err != nil {
            fmt.Printf("Read from STDIN error %v", err.Error())
            continue
        }

        if s = strings.TrimSpace(s); s == "quit" {
            break
        }
        ch <- s
    }
}
