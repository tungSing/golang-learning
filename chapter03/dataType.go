/**
数据类型
*/
package dataType

import (
	"fmt"
	"unsafe"
)

func DataType() {
	//布尔型
	var b bool = true
	fmt.Println(b)

	//数字型(Numeric types)
	//无符号8位整型(0到255)
	var numUint8 uint8 = 255
	fmt.Println(numUint8)

	//无符号 16 位整型(0 到65535)
	var numUint16 uint16 = 65535
	fmt.Println(numUint16)

	//无符号 32 位整型(0 到 4294967295)
	var numUint32 uint32 = 4294967295
	fmt.Println(numUint32)

	//无符号 64 位整型(0 到 18446744073709551615)
	var numUint64 uint64 = 18446744073709551615
	fmt.Println(numUint64)

	//有符号 8 位整型 (-128 到 127)
	var numInt8 int8 = 127
	fmt.Println(numInt8)

	//有符号 16 位整型 (-32768 到 32767)
	var numInt16 int32 = 32767
	fmt.Println(numInt16)

	//有符号 32 位整型 (-2147483648 到 2147483647)
	var numInt32 int32 = 2147483647
	fmt.Println(numInt32)

	//有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	var numInt64 int64 = 9223372036854775807
	fmt.Println(numInt64)

	//取决于操作系统（32位机器上就是32字节，64位机器上就是64字节）
	var numInt int = 11
	fmt.Println(numInt)

	//取决于操作系统（32位机器上就是32字节，64位机器上就是64字节）
	var numUint uint = 11
	fmt.Println(numUint)

	//IEEE-754 32位浮点型数
	var numFloat32 float32 = 12.22
	fmt.Println(numFloat32)

	//IEEE-754 64位浮点型数
	var numFloat64 float64 = 12.22
	fmt.Println(numFloat64)

	//32 位实数和虚数
	var numComplex64 complex64 = 1
	fmt.Println(numComplex64)

	//64 位实数和虚数
	var numComplex128 complex128 = 11
	fmt.Println(numComplex128)

	//byte就是unit8的别名
	var numByte byte = 123
	fmt.Println(numByte)

	//rune就是int32的别名
	var numRune rune = 123
	fmt.Println(numRune)

	//无符号整型，用于存放一个指针
	var numUintptr uintptr = uintptr(unsafe.Pointer(&numRune))
	fmt.Println(numUintptr)

	//字符串类型
	var str string = "hello world"
	fmt.Println(str)

	//派生类型(复合类型)-array，struct，指针，function，interface，slice，map，channel类型（可以使用type构造）
	/**
	指针类型 *号用于指定变量是作为一个指针
	1. 第一指针变量
	2. 为指针变量赋值
	3. 访问指针变量中指向地址的值
	*/
	var aa int = 20      //声明实际变量
	var pointerType *int // 声明指针变量
	pointerType = &aa    //指针变量的存储地址
	//指针变量的存储地址
	fmt.Printf("aa变量的地址:%x\n", &aa)
	//指针变量的存储地址
	fmt.Printf("pointerType变量的地址:%x\n", pointerType)
	//使用指针访问值
	fmt.Printf("*pointerType变量的值: %d\n", *pointerType)

	//数组类型
	var array = [5]int{0, 1, 2, 3, 4}
	fmt.Println(array)

	//结构化类型(struct)
	type person struct {
		id   int
		name string
	}
	var zhangSan = new(person)
	zhangSan.id = 1
	zhangSan.name = "张三"
	fmt.Println(zhangSan)

	/**
	管道类型(channel)
		a. 类似unix中管道(pipe)
	    b. 先进先出
	    c. 线程安全，多个goroutine同时访问，不需要加锁
	    d. channel是有类型的，一个整数的channel只能存放整数
	*/
	//定义变量
	var chanName chan int
	//变量赋值
	chanName = make(chan int, 10)
	var ch02 chan int = make(chan int)
	fmt.Printf("有缓冲 局部变量 chan chanName : %v\n", chanName)
	fmt.Printf("无缓冲 局部变量 chan ch2 : %v\n", ch02)

	//函数类型
	//本方法就是一个简单的函数

	//切片类型(slice)
	var arr = [...]int{0, 1, 2, 3, 4, 5}
	var slice1 = arr[2:4]
	fmt.Printf("局部变量： slice1 %v\n", slice1)

	//接口类型(interface)
	var peppa Animal = Pig{1, "佩奇"}
	peppa.Eat()
	peppa.Sleep()

	//Map类型

	var mapDemo = map[int]string{
		1: "张三",
		2: "李四",
	}
	fmt.Println(mapDemo)

}

type Animal interface {
	Eat()
	Sleep()
}
type Pig struct {
	id   int
	name string
}

func (pig Pig) Eat() {
	fmt.Println(pig.name + "吃...")
}

func (pig Pig) Sleep() {
	fmt.Println(pig.name + "睡...")
}
