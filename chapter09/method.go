/**
方法：
Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针

方法定义：
func (recevier type) methodName(参数列表)(返回值列表){}
	参数和返回值可以省略

例子：
package main

type Test struct{}

// 无参数、无返回值
func (t Test) method0() {

}

// 单参数、无返回值
func (t Test) method1(i int) {

}

// 多参数、无返回值
func (t Test) method2(x, y int) {

}

// 无参数、单返回值
func (t Test) method3() (i int) {
	return
}

// 多参数、多返回值
func (t Test) method4(x, y int) (z int, err error) {
	return
}

// 无参数、无返回值
func (t *Test) method5() {

}

// 单参数、无返回值
func (t *Test) method6(i int) {

}

// 多参数、无返回值
func (t *Test) method7(x, y int) {

}

// 无参数、单返回值
func (t *Test) method8() (i int) {
	return
}

// 多参数、多返回值
func (t *Test) method9(x, y int) (z int, err error) {
	return
}

func main() {}

• 只能为当前包内命名类型定义方法。
• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
• 不支持方法重载，receiver 只是参数签名的组成部分。
• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。

*/
package method

import "fmt"

func Demo() {
	// 值类型调用方法
	u1 := User{1, "zhangsan", "man"}
	fmt.Println("值类型方法调用前:", u1)
	u1.ValueNotify()
	fmt.Println("值类型方法调用后:", u1)

	// 指针类型调用方法
	u2 := User{2, "lisi", "man"}
	u3 := &u2
	fmt.Println("指针方法调用前:", u2)
	u3.PointerNotify()
	fmt.Println("指针方法调用后:", u2)

	//匿名字段-变相实现继承
	m := Manager{User{1, "zhangsan", "man"}, 1}
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())

	/**
	方法集
		• 类型 T 方法集包含全部 receiver T 方法。
		• 类型 *T 方法集包含全部 receiver T + *T 方法。
	*/

	u4 := User{3, "xiaoming", "man"}
	u5 := &u4
	fmt.Printf("u5 is : %v\n", u5)
	u5.ValueNotify()
	u5.PointerNotify()


	//方法表达式

	u6 :=User{4,"xiaohong","famale"}

	uValue :=u6.ValueNotify
	fmt.Println(uValue)
	uValue()

	uExpression := (*User).PointerNotify
	fmt.Println(uExpression)
	uExpression(&u6)

}

//用户结构体
type User struct {
	id   int
	name string
	sex  string
}

//方法
func (u User) ValueNotify() {
	u.name = "wangwu"
}

func (u *User) PointerNotify() {
	u.name = "zhaoliu"
}

type Manager struct {
	User
	managerId int
}

func (u *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

func (m *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", m, m)
}
