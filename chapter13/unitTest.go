/*
单元测试
Golang单元测试对文件名和方法名，参数都有很严格的要求。
1、文件名必须以xx_test.go命名
2、方法必须是Test[^a-z]开头
3、方法参数必须 t *testing.T
4、使用go test执行单元测试
*/
package unitTest

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
