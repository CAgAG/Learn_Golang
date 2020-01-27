package Labyrinth

import "fmt"

func Show(arr [M][N]int) {
	fmt.Println("status: ", "-----------------------------------------")
	for i:=0;i<M;i++{
		for j:=0;j<N;j++{
			fmt.Printf("%4d",arr[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("status: ", "-----------------------------------------")
}
