/**
流程控制
条件语句if switch select
循环语句for range
循环控制Goto break commit
*/
package decisionMaking

import (
	"fmt"
	"go/types"
)

func Demo() {

	//if语句
	var a, b = 10, 11
	if a < 20 {
		fmt.Println("a小于20")
	} else {
		fmt.Println("a大于等于20")
	}

	if b < 11 {
		fmt.Println("b小于11")
	} else if b == 11 {
		fmt.Println("b等于11")
	} else {
		fmt.Println("b大于11")
	}

	//switch语句
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(grade)

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B" || grade == "C":
		fmt.Println("良")

	default:
		fmt.Println("差")
	}

	var x interface{}
	switch i := x.(type) {
	case types.Nil:
		fmt.Printf("x的类型为%T\r\n", i)
	case int:
		fmt.Printf("x的类型为int")
	default:
		fmt.Printf("未知类型\n")
	}

	/**
	select 语句 主要用于处理异步IO操作
	每个case都必须是一个通信
	所有channel表达式都会被求值
	所有被发送的表达式都会被求值
	如果任意某个通信可以进行，它就执行；其他被忽略。
	如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	否则：
	如果有default子句，则执行该语句。
	如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select { //不停的在这里检测
	case i1 = <-c1: //检测有没有数据可以读
		//如果chanl成功读取到数据，则进行该case处理语句
		fmt.Printf("received ", i1, "from c1\n")
	case c2 <- i2: // //检测有没有可以写
		//如果成功向chan2写入数据，则进行该case处理语句
		fmt.Printf("sent ", i2, "to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
		//假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞
		// 一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
	default:
		//如果以上都没有符合条件，那么则进行default处理流程
		fmt.Printf("no communication\n")
	}

	//for 循环
	//方式一
	fmt.Println("---for 方式一")
	s := "abcd"
	for i, n := 0, len(s); i < n; i++ {
		fmt.Println(i, string(s[i]))
	}

	//方式二
	fmt.Println("---for 方式二")
	k := len(s) - 1
	for k >= 0 {
		fmt.Println(string(s[k]))
		k--
	}

	//方式三 无限循环
	//for {
	//	fmt.Println(s)
	//}

	//range 循环语句 类似迭代器
	fmt.Println("---range循环语句")
	for i := range s {
		fmt.Println(s[i])
	}

	/**
	for 和 for range有什么区别?

	主要是使用场景不同

	for可以

	遍历array和slice
	遍历key为整型递增的map
	遍历string

	for range可以完成所有for可以做的事情，却能做到for不能做的，包括

	遍历key为string类型的map并同时获取key和value
	遍历channel
	*/

	//range会复制对象

	duplication := [3]int{0, 1, 2}
	// index、value 都是从复制品中取出。
	for i, v := range duplication {
		if i == 0 {
			// 在修改前，我们先修改原数组。
			duplication[1], duplication[2] = 999, 999
			fmt.Println(duplication)
		}

		// 使用复制品中取出的 value 修改原数组。
		duplication[i] = v + 100
	}

	fmt.Println(duplication)

	// 引用类型，其底层数据不会被复制
	reference := []int{1, 2, 3, 4, 5}

	// 复制 struct slice { pointer, len, cap }
	for i, v := range reference {
		if i == 0 {
			// 对 slice 的修改，不会影响 range。
			reference = reference[:3]
			// 对底层数据的修改
			reference[2] = 100
		}

		fmt.Println(i, v)
	}

	/**
	循环控制语句Goto、Break、Continue
		1.三个语句都可以配合标签(label)使用
	    2.标签名区分大小写，定以后若不使用会造成编译错误
	    3.continue、break配合标签(label)可用于多层循环跳出
	    4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同
	*/

	fmt.Println("---goto 跳转---")
	var jump int
	for {
		fmt.Println(jump)
		jump++
		if jump > 2 {
			goto BREAK
		}
	}
	//goto语句与标签之间不能有变量声明。否则编译错误
	//var ss int = 1
BREAK:
	fmt.Println("break")
	//fmt.Println(ss)

	//跳转语句 goto语句可以跳转到本函数内的某个标签
	gotoCount := 0
GotoLabel:
	gotoCount++
	if gotoCount < 10 {
		//如果小于10的话就跳转到GotoLabel
		goto GotoLabel
	}
	fmt.Println(gotoCount)

	/**
	1.用于循环语句中跳出循环，并开始执行循环之后的语句。
	2.break在switch（开关语句）中在执行一条case后跳出语句的作用。
	3.break 可用于 for、switch、select
	*/
	fmt.Println("---break ---")

	var breakDemo int = 10
	for breakDemo < 20 {
		fmt.Printf("breakDemo 的值为 : %d\n", breakDemo)
		breakDemo++
		if breakDemo > 15 {
			//使用 break 语句跳出循环
			break
		}
	}

	fmt.Println("---label break ---")
Exit:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i+j > 15 {
				fmt.Println("exit")
				break Exit
			}
		}
	}

	fmt.Println("3")

	/**
	  continue 不是跳出循环，而是跳过当前循环执行下一次循环语句
	  continue 仅能用于 for 循环
	*/
	fmt.Println("---continue ---")

	continueDemo := 5
	for continueDemo < 10 {
		continueDemo++
		if continueDemo == 7 {
			continue
		}
		fmt.Printf("continueDemo 的值为 : %d\n", continueDemo)

	}
}
