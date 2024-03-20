package main

import "fmt"

/*
func funcName(a interface{}) string {
        return string(a)
}
*/

func funcName(a interface{}) string {
	// 类型断言
	value, ok := a.(string)
	if !ok {
		fmt.Println("It is not ok for type string")
		return ""
	}
	fmt.Println("The value is ", value)

	return value
}

func main1() {
	var a interface{}
	a = "123"
	funcName(a)
	a = 10
	funcName(a)

	var t interface{}
	t = 10
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

	var x interface{}
	x = 100
	if x == nil {
		fmt.Println("NULL")
	} else if _, ok := x.(int); ok {
		fmt.Printf("%d", x)
	} else if _, ok := x.(uint); ok {
		fmt.Printf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			fmt.Println("TRUE")
		}
		fmt.Println("FALSE")
	} else if s, ok := x.(string); ok {
		fmt.Println(s) // (not shown)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
