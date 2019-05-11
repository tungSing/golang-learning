/**
异常处理
Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。

异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

panic：
1、内置函数
2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
4、直到goroutine整个退出，并报告错误

recover：
1、内置函数
2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
3、一般的调用建议
	a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
	b). 可以获取通过panic传递的error

注意:

1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

如何区别使用 panic 和 error 两种方式?

惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。
*/

package exception

import (
	"fmt"
	"log"
)

func Demo() {
	//向已关闭的通道发送数据会引发panic
	test()
	//如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执
	ProtectCode(5, 6)
	//实现类似 try catch 的异常处理
	Biz()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

func ProtectCode(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()

		panic("出现意想不到到错误")

		z = x / y
		return
	}()

	fmt.Printf("x/y = %d\n", z)
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func Biz() {
	Try(func() {
		//业务逻辑
		panic("出现意想不到的错误")
	}, func(err interface{}) {
		log.Println(err)
	})
}
