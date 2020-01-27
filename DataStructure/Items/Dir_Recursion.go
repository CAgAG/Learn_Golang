package Items

import (
	"errors"
	"io/ioutil"
)

func GetFileDir(path string, files [] string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, f := range read {
		dir := path+"/"+f.Name()
		if f.IsDir() {
			// 文件夹处理
			files = append(files, dir)
			files, _ = GetFileDir(dir, files) // 递归处理
		} else {
			// 文件
			files = append(files, dir)
		}
	}
	return files, nil
}
