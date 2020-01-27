package RPC

import (
	"errors"
	"fmt"
)

type Args struct {
	A, B int
}

type Query struct {
	X, Y int
}

type Last int

func (t *Last)Mul(args *Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println(reply, "乘法已执行")
	return nil
}

func (t *Last)Div(args *Args, query *Query) error {
	if args.B == 0 {
		return errors.New("不能除0")
	}
	query.X = args.A / args.B
	query.Y = args.A % args.B
	fmt.Println(query,"除法已执行")
	return nil
}
