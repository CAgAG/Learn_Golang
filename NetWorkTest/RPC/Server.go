package RPC

import (
	"fmt"
	"net/rpc"
)

type ArgsX struct{
	A,B int  //两个数据
}

type Query struct{
	X,Y int  //两个数据
}

func Build() {
	client, err := rpc.DialHTTP("tcp","127.0.0.1:8848")
	if err == nil {
		fmt.Println(err)
	}
	data1 := 1
	data2 := 2

	args := ArgsX{data1, data2}
	var reply int
	err = client.Call("Last.Mul", args, &reply)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(args.A,args.B,reply)///乘法

	var qu Query
	err=client.Call("Last.Div",args,&qu)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(args.A,args.B,qu.X,qu.Y)
}

