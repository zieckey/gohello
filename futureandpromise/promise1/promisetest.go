package main

import (
	"errors"
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
	testResolve()
	testReject()
	testCancel()
	testCallbacks()
}

func testResolve() {
	p := promise.NewPromise()
	go func() {
		time.Sleep(500 * time.Millisecond)
		p.Resolve("okxxx")
	}()
	t0 := time.Now()
	r, err := p.Get()
	fmt.Printf("Waited %v result [%v] err=%v\n", time.Since(t0), r, err) // Output : Waited 500.05ms result [okxxx] err=<nil>
}

func testReject() {
	p := promise.NewPromise()
	go func() {
		time.Sleep(500 * time.Millisecond)
		p.Reject(errors.New("failed"))
	}()
	t0 := time.Now()
	r, err := p.Get()
	fmt.Printf("Waited %v result [%v] err=%v\n", time.Since(t0), r, err) // Output : Waited 500.05ms result [<nil>] err=failed
}

func testCancel() {
	p := promise.NewPromise()
	go func() {
		time.Sleep(500 * time.Millisecond)
		p.Cancel()
	}()
	t0 := time.Now()
	r, err := p.Get()
	fmt.Printf("Waited %v result [%v] err=%v\n", time.Since(t0), r, err) // Output : Waited 500.05ms result [<nil>] err=Task be cancelled
}

func testCallbacks() {
	timeout := 50 * time.Millisecond
	done, always, fail := false, false, false

	p := promise.NewPromise()
	go func() {
		<- time.After(timeout)
		p.Resolve("ok")
	}()

	p.OnSuccess(func(v interface{}) {
		done = true
		fmt.Printf("The argument of Done should be 'ok'\n")
	}).OnComplete(func(v interface{}) {
		always = true
		fmt.Printf("The argument of Always should be 'ok'\n")
	}).OnFailure(func(v interface{}) {
		fail = true
		panic("Unexpected calling")
	})
	t0 := time.Now()
	r, err := p.Get()

	//The code after Get() and the callback will be concurrent run
	//So sleep 52 ms to wait all callback be done
	time.Sleep(timeout)

	fmt.Printf("Waited %v r=%v err=%v done=%v always=%v fail=%v\n", time.Since(t0), r, err, done, always, fail)
}
