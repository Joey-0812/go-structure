package main

import "fmt"

//func main() {
//	var p *int
//	i := 42
//	p = &i //p 指向i的地址
//
//	fmt.Println(*p) // 通过指针 p 读取 i
//	*p = 21         // 通过指针 p 设置 i
//}
func main() {

	//调用函数
	n := rpf()

	//显示值
	fmt.Println("n的值: ", *n)

}

//定义具有整数的函数
//指针作为返回类型
func rpf() *int {

	//局部变量
	//函数内部使用简短运算符声明
	lv := 100

	// 返回lv的地址
	return &lv
}

//只有一个特殊的地方，尽管b所表示的是Book1对象的内存地址，但是，在从b对应的内存地址取属性值的时候，
//就不是*b.title了。而是直接使用b.title，这点很特殊，它的效果就相当于Book1.title:
