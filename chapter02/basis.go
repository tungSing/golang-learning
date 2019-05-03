/**
基础语法
- 行分隔符
- 注释
- 标识符
- 关键字
- 空格
 */
package basis

import "fmt"

func Base(){
	//行分隔符  一行代表一个语句结束，如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中并不鼓励这种做法。
	//注释分单行注释和多行注释
	//标识符 标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。
	var flag = true;
	fmt.Println(flag)
	/**
	关键字
	25个关键字或保留字
	break
	default
	func
	interface
	select
	case
	defer
	go
	map
	struct
	chan
	else
	goto
	package
	switch
	const
	fallthrough
	if
	range
	type
	continue
	for
	import
	return
	var
	36个预定义标识符
	append
	bool
	byte
	cap
	close
	complex
	complex64
	complex128
	uint16
	copy
	false
	float32
	float64
	imag
	int
	int8
	int16
	uint32
	int32
	int64
	iota
	len
	make
	new
	nil
	panic
	uint64
	print
	println
	real
	recover
	string
	true
	uint
	uint8
	uintptr
	 */

	//空格 变量的声明必须使用空格隔开
	var ss = 123
	ss++


}

