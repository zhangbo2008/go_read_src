package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	var a = 3
	b := &a
	print(b)
	var locker = new(sync.Mutex)
	var cond = sync.NewCond(locker)
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知,阻塞当前goroutine
			fmt.Println(x)
		}(i)
	}

	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")

	select {}
}
