// 泛型定义
package main

import "fmt"

// 泛型接口
// 表明 该类型可以代表以下类型
type TInt interface {
	int | ~int32 | ~int64 | float32 | float64 // ~T表示底层类型是T的所有类型, 例如 rune的底层就是 int32 ==> type rune = int32
}

// 泛型函数: 加 []
// ==> 注意输入的 a, b类型必须相同
// 例如 a是int, b也要是int。同理 a是float32, b也要是float32
func compare[T TInt](a, b T) bool {
	return a > b
}

// 不使用接口, 直接定义也可以
func compare2[T int | float32 | float64 | ~int32](a, b T) bool {
	return a > b
}

// 泛型结构体
type compare3[T int | float32 | float64 | int32 | int64, T2 string] struct {
	num1 T
	num2 T
	name T2
}

func main() {
	fmt.Println(compare(1, 2))
	fmt.Println(compare(1.1, 2.1))

	fmt.Println(compare2(1, 2))
	fmt.Println(compare2(1.1, 2.1))

	com := compare3[int, string]{num1: 1, num2: 2, name: "cag1"}
	fmt.Println(com.num1)
	fmt.Println(com.num2)
	fmt.Println(com.name)

	com2 := compare3[float32, string]{num1: 1.1, num2: 2.2, name: "cag2"}
	fmt.Println(com2.num1)
	fmt.Println(com2.num2)
	fmt.Println(com2.name)

	com3 := new(compare3[float32, string])
	com3.num1 = 1.1
	com3.num2 = 2.1
	com3.name = "cag3"
	fmt.Println(com3.num1)
	fmt.Println(com3.num2)
	fmt.Println(com3.name)
}
