package main

import "fmt"

// 变量声明
var (
	var_a int
	var_b string
	var_c float64
	var_d byte
)

// 常量声明
const (
	const_a int     = 1
	const_b string  = "hello"
	const_c float64 = 3.14
	const_d byte    = 1
)

// 函数定义
func swap(x, y string) (string, string) {
	return y, x
}

// 函数定义: 指针参数
func swap2(x *int, y *int) {
	var tp int
	tp = *x
	*x = *y
	*y = tp
}

// 特殊函数: init. 每个包会优先执行的函数
// 注意: 如果一个包会被多个包同时导入，那么它只会被导入一次, 也就是只会执行一次init函数
func init() {
	fmt.Println("========================")
	fmt.Println("init func start")
	fmt.Println("========================")
}

// defer 语句: 延迟函数
/* defer作用：
 * 释放占用的资源
 * 捕捉处理异常
 * 输出日志 */
// 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。
func defer_demo() {
	fmt.Println("func start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("func end")
}

func Demo(i int) {
	// 定义10个元素的数组
	var arr [10]int
	// 错误拦截要在产生错误前设置
	defer func() {
		// 设置recover拦截错误信息
		err := recover()
		// 产生panic异常  打印错误信息
		if err != nil {
			fmt.Println("error log:")
			fmt.Println(err)
		}
	}()
	// 根据函数参数为数组元素赋值
	// 如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 10
}

func main() {
	fmt.Println("hello world")

	a := 10
	fmt.Printf("=%d\n", a)

	const b = 101
	fmt.Println("const:", 101)

	s_a := 1
	s_b := 2
	swap2(&s_a, &s_b)
	fmt.Printf("%d %d\n", s_a, s_b)

	defer_demo()
	Demo(11)

	// 声明切片(动态数组)。注意：Go 数组的长度不可改变
	var slice_a = []int{1, 2, 3, 4}
	//var slice_b = make([]int, 10, 20) // 10 是初始长度, 20 是最大长度。这个20这个参数都是可选的

	// 初始化切片copy_slice_a,是slice_a的引用
	var copy_slice_a = slice_a[:]
	copy_slice_a[0] = 100
	for i, v := range slice_a {
		fmt.Print("(", i, ",", v, ") ")
	}
	fmt.Println()

	// 初始化切片copy_slice_b,是slice_a的引用
	var copy_slice_b = slice_a[0:3]
	// 获取切片的长度和容量(slice_a)
	fmt.Println(len(copy_slice_b), cap(copy_slice_b))
	copy_slice_b[1] = 101
	for i, v := range slice_a {
		fmt.Print("(", i, ",", v, ") ")
	}
	fmt.Println()
	for i, v := range copy_slice_b {
		fmt.Print("(", i, ",", v, ") ")
	}
	fmt.Println()

	// 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	// 拷贝切片的 copy 方法和向切片追加新元素的 append 方法
	copy_slice_b = append(copy_slice_b, 105)
	for i, v := range copy_slice_b {
		fmt.Print("(", i, ",", v, ") ")
	}
	fmt.Println()
	double_len_b := make([]int, len(copy_slice_b), len(copy_slice_b)*2)
	// 将 copy_slice_b 的内容复制到 double_len_b
	copy(double_len_b, copy_slice_b)
	for i, v := range double_len_b {
		fmt.Print("(", i, ",", v, ") ")
	}
	fmt.Println()

	// map
	//第一种声明
	var test1 map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "golang"
	test1["three"] = "java"
	fmt.Println(test1) //map[two:golang three:java one:php]

	//第二种声明
	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "golang"
	test2["three"] = "java"
	fmt.Println(test2) //map[one:php two:golang three:java]

	//第三种声明
	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println(test3) //map[one:php two:golang three:java]

	language := make(map[string]map[string]string)
	language["php"] = make(map[string]string, 2)
	language["php"]["id"] = "1"
	language["php"]["desc"] = "php是世界上最美的语言"
	language["golang"] = make(map[string]string, 2)
	language["golang"]["id"] = "2"
	language["golang"]["desc"] = "golang抗并发非常good"

	fmt.Println(language) //map[php:map[id:1 desc:php是世界上最美的语言] golang:map[id:2 desc:golang抗并发非常good]]

	//增删改查
	val, key := language["php"] //查找是否有php这个子元素
	if key {
		fmt.Printf("%v \n", val)
	} else {
		fmt.Printf("no")
	}

	language["php"]["id"] = "3"         //修改了php子元素的id值
	language["php"]["nickname"] = "啪啪啪" //增加php元素里的nickname值
	delete(language, "php")             //删除了php子元素
	fmt.Println(language)
}
