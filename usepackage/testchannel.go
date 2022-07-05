package usepackage

import (
	"fmt"
	"runtime"
	"time"
)

// 斐波那契
func Fibonacci() {
	ch := make(chan int)    // 用来进行数据通信的channel
	quit := make(chan bool) // 用来判断是否退出的channel

	go f(ch, quit)

	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}

func f(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			//return
			runtime.Goexit()
		}
	}
}

func timeout() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
				// 设置5秒读取不到数据就退出，避免阻塞
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				quit <- true
				break
			}
		}
	}()
	ch <- 666 // 写完5秒后退出
	<-quit
}
