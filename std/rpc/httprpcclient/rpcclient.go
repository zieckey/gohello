package main

import (
  "net/rpc"
  "fmt"
  "log"
  "os"
)

type Args struct {
  A,B int
}

type Quotient struct {
  Quo,Rem int
}

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: ",os.Args[0],"server")
	os.Exit(1)
  }
  serverAddress := os.Args[1]
  
  client,err := rpc.DialHTTP("tcp", serverAddress+":1234")
  if err != nil {
    log.Fatal("dialing:",err)
  }
  
  // Synchronous call
  args := Args{17,8}
  var reply int
  //远程调用，args是传给远程函数的参数，reply用来接收函数的结果
  err = client.Call("Arith.Multiply",args,&reply)
  if err != nil {
    log.Fatal("arith error:",err)
  }
  fmt.Printf("Arith.Multiply: %d*%d=%d\n",args.A,args.B,reply)

  var quot Quotient
  //远程调用，args是传给远程函数的参数，reply用来接收函数的结果
  err = client.Call("Arith.Divide",args, &quot)
  if err != nil {
    log.Fatal("arith error:",err)
  }
  fmt.Printf("Arith.Divide: %d/%d=%d remainder %d\n",args.A,args.B,quot.Quo,quot.Rem)
  
  
  hi := "你好"
  var resp string
  //远程调用，args是传给远程函数的参数，reply用来接收函数的结果
  err = client.Call("Arith.HandShake", hi, &resp)
  if err != nil {
    log.Fatal("arith error:",err)
  }
  fmt.Printf("Arith.HandShake: %v -> [%v]\n",  hi, resp)
}