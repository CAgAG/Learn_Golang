package Items

import (
	Queue2 "DataStructure/QueueArray"
	"io/ioutil"
)

func GetFileDir_Queue(path string) []string {
	files := []string{}

	Queue := Queue2.NewQueue()
	Queue.EnQueue(path)

	for true {
		path := Queue.DeQueue()
		if path == nil {
			break
		}
		read, _ := ioutil.ReadDir(path.(string))
		for _, f := range read{
			dir := path.(string)+"/"+f.Name()
			if f.IsDir() {
				Queue.EnQueue(dir)
			} else {
				files = append(files, dir)
			}
		}
	}

	return files
}
