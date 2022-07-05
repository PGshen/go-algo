package usepackage

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex // 定义互斥锁变量 mutex

func printer(str string) {
	mutex.Lock()         // 添加互斥锁
	defer mutex.Unlock() // 使用结束时解锁

	for _, data := range str { // 迭代器
		fmt.Printf("%c", data)
		time.Sleep(time.Second) // 放大协程竞争效果
	}
	fmt.Println()
}

func person1(s1 string) {
	printer(s1)
}

func person2() {
	printer("world") // 调函数时传参
}

func runLock() {
	go person1("hello") // main 中传参
	go person2()
	time.Sleep(time.Second * 10)
}

var count int           // 全局变量count
var rwlock sync.RWMutex // 全局读写锁 rwlock

func read(n int) {
	rwlock.RLock()
	fmt.Printf("读 goroutine %d 正在读取数据...\n", n)
	num := count
	fmt.Printf("读 goroutine %d 读取数据结束，读到 %d\n", n, num)
	defer rwlock.RUnlock()
}
func write(n int) {
	rwlock.Lock()
	fmt.Printf("写 goroutine %d 正在写数据...\n", n)
	num := rand.Intn(1000)
	count = num
	fmt.Printf("写 goroutine %d 写数据结束，写入新值 %d\n", n, num)
	defer rwlock.Unlock()
}

func runRwLock() {
	for i := 0; i < 5; i++ {
		go read(i + 1)
	}
	for i := 0; i < 5; i++ {
		go write(i + 1)
	}
	time.Sleep(time.Second * 3)
}
