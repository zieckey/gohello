package main

import (
	"fmt"
	"time"
	//"github.com/astaxie/beego/orm"
	"github.com/fanliao/go-promise"
)

type User struct {
	Id   int
	Name string
}

//func main() {
//	ch := make(chan bool)
//	p := promise.NewPromise() //建立Promise对象
//	go func() {
//		o := orm.NewOrm()
//		u := User{Id: 1}
//
//		//读取一个User
//		if err := o.Read(&u); err != nil {
//			p.Reject(err) //成功
//			fmt.Printf("Reject\n")
//		} else {
//			p.Resolve(u) //失败
//			fmt.Printf("Resolve\n")
//		}
//		ch <- true
//	}()
//
//	<- ch
//}

func main() {
	p := promise.NewPromise()
	go func() {
		time.Sleep(500 * time.Millisecond)
		p.Resolve("okxxx")
	}()
	t0 := time.Now()
	r, err := p.Get()
	fmt.Printf("Waited %v for goroutine to stop. result [%v] err=%v\n", time.Since(t0), r, err) // Output : Waited 500.05ms for goroutine to stop. result [okxxx] err=<nil>
}
