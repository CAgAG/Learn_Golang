package FileSearch

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func ReadFile(path string) {
	startTime := time.Now()
	file, _ := os.Open(path)
	defer file.Close()
	line_num := 0
	SearchData := ""
	fmt.Print("Search Data: ")
	fmt.Scanln(&SearchData)
	buffer := bufio.NewReader(file)
	for true {
		line, _, end := buffer.ReadLine()
		if end == io.EOF {
			break
		}
		data := string(line)
		if strings.Contains(data, SearchData){
			fmt.Println(data)
		}
		line_num++
	}
	fmt.Println("用时: ", time.Since(startTime))
	fmt.Println("行数: ", line_num)
}