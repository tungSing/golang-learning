/**
接口
 一个 interface 类型定义了一个方法集作为其接口。
interface 类型的变量可以保存含有属于这个 interface 类型方法集超集的任何类型的值，这时我们就说这个类型实现了这个接口，未被初始化的 interface 类型变量的零值为 nil。

*/
package interfaceLearning

import "fmt"

func Demo() {
	// 接口定义
	var peppa = Pig{1, "佩奇"}
	peppa.Eat()
	peppa.Sleep()
	peppa.Wear()

	//多态
	var tom = Cat{"汤姆", 2}
	Eat(tom)
	Eat(peppa)

	//空接口
	Print(1)
	Print(tom)

}

func (pig Pig) Eat() {
	fmt.Println(pig.name + "吃...")
}

func (pig Pig) Sleep() {
	fmt.Println(pig.name + "睡...")
}

//实现多个接口。
func (pig Pig) Wear() {
	fmt.Println(pig.name + "穿...")
}

func (cat Cat) Eat() {
	fmt.Println(cat.id + "吃鱼...")
}

func (cat Cat) Sleep() {
	fmt.Println(cat.id + "睡...")
}

func Eat(animal Animal) {
	switch i := animal.(type) {
	case Cat:
		i.Eat()
	case Pig:
		i.Eat()
	default:
		fmt.Println("未知动物")
	}
}

//空接口 interface{} 没有任何方法签名，也就意味着任何类型都实现了空接口。其作用类似面向对象语言中的根对象 object。
func Print(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}
