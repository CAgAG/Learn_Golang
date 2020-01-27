package main

import (
	"DistributedSystem/RPC"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	//arr := []int{1,9,2,8,3,7,4,6,5,10}
	//Sort.ShellSortGoRoutine(arr)
	//fmt.Println(arr)

	//Map.Smap = Map.SyncMap{make(map[string] string), new(sync.RWMutex)}
	//Map.Done = make(chan bool, 2)
	//
	//smap := Map.Smap
	//done := Map.Done
	//
	//go Map.WriteTest_1()
	//go Map.Writetest_2()
	//
	//for true {
	//	Map.Read()
	//	if len(done) == 2 {
	//		fmt.Println(smap.Map)
	//		for k, v := range smap.Map {
	//			fmt.Println(k, v)
	//		}
	//		break
	//	} else {
	//		time.Sleep(1 * time.Second)
	//	}
	//}

	//go func() {
	//	time.Sleep(1000 * time.Second)
	//}()
	//path := "./src/DistributedSystem/PipLineMiddleWare/"
	//p := PipLineMiddleWare.CreateNetWordPipeLine( path+"data.in", 100, 2)
	//PipLineMiddleWare.Write2File(p, path+"data.out")
	//PipLineMiddleWare.UseTime()
	//PipLineMiddleWare.ShowFile(path+"data.out")

	//server_listener, err := net.Listen("tcp", "127.0.0.1:8848")
	//if err != nil {
	//	panic(err)
	//}
	//defer server_listener.Close()
	//for true {
	//	new_conn, err := server_listener.Accept()
	//	if err != nil {
	//		break
	//		panic(err)
	//	}
	//	go NetWork.MsgHandler(new_conn)
	//}

	last := new(RPC.Last)
	fmt.Println("Last: ", last)
	rpc.Register(last) // 注册类型
	rpc.HandleHTTP() // 设定http类型
	listener, err := net.Listen("tcp", "127.0.0.1:8848")
	if err == nil {
		fmt.Println(err)
	}
	http.Serve(listener, nil)

}
