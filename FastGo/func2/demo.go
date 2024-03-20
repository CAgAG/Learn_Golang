package func2

import "fmt"

func Hello() {
	fmt.Println("hello func2 demo")
}

type ST struct {
	name string
}

// 结构体的函数
func (st ST) func1() {
	st.name = "func1 change"
}

// func2 的引用传递会直接影响到 结构体的值
func (st *ST) func2() {
	st.name = "func2 change"
}

//func main() {
//	ss := ST{name: string("First name")}
//
//	fmt.Println(ss.name)
//	ss.func1()
//	fmt.Println(ss.name)
//	fmt.Println("================")
//
//	fmt.Println(ss.name)
//	ss.func2()
//	fmt.Println(ss.name)
//	fmt.Println("================")
//
//	fmt.Println(ss.name)
//	ss.func2()
//	fmt.Println(ss.name)
//	fmt.Println("================")
//
//}
