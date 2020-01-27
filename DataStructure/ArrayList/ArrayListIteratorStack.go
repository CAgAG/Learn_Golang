package ArrayList

type StackArrayX interface {
	Clear() //清空
	Size()int  //大小
	Pop()interface{}  //弹出
	Push(data interface{})//压入
	IsFull() bool //是否满了
	IsEmpty() bool//是否为空
}
type StackX struct{
	myarray *ArrayList
	iter Iterator

}
func  NewArrayListStackX()*StackX{
	mystack:=new(StackX)
	mystack.myarray=NewArrayList()           //数组
	mystack.iter =mystack.myarray.Iterator() //迭代
	return mystack


}
func (mystack * StackX )Clear() {
	mystack.myarray.Clear()
	mystack.myarray.theSize=0
}
func (mystack * StackX )Size()int  {
	return  mystack.myarray.theSize
}
func (mystack * StackX )Pop()interface{} {
	if  !mystack.IsEmpty(){
		last:=mystack.myarray.dataStore[mystack.myarray.theSize-1]
		mystack.myarray.Delete(mystack.myarray.theSize-1)
		return last

	}
	return nil
}
func (mystack * StackX ) Push(data interface{}){
	if !mystack.IsFull(){
		mystack.myarray.Append(data)
	}
}
func (mystack * StackX ) IsFull() bool { //判断满了
	if mystack.myarray.theSize>=10{
		return true
	}else{
		return false
	}
}
func (mystack * StackX )IsEmpty() bool{//判断为空
	if  mystack.myarray.theSize==0{
		return true
	}else{
		return false
	}
}
