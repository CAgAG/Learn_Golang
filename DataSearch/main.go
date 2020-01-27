package main

import (
	"DataSearch/ThirdSearch"
	"fmt"
)

func main() {
	//FileSearch.ReadFile("./src/DataSearch/FileSearch/data.txt")

	//arr:=[]int {1,2,3,3,3,3,3,4,5,6,6,6,6,7,9,10}
	//fmt.Println(BinSearch.BinSearch_lesser(arr, 4))
	arr := []int{19, 3, 21, 89, 76, 22, 90, 56, 34}
	//fmt.Println(FindNLargest.FindNLargest(arr, 2))
	//fmt.Println(arr)
	fmt.Println(ThirdSearch.ThirdSearch(arr, 56))
}
