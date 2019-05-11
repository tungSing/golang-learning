/*
压力测试
压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：
压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母*/
package unitTest

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < 100; i++ {
		Add(1, 2)
	}
}
