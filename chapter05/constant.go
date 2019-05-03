/**
常量
*/
package constant

import "fmt"

//显式类型定义
const MysqlPort int = 3306

//隐式类型定义
const MongoDBPort = 27017

//多常量初始化
const (
	Red                        = "#ff0000"
	Yellow                     = "#ffff00"
	OraclePart      int        = 1521
	Man             int        = 1
	Male                       //和Man同样的值
	ColorCodeLength = len(Red) //支持编译期可确定的函数返回值
)

//iota
/**
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
*/
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//iota 其他用法

const (
	_        = iota
	KB int64 = 1 << (10 * iota) //0100 0000 0000
	MB
	GB = 100 //独立值
	TB
	PB = iota //恢复计数
	EB
	ZB
	YB
	BB
)

func Demo() {
	fmt.Println(MysqlPort)
	fmt.Println(MongoDBPort)
	fmt.Println(Red)
	fmt.Println(Yellow)
	fmt.Println(OraclePart)
	fmt.Println(Man)
	fmt.Println(Male)
	fmt.Println(ColorCodeLength)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB, BB)
}
