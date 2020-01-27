package Labyrinth

import "strings"

// wsad -> 上下左右
func Run(direct string) {
	// Tip: 2为障碍, 不能移动
	if strings.ToLower(direct) == "w" {
		if Ipos-1 >= 0 && Data[Ipos-1][Jpos] < 2 {
			// 交换
			Swap(Ipos, Jpos, Ipos-1, Jpos)
			Ipos -= 1
		}
	} else if strings.ToLower(direct) == "s" {
		if Ipos+1 <= M-1 && Data[Ipos+1][Jpos] < 2 {
			// 交换
			Swap(Ipos, Jpos, Ipos+1, Jpos)
			Ipos += 1
		}
	} else if strings.ToLower(direct) == "a" {
		if Jpos-1 >= 0 && Data[Ipos][Jpos-1] < 2 {
			// 交换
			Swap(Ipos, Jpos, Ipos, Jpos-1)
			Jpos -= 1
		}
	} else if strings.ToLower(direct) == "d" {
		if Jpos+1 >= N-1 && Data[Ipos][Jpos+1] < 2 {
			// 交换
			Swap(Ipos, Jpos, Ipos, Jpos+1)
			Jpos += 1
		}
	}

	// 移动后, 展示结果
	Show(Data)
}

func Swap(bi, bj int, ai, aj int) {
	Data[bi][bj], Data[ai][aj] = Data[ai][aj], Data[bi][bj]
}
