/**
并发
Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的
*/
package concurrent

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func Demo() {

	num := runtime.NumCPU()
	fmt.Println("获取主机的逻辑CPU个数:" + strconv.Itoa(num))
	runtime.GOMAXPROCS(1) //设置可同时执行的最大CPU数

	//并发基础示例
	for i := 0; i < 10; i++ {
		go Say(i)
	}
	//这里如果不使用time.Sleep, main函数很快就执行完了
	time.Sleep(1 * time.Second)

	//并发和channel结合使用
	CSP()

	//同步goroutine实现方式
	var goSync sync.WaitGroup //声明一个WaitGroup变量
	for i := 0; i < 10; i++ {
		goSync.Add(1) // WaitGroup的计数加1
		go Cal(i, i+1, &goSync)
	}
	goSync.Wait() //等待所有goroutine执行完毕

	//通过channel实现goroutine之间的同步
	exitChan := make(chan bool, 10) //声明并分配管道内存
	for i := 0; i < 10; i++ {
		go Add(i, i+1, exitChan)
	}
	for j := 0; j < 10; j++ {
		<-exitChan //取信号数据，如果取不到则会阻塞
	}
	close(exitChan) // 关闭管道

	//goroutine之间的通讯
	dataChan := make(chan int, 100) //通讯数据管道
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go Producer(dataChan, i, &wg) //生产数据
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(dataChan, &wg) //消费数据
		wg.Add(1)
	}
	wg.Wait()

}

func Say(i int) {
	fmt.Println(strconv.Itoa(i) + ":Hello World!")
}

func CSP() {
	message := make(chan string)
	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)
}

func Cal(a, b int, n *sync.WaitGroup) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	//goroutinue完成后, WaitGroup的计数-1
	defer n.Done()
}

func Add(a, b int, exitChan chan bool) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	time.Sleep(time.Second * 2)
	exitChan <- true
}

func Producer(dataChan chan int, data int, wait *sync.WaitGroup) {
	dataChan <- data
	fmt.Println("生产数据:", data)
	wait.Done()
}

func Consumer(dataChan chan int, wait *sync.WaitGroup) {
	a := <-dataChan
	fmt.Println("消费数据:", a)
	wait.Done()
}
