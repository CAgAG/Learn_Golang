package Map

import (
	"fmt"
	"sync"
	"time"
)

type SyncMap struct {
	Map map[string] string
	*sync.RWMutex // 读写锁
}

var (
	Smap SyncMap // 公有map
	Done chan bool // 通道, 是否完成
	)

func WriteTest_1() {
	keys := []string{"a1", "a2", "a3"}
	for _, k := range keys {
		Smap.Lock()
		Smap.Map[k] = k
		Smap.Unlock()
		time.Sleep(1 * time.Second)
	}
	Done <- true // 通道写入
}

func Writetest_2() {
	keys := []string{"b1", "b2", "b3"}
	for _, k := range keys {
		Smap.Lock()
		Smap.Map[k] = k
		Smap.Unlock()
		time.Sleep(1 * time.Second)
	}
	Done <- true // 通道写入
}

func Read() {
	Smap.RLock()
	fmt.Println("Read Lock: ")
	for k, v := range Smap.Map {
		fmt.Println(k, v)
	}
	Smap.RUnlock()
}




















