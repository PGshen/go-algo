package usepackage

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond // 创建全局条件变量

// 生产者
func producer(out chan<- int, idx int) {
	for {
		cond.L.Lock()       // 条件变量对应互斥锁加锁
		for len(out) == 3 { // 产品区满 等待消费者消费
			cond.Wait() // 挂起当前协程， 等待条件变量满足，被消费者唤醒
		}
		num := rand.Intn(1000) // 产生一个随机数
		out <- num             // 写入到 channel 中 （生产）
		fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(out))
		cond.L.Unlock()         // 生产结束，解锁互斥锁
		cond.Signal()           // 唤醒 阻塞的 消费者
		time.Sleep(time.Second) // 生产完休息一会，给其他协程执行机会
	}
}

//消费者
func consumer(in <-chan int, idx int) {
	for {
		cond.L.Lock()      // 条件变量对应互斥锁加锁（与生产者是同一个）
		for len(in) == 0 { // 产品区为空 等待生产者生产
			cond.Wait() // 挂起当前协程， 等待条件变量满足，被生产者唤醒
		}
		num := <-in // 将 channel 中的数据读走 （消费）
		fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", idx, num, len(in))
		cond.L.Unlock()                    // 消费结束，解锁互斥锁
		cond.Signal()                      // 唤醒 阻塞的 生产者
		time.Sleep(time.Millisecond * 500) //消费完 休息一会，给其他协程执行机会
	}
}
func main() {
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

/*
1. main函数中定义quit，其作用是让主协程阻塞。
2. 定义product作为队列，生产者产生数据保存至队列中，最多存储3个数据，消费者从中取出数据模拟消费
3. 条件变量要与锁一起使用，这里定义全局条件变量cond，它有一个属性：L Locker。是一个互斥锁。
4. 开启5个消费者协程，开启3个生产者协程。
5. producer生产者，在该方法中开启互斥锁，保证数据完整性。并且判断队列是否满，如果已满，调用wait()让该goroutine阻塞。当消费者取出数后执行cond.Signal()，会唤醒该goroutine，继续生产数据。
6. consumer消费者，同样开启互斥锁，保证数据完整性。判断队列是否为空，如果为空，调用wait()使得当前goroutine阻塞。当生产者产生数据并添加到队列，执行cond.Signal() 唤醒该goroutine。
*/
