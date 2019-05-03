/**
函数
特点：
• 无需声明原型。
• 支持不定 变参。
• 支持多返回值。
• 支持命名返回参数。
• 支持匿名函数和闭包。
• 函数也是一种类型，一个函数可以赋值给变量。

• 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
• 不支持 重载 (overload)
• 不支持 默认参数 (default parameter)。
*/
package function

import (
	"fmt"
	"math"
	"net/http"
)

/**
函数定义：
函数声明包含一个函数名，参数列表， 返回值列表和函数体。如果函数没有返回值，则返回列表可以省略。函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。
函数可以没有参数或接受多个参数。
注意类型在变量名之后 。
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
函数可以返回任意数量的返回值。
使用关键字 func 定义函数，左大括号依旧不能另起一行。
*/

func Demo() {

	//值传递
	f, g := 2, 3
	fmt.Println("传递前", f, g)
	ValueTransfer(f, g)
	fmt.Println("传递后", f, g)
	/**
	引用传递
	&h 指向 h 指针，h 变量的地址
	&i 指向 i 指针，i 变量的地址
	*/
	var h, i int = 100, 200
	fmt.Printf("交换前，h 的值 : %d\n", h)
	fmt.Printf("交换前，i 的值 : %d\n", i)
	Swap(&h, &i)
	fmt.Printf("交换后，h 的值 : %d\n", h)
	fmt.Printf("交换后，i 的值 : %d\n", i)

	//可变参数
	Varargs(0, 1, 2, 3, 4, 5)

	//不同类型可变参数
	DifferentType(0, "demo", 12, 5)

	//函数返回多个值
	a, b := ReturnValue("Mahesh", "Kumar")
	fmt.Println(a, b)

	//裸返回
	c := BareReturn()
	fmt.Println(c)

	//返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略
	d, _ := ReturnValue("张三", "李四")
	fmt.Println(d)

	//命名返回值
	j := NamedReturnParams()
	fmt.Println("命名返回值结果", j)

	//匿名函数
	AnonymousFunc()

	//闭包
	Closures()

	//递归函数
	result := Factorial(7)
	fmt.Println("7!=", result)

	//延迟调用
	Defer()

}

/**
函数可以通过两种方式来传递参数：
	值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数
*/
func ValueTransfer(x, y int) {
	x, y = 5, 6
}

func Swap(x *int, y *int) {
	var temp int
	//保持 x 地址上的值
	temp = *x
	//将 y 值赋给 x
	*x = *y
	//将 temp 值赋给 y
	*y = temp
}

//可变参数本质上就是 slice
func Varargs(args ...int) {
	fmt.Println("参数长度:", len(args))
	for _, v := range args {
		fmt.Println("参数值：", v)
	}
}

//用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的
func DifferentType(args ...interface{}) {
	for _, v := range args {
		fmt.Println("不同类型参数", v)
	}
}

func ReturnValue(x, y string) (string, string) {
	return y, x
}

func NamedReturnParams() (name string) {
	return "tungSing"
}

/**
没有参数的 return 语句返回各个返回变量的当前值
直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性
*/
func BareReturn() (c int) {
	c = 1 + 2
	return
}

/**
匿名函数是指不需要定义函数名的一种函数实现方式
函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。
匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
*/
func AnonymousFunc() {
	//function variable
	getSqrt := func(e float64) float64 {
		return math.Sqrt(e)
	}
	fmt.Println(getSqrt(4))

	//function collection
	fns := []func(x int) int{
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}

	fmt.Println(fns[0](100))

	//function as field
	d := struct {
		fn func() string
	}{
		fn: func() string {
			return "Hello World!"
		},
	}

	fmt.Println(d.fn())

	//channel of function
	fc := make(chan func() string, 2)
	fc <- func() string {
		return "Hello World!"
	}
	fmt.Println((<-fc)())
}

/**
Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
*/
func Closures() {
	nextInt := getSequence()
	fmt.Println("序号:", nextInt())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
递归函数
构成递归需具备的条件：
1. 子问题须与原始问题为同样的事，且更为简单。
2. 不能无限制地调用本身，须有个出口，化简为非递归状况处理。
*/
//n的阶乘指的是所有小于等于n的正整数的乘积
func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}

/**
延迟调用(defer)
特性：
1. 关键字 defer 用于注册延迟调用。
2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
3. 多个defer语句，按先进后出的方式执行。
4. defer语句中的变量，在defer声明时就决定了。
用途：
1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放
*/
func Defer() {

	//defer named return 匿名函数
	a := DeferNamedReturnClosures()
	fmt.Println("命名返回值函数结果", a)

	// defer return 匿名函数
	b := DeferReturnClosures()
	fmt.Println("普通返回值函数结果", b)

	//不使用匿名函数情况
	DeferGeneralFunc()

	//读取资源
	c := ReadResource()
	fmt.Println(c)

}

//当函数有命名结果形参时，defer函数是可以修改它，然后再将它的值返回
//当函数有命名结果形参时，结果形参的初始值被设置为零值，函数的return语句会设置结果形参的值
func DeferNamedReturnClosures() (i int) {
	defer func() {
		i = 1
		fmt.Println("执行a函数", i)
	}()
	return i
}

func DeferReturnClosures() int {
	i := 0
	defer func() {
		i = 1
		fmt.Println("执行b函数", i)
	}()
	return i
}

func DeferGeneralFunc() {
	defer b(1)
	fmt.Println("执行DeferDemo函数")
}

func b(i int) {
	fmt.Println("执行b函数", i)
}

func ReadResource() error {
	fmt.Println("访问草原天路首页--->>>")
	res, err := http.Get("http://www.caoytl.com/")

	//defer res.Body.Close() 错误写法
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	fmt.Println(res.Status)
	return nil

}
