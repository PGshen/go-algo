package usepackage

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_consumer(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	quit := make(chan bool)          // 创建用于结束通信的 channel

	product := make(chan int, 3) // 产品区（公共区）使用channel 模拟
	cond.L = new(sync.Mutex)     // 创建互斥锁和条件变量

	for i := 0; i < 5; i++ { // 5个消费者
		go producer(product, i+1)
	}
	for i := 0; i < 3; i++ { // 3个生产者
		go consumer(product, i+1)
	}
	<-quit // 主协程阻塞 不结束
}
