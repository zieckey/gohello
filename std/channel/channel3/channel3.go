package main

// http://golanghome.com/post/554 select chan原子性的问题 

import (
	"fmt"
	"time"
)

var (
	ErrClosed = fmt.Errorf("worker closed")
)

type someTask struct {
}

func (s someTask) Delete() {
}

//下面代码能否保证在Stop被调用时，多个调用Send产生任务的线程不会因为向已close的chan发送内容而panic？

type Worker struct {
	quitCh chan struct{}
	taskCh chan someTask
}

func NewWorker() *Worker {
	w := &Worker{
		quitCh: make(chan struct{}),
		taskCh: make(chan someTask, 128),
	}
	return w
}

func (w *Worker) run() {
	for {
		select {
		case <-w.quitCh:
			close(w.taskCh)
			for t := range w.taskCh {
				t.Delete()
			}
			return

		case _ = <-w.taskCh:
		}
	}
}

func (w *Worker) Send(t someTask) error {
	// 问题在这里：由Stop关闭quitCh，再有run检测到quitCh已关闭而随后关闭taskCh，那这里是否一定能先检查到w.quitCh已关闭，而不会错误地向taskCh发送变量而panic？
	select {
	case <-w.quitCh:
		return ErrClosed

	case w.taskCh <- t:
	}
	return nil
}

func (w *Worker) Stop() {
	close(w.quitCh)
}

func main() {
	w := NewWorker()

	// 多个生产线程
	for i := 0; i < 100; i++ {
		go func() {
			for {
				if err := w.Send(someTask{}); err != nil {
					break
				}
			}
		}()
	}

	// 某个时刻需要结束w的工作
	go func() {
		time.Sleep(10 * time.Second)
		w.Stop()
	}()

	w.run()

	// 其他任务不会导致程序结束
	time.Sleep(time.Second * 2)
}
