package main

import (
	"fmt"
	"work_as/func2"
)

// 定义一个Point切片类型 Path
type Path []func2.Point

// 方法的接收器 是Path类型数据, 方法的选择器是TranslateBy(Point, bool)
func (path Path) TranslateBy(another func2.Point, add bool) {
	var op func(p, q func2.Point) func2.Point //定义一个 op变量 类型是方法表达式 能够接收Add,和 Sub方法
	if add == true {
		op = func2.Point.Add //给op变量赋值为Add方法
	} else {
		op = func2.Point.Sub //给op变量赋值为Sub方法
	}

	for i := range path {
		//调用 path[i].Add(another) 或者 path[i].Sub(another)
		path[i] = op(path[i], another)
		path[i].Print()
	}
}

func main3() {
	func2.Hello()

	points := Path{
		{10, 10},
		{11, 11},
	}

	anotherPoint := func2.Point{5, 5}

	points.TranslateBy(anotherPoint, false)

	fmt.Println("------------------")

	points.TranslateBy(anotherPoint, true)
}
