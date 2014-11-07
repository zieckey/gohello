package main

import (
  "fmt"
  "net/rpc"
  "errors"
  "net/http"
)

type Args struct {
  A,B int
}

type Quotient struct {
  Quo,Rem int
}

type Arith struct {
	
}

// 需要远程调用的方法
func (t *Arith) Multiply(args *Args, reply *int) error {
  *reply = args.A * args.B
  fmt.Printf("Multiply : %v * %v = %v\n",  args.A, args.B, *reply)
  return nil
}
// 需要远程调用的方法
func (t *Arith) Divide(args *Args,quo *Quotient) error {
  if args.B == 0 {
    return errors.New("divide by zero")
  }
  quo.Quo = args.A / args.B
  quo.Rem = args.A % args.B
  fmt.Printf("Divide : %v / %v = %v|%v\n",  args.A, args.B, quo.Quo, quo.Rem)
  return nil
}

func (t *Arith) HandShake(sayhi *string, resp *string) error {
	*resp = "reponsed message: " + *sayhi
	fmt.Printf("HandShake : %v -> [%v]\n",  *sayhi, *resp)
	return nil
} 

func main() {
  arith := new(Arith)
  // 注册
  rpc.Register(arith)
  rpc.HandleHTTP()
  
  err := http.ListenAndServe(":1234",nil)
  if err != nil {
    fmt.Println(err.Error())
  }
}