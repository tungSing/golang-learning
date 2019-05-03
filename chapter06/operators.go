/**
运算符
内置有
	算术运算符
	关系运算符
	逻辑运算符
	位运算符
	赋值运算符
	其他运算符
*/
package operators

import "fmt"

func Demo() {

	/**
	算术运算符
	*/
	fmt.Println("---Arithmetic Operators---")
	Arithmetic()

	/**
	关系运算符
	*/
	fmt.Println("---Relational Operators---")
	Relational()

	/**
	逻辑运算符
	*/
	fmt.Println("---Logical Operators---")
	Logical()

	/**
	位运算符
	*/
	fmt.Println("---Bit Operators---")
	Bit()

	/**
	赋值运算符
	*/
	fmt.Println("---Assignment Operators---")
	Assignment()

	/**
	其他运算符
	 */
	fmt.Println("---Other Operators---")
	Other()
}

func Arithmetic() {
	var a, b = 1, 2
	//相加
	c := a + b
	fmt.Println(c)

	//相减
	d := a - b
	fmt.Println(d)

	//相乘
	e := a * b
	fmt.Println(e)

	//相除
	f := a / b
	fmt.Println(f)

	//求余
	g := a / b
	fmt.Println(g)

	//自增
	a++
	fmt.Println(a)

	//自减
	b--
	fmt.Println(b)
}

func Relational() {
	var a, b = 21, 10

	//检查两个值是否相等，如果相等返回 True 否则返回 False。
	fmt.Println(a == b)

	//检查两个值是否不相等，如果不相等返回 True 否则返回 False。
	fmt.Println(a != b)

	//检查左边值是否大于右边值，如果是返回 True 否则返回 False。
	fmt.Println(a > b)

	//检查左边值是否小于右边值，如果是返回 True 否则返回 False。
	fmt.Println(a < b)

	//检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。
	fmt.Println(a >= b)

	//检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。
	fmt.Println(a <= b)

}

func Logical() {
	var a, b = true, false

	//逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False
	fmt.Println(a && b)

	//逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。
	fmt.Println(a || b)

	//逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。
	fmt.Println(!a)
}

func Bit() {
	//60 = 0011 1100
	//13 = 0000 1101
	var a, b = 60, 13

	//按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。
	// 12 = 0000 1100
	fmt.Println(a & b)

	//按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或
	//61 = 0011 1101
	fmt.Println(a | b)

	//按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。
	//49 = 0011 0001
	fmt.Println(a ^ b)

	//左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
	//0011 1100 00
	//1111 0000
	fmt.Println(a << 2)

	//右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。
	// 0011 1100
	// 0000 1111
	fmt.Println(a >> 2)
	//0000 1101
	//0000 0011
	fmt.Println(b >> 2)
}

func Assignment() {
	var a = 21
	var b int

	//简单的赋值运算符，将一个表达式的值赋给一个左值
	b = a
	fmt.Println(b)

	//相加后再赋值
	b += a
	fmt.Println(b)

	//相减后再赋值
	b -= a
	fmt.Println(b)

	//相乘后再赋值
	b *= a
	fmt.Println(b)

	//相除后再赋值
	b /= a
	fmt.Println(b)

	//求余后再赋值
	b %= a
	fmt.Println(b)

	//1100 1000
	b = 200

	//左移后赋值 b = b << 2
	//0011 0010 0000
	b <<= 2
	fmt.Println(b)

	//右移后赋值
	//1100 1000
	b >>= 2
	fmt.Println(b)

	//按位与后赋值 b = b & 2
	//200 = 1100 1000
	//2   = 0000 ‭0010‬
	//      0000 0000
	b &= 2
	fmt.Println(b)

	//按位异或后赋值
	//0000
	//‭0010‬
	b ^= 2
	fmt.Println(b)

	//按位或后赋值
	//‭0010‬
	//‭0010‬
	b |= 2
	fmt.Println(b)

}

func Other() {
	var a = 4
	var ptr *int

	//返回变量存储地址
	fmt.Println(&a)

	//指针变量
	ptr = &a
	fmt.Println(*ptr)
}

