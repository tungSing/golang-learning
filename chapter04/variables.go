/**
变量
	函数内定义的变量称为局部变量
	函数外定义的变量称为全局变量
	函数定义中的变量称为形式参数
*/
package variables

import "fmt"

//全局变量
var cc, python, java = true, false, "no!"
var (
	c1 int
	d1 string
)

func Demo() {
	var a = "initial"
	var b int
	fmt.Println(a)
	fmt.Println(b)

	//一次性定义多个变量
	var c, d int = 3, 4
	fmt.Println(c, d)
	fmt.Println(c1, d1)

	var e = true
	fmt.Println(e)

	//短声明变量-在函数内部，可以使用更简略的 ":=" 方式定义变量。
	f := "short"
	//注意检查，是定义新的局部变量，还是修改全局变量。该方式容易造成错误！
	fmt.Println(&cc, &python, &java)
	cc, python, java := false, true, "yes!"
	fmt.Println(&cc, &python, &java)
	fmt.Println(cc, python, java, f)

	//注意重新赋值与定义新同名变量的区别
	Difference()

	//形式参数
	sum(c, d)

}

func Difference() {
	s := "abc"
	fmt.Println(&s)
	//重新赋值: 与前 s 在同 层次的代码块中，且有新的变量被定义。
	s, y := "hello", 20
	fmt.Println(&s, y)
	{
		// 定义新同名变量: 不在同 层次代码块。
		s, z := 1000, 30
		fmt.Println(&s, z)
	}

}

func sum(k, f int) int {
	fmt.Printf("sum() 函数中 k = %d\n", k)
	fmt.Printf("sum() 函数中 f = %d\n", f)
	return k + f
}
