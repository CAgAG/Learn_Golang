package Items

import (
	"DataStructure/StackArray"
	"io/ioutil"
)

func GetFileDir_Stack(path string) []string {
	files := []string{}

	Stack := StackArray.NewStack()
	Stack.Push(path)

	for !Stack.IsEmpty() {
		Path := Stack.Pop().(string) // 类型断言, 因为Stack是interface{}类型
		read, _ := ioutil.ReadDir(Path)
		for _, f := range read{
			dir := Path +"/"+f.Name()
			if f.IsDir() {
				// 文件夹
				Stack.Push(dir)
			} else {
				// 文件
				files = append(files, dir)
			}
		}
	}

	return files
}
