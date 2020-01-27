package main

import (
	"DataStructure/Items/Labyrinth"
	"fmt"
)

func main() {
	//list := ArrayList.NewArrayList()
	//list.Append(1)
	//list.Append(2)
	//list.Append(3)
	//
	//list.Insert(3, 4)
	////list.Insert(3, 4)
	////list.Delete(3)
	//fmt.Println(list)
	//list.ToDataStore()
	//fmt.Println("hello world")
	//for it:= list.Iterator(); it.HasNext();{
	//	item,  _ := it.Next()
	//	fmt.Println(item)
	//}
	//myStack := StackArray.NewStack()
	//myStack.Push(1)
	//myStack.Push(2)
	//fmt.Println(myStack.Pop())
	//fmt.Println(myStack.Size())

	//files, _ := GetFileDir(".", []string{})
	//files := GetFileDir_Stack(".")
	//files := GetFileDir_Queue(".")
	//for  i:=0;i<len(files);i++{ //打印
	//	fmt.Println(files[i])
	//}

	//Convert10toN(10, 2)
	//fmt.Println("9" >= "0")
	//fmt.Println(GetResult(InToPost("(1+25)*5")))

	//set := SetArray.NewSet()
	//set.Add(1)
	//set.Add(2)
	//set.Add(1)
	//fmt.Println(set.Strings())

	//SinglnLink := SingleLink.NewSingleLinkList()
	//SinglnLink.InsertNodeFront(SingleLink.NewSingleLinkNode(1))
	//SinglnLink.InsertNodeFront(SingleLink.NewSingleLinkNode(2))
	//SinglnLink.InsertNodeTail(SingleLink.NewSingleLinkNode(3))
	//fmt.Println(SinglnLink.String())
	//SinglnLink.InsertNodeTail(SingleLink.NewSingleLinkNode(4))
	//SinglnLink.InsertNodeValueBack(3, SingleLink.NewSingleLinkNode(33))
	//SinglnLink.InsertNodeValueFront(3, SingleLink.NewSingleLinkNode(333))
	//fmt.Println(SinglnLink.String())
	//SinglnLink.DeleteIndex(1)
	//fmt.Println(SinglnLink.GetMid().Value())
	//fmt.Println(SinglnLink.String())
	//SinglnLink.ReverseList()
	//fmt.Println(SinglnLink.String())

	//DoubleLink := DoubleLink2.NewDoubleLinkList()
	//DoubleLink.InsertHead(DoubleLink2.NewDoubleLinkNode(1))
	//DoubleLink.InsertHead(DoubleLink2.NewDoubleLinkNode(2))
	//DoubleLink.InsertBack(DoubleLink2.NewDoubleLinkNode(3))
	//fmt.Println(DoubleLink.String())
	//DoubleLink.InsertValueBackByValue(3, DoubleLink2.NewDoubleLinkNode(33))
	//DoubleLink.InsertValueHeadByValue(3, DoubleLink2.NewDoubleLinkNode(333))
	//fmt.Println(DoubleLink.String())
	//DoubleLink.DeleteNodeAtIndex(1)
	//fmt.Println(DoubleLink.String())
	//fmt.Println(DoubleLink.GetNodeAtIndex(2).Value())

	//MTable, _ := HashTableArray.NewHashTable(100, HashTableArray.SHA)
	//MTable.Insert("a")
	//MTable.Insert("b")
	//MTable.Insert("c")
	//pos := MTable.Find("a")
	//fmt.Println(MTable.GetValue(pos))
	//pos = MTable.Find("c")
	//fmt.Println(MTable.GetValue(pos))
	//MTable.Empty()
	//fmt.Println(MTable.GetValue(pos))

	//h := HeapArray.NewHeap()
	//h.Insert(HeapArray.DataType(1))
	//h.Insert(HeapArray.DataType(5))
	//h.Insert(HeapArray.DataType(7))
	//h.Insert(HeapArray.DataType(2))
	//h.Insert(HeapArray.DataType(4))
	//fmt.Println(h.Extract())
	//fmt.Println(h.Extract())
	//fmt.Println(h.Extract())

	//h := QueueArray.NewMinPriorityQueue()
	//h.Insert(*QueueArray.NewPriorityitem(101, 10))
	//h.Insert(*QueueArray.NewPriorityitem(102, 11))
	//h.Insert(*QueueArray.NewPriorityitem(103, 12))
	//h.Insert(*QueueArray.NewPriorityitem(104, 7))
	//fmt.Println(h.Extract())
	//fmt.Println(h.Extract())
	//fmt.Println(h.Extract())

/*
				 1
			2		 3
		 4	  5   6		7
				          8
*/
	//bst := BinTree.NewTree()
	//
	//node1 := &BinTree.Node{1, nil, nil}
	//node2 := &BinTree.Node{2, nil, nil}
	//node3 := &BinTree.Node{3, nil, nil}
	//node4 := &BinTree.Node{4, nil, nil}
	//node5 := &BinTree.Node{5, nil, nil}
	//node6 := &BinTree.Node{6, nil, nil}
	//node7 := &BinTree.Node{7, nil, nil}
	//node8 := &BinTree.Node{8, nil, nil}
	//
	//bst.Root=node1
	//
	//node1.Left=node2
	//node1.Right=node3
	//
	//node2.Left=node4
	//node2.Right=node5
	//
	//node3.Left=node6
	//node3.Right=node7
	//
	//node7.Right=node8
	//
	//bst.Size=7
	//nodelast:=bst.FindAncestor(bst.Root,node6,node7)
	//fmt.Println(nodelast.Data)
	//fmt.Println("Depth: ", bst.GetDepth(bst.Root))
	//fmt.Println("PreOrder: ")
	//bst.PreOrder()
	//fmt.Println()
	//fmt.Println(bst.PreOrderStack())
	//bst.LevelShow()
	//fmt.Println("\nmax: ", bst.FindMax())
	//bst.RemoveMax()
	//fmt.Println("max: ", bst.FindMax())

	//H := HeapLink.NewLeftHeap(3)
	//H = HeapLink.Insert(2, H)
	//H = HeapLink.Insert(4, H)
	//H = HeapLink.Insert(9, H)
	//H = HeapLink.Insert(6, H)
	//H = HeapLink.Insert(7, H)
	//
	//HeapLink.PrintHeap(H)
	//H, data := HeapLink.DeleteMax(H)
	//fmt.Println("\n", data)
	//HeapLink.PrintHeap(H)

	//for i := 0; i < 10; i++ {
	//	n := &CircleLink.Node{i, nil}
	//	CircleLink.Add(n)
	//}
	//CircleLink.Show(CircleLink.Head)
	//CircleLink.Jose(3, 3)
	//CircleLink.Show(CircleLink.Head)

	//head := CircleLink.NewDCircleLinkNode("A")
	//head.Show()
	//node1 := &CircleLink.DCircleLink{3,"b",nil,nil}
	//node2 := &CircleLink.DCircleLink{2,"c",nil,nil}
	//node3 := &CircleLink.DCircleLink{5,"d",nil,nil}
	//node4 := &CircleLink.DCircleLink{4,"e",nil,nil}
	//
	//head.AddNode(node1)
	//head.AddNode(node2)
	//head.AddNode(node3)
	//head.AddNode(node4)
	//
	//head.Show()
	//
	//head.DeleteNodeById(5)
	//head.Show()
	//
	//head.ChangeNodeById(3, "x")
	//head.Show()

	//deQ := DeQueArray.NewDeque(4)
	//deQ.Addleft(1)
	//deQ.Addleft(2)
	//deQ.Addleft(3)
	//deQ.Show()
	//deQ.Delleft()
	//deQ.Show()
	//deQ.Delright()
	//deQ.Show()

	//test := []string{"123","456","789"}
	//hashMap := &HashMap.Ring{map[uint32]string{}, HashMap.Rindex{}, new(sync.RWMutex)}
	//
	//for _, v := range test {
	//	index:=crc32.ChecksumIEEE([]byte(v)) // 循环索引
	//	hashMap.Rmap[index] = v
	//	hashMap.RindexArr = append(hashMap.RindexArr, index)
	//}
	//fmt.Println(hashMap)
	////处理索引数组
	//sort.Sort(hashMap.RindexArr)
	//fmt.Println(hashMap)
	//
	//hashMap.AddNode("101112")
	//fmt.Println(hashMap)
	//hashMap.RemoveNode("101112")
	//fmt.Println(hashMap)
	//node := hashMap.GetNode("123")
	//fmt.Println(node)

	//ht := HashTableLink.NewHashTableLink(1000)
	//ht.Put("1","11")
	//ht.Put("2","22")
	//ht.Put("3","33")
	//ht.Put("4","44")
	//ht.Put("5","55")
	//
	//ht.Put("3", "66")
	////ht.Del("3")
	//val, err := ht.Get("3")
	//fmt.Println(val, err)

	//matcher := SetLink.DefaultMatch
	//set1 := new(SetLink.Set)
	//set1.Init(matcher)
	//set2 := new(SetLink.Set)
	//set2.Init(matcher)
	//
	//set1.Insert(1)
	//set1.Insert(2)
	//set1.Insert(3)
	//
	//set2.Insert(2)
	//set2.Insert(3)
	//set2.Insert(4)
	//
	//union := set1.Union(set2)
	//for it := union.GetIterator(); it.HashNext(); {
	//	fmt.Print(it.Next(), " ")
	//}
	//
	//fmt.Println()
	//
	//share := set1.Share(set2)
	//for it := share.GetIterator(); it.HashNext(); {
	//	fmt.Print(it.Next(), " ")
	//}
	//
	//fmt.Println()
	//
	//dif := set1.Different(set2)
	//for it := dif.GetIterator(); it.HashNext(); {
	//	fmt.Print(it.Next(), " ")
	//}
	//
	//fmt.Println()
	//
	//fmt.Println(set1.IsEquals(set2))
	//fmt.Println(set1.IsSub(set2))

	Ai := Labyrinth.AiQueue
	ok := Ai(Labyrinth.AIData, 0,0)
	if ok {
		Labyrinth.Show(Labyrinth.AIData)
		//Labyrinth.AiMove()
		//Labyrinth.Show(Labyrinth.AIData)
	} else {
		fmt.Println("走不出")
	}
}
