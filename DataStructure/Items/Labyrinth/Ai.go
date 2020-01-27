package Labyrinth

import (
	"DataStructure/QueueArray"
	"DataStructure/StackArray"
	"fmt"
)

// 队列实现
func AiQueue(AiData [M][N]int,  i, j int) bool {
	q := QueueArray.NewQueue()
	q.EnQueue(&Pos{i, j})

	for true {
		nowPos := q.DeQueue().(*Pos)
		if nowPos == nil {
			break
		}

		i = nowPos.x
		j = nowPos.y
		// 到尽头
		if nowPos.x == M-1 && nowPos.y == N-1 {
			// 标记, 避免走回头路, 相当于障碍
			AiData[i][j] = 3
			CanGoOut = true
			// 更新Ai 地图
			AIData = AiData
			fmt.Println("Can go out")
			break
		} else {
			if CanGoOut {
				// 走不通设置为0 (空地)
				AiData[i][j] = 0
			} else {
				// 标记, 避免走回头路
				AiData[i][j] = 3
			}

			// 可能的有效步加入队列
			if j+1 <= N-1 && AiData[i][j+1] < 2 && CanGoOut != true {
				q.EnQueue(&Pos{i, j+1})
			}
			if i+1 <= M-1 && AiData[i+1][j] < 2 && CanGoOut != true {
				q.EnQueue(&Pos{i+1, j})
			}
			if j-1 >= 0 && AiData[i][j-1] < 2 && CanGoOut != true {
				q.EnQueue(&Pos{i, j-1})
			}
			if i-1 >= 0 && AiData[i-1][j] < 2 && CanGoOut != true {
				q.EnQueue(&Pos{i-1, j})
			}

		}
	}
	return CanGoOut
}

func AiStack(AiData [M][N]int,  i, j int) bool {
	stack := StackArray.NewStack()
	stack.Push(&Pos{i, j})

	for !stack.IsEmpty() {
		nowPos := stack.Pop().(*Pos)
		if nowPos == nil {
			break
		}

		i = nowPos.x
		j = nowPos.y
		// 标记, 避免走回头路
		AiData[i][j] = 3
		// 到尽头
		if nowPos.x == M-1 && nowPos.y == N-1 {
			CanGoOut = true
			// 更新Ai 地图
			AIData = AiData
			fmt.Println("Can go out")
			break
		} else {
			// 可能的有效步加入栈
			if j+1 <= N-1 && AiData[i][j+1] < 2 && CanGoOut != true {
				stack.Push(&Pos{i, j+1})
			}
			if i+1 <= M-1 && AiData[i+1][j] < 2 && CanGoOut != true {
				stack.Push(&Pos{i+1, j})
			}
			if j-1 >= 0 && AiData[i][j-1] < 2 && CanGoOut != true {
				stack.Push(&Pos{i, j-1})
			}
			if i-1 >= 0 && AiData[i-1][j] < 2 && CanGoOut != true {
				stack.Push(&Pos{i-1, j})
			}

		}
	}
	return CanGoOut
}

func AiRecursion(AiData [M][N]int, i, j int) bool {
	// 标记, 避免走回头路
	AiData[i][j] = 3

	if i == M-1 &&  j == N-1 {
		CanGoOut = true
		// 更新Ai 地图
		AIData = AiData
		fmt.Println("Can go out")
	} else {
		// 可能的有效步
		if j+1 <= N-1 && AiData[i][j+1] < 2 && CanGoOut != true {
			AiRecursion(AiData, i, j+1)
		}
		if i+1 <= M-1 && AiData[i+1][j] < 2 && CanGoOut != true {
			AiRecursion(AiData, i+1, j)
		}
		if j-1 >= 0 && AiData[i][j-1] < 2 && CanGoOut != true {
			AiRecursion(AiData, i, j-1)
		}
		if i-1 >= 0 && AiData[i-1][j] < 2 && CanGoOut != true {
			AiRecursion(AiData, i-1, j)
		}
	}
	return CanGoOut
}


