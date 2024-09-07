package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func InArray(list []string, data string) bool {
	for _, ele := range list {
		if reflect.DeepEqual(ele, data) {
			return true
		}
	}
	return false
}

func main() {
	cur_dirs, err := os.Open(".")
	if err != nil {
		return
	}

	defer cur_dirs.Close()

	files, err := cur_dirs.ReadDir(-1)
	if err != nil {
		return
	}
	file_map := make(map[string][]string)
	code_ext := []string{".go", ".cpp", ".hpp", ".py"}

	for _, file := range files {
		if file.IsDir() {
			if file.Name()[0] == '.' {
				continue
			}

			dir_files, _ := os.ReadDir("./" + file.Name())
			for _, dir_file := range dir_files {
				file_ext := filepath.Ext(dir_file.Name())
				if file_ext == "" {
					continue
				}

				if !InArray(code_ext, file_ext) {
					continue
				}

				file_map[file.Name()] = append(file_map[file.Name()], fmt.Sprintf("./%s", dir_file.Name()))
			}
		}
	}

	var cmd string
	var next_cmd string
	var cmd2files = make(map[string]string)
	map_keys := []string{}
	for mk := range file_map {
		map_keys = append(map_keys, mk)
	}
	slices.Sort(map_keys)

	for {
		fmt.Println("======================================================")
		for menu_i, key := range map_keys {
			fmt.Println(menu_i, key)
			cmd2files[strconv.Itoa(menu_i)] = key
		}

		fmt.Println(">>> 输入题目编号 >>>")
		_, err := fmt.Scanf("%s\n", &cmd)
		if err != nil {
			fmt.Println("输入读取出错！！！")
			fmt.Println(err)
			continue
		}

		select_dir, ok := cmd2files[cmd]
		if !ok {
			fmt.Println("命令不存在:", cmd)
			continue
		}
		// 随机 ==============================
		selected_code_files := []string{}
		for {
			values := file_map[select_dir]
			value_index := rand.Intn(len(values))
			value := values[value_index]
			selct_code_file := path.Join(select_dir, value)
			if len(selected_code_files) == len(values) {
				fmt.Println("已全部抽取过一次!!!")
				break
			}
			if InArray(selected_code_files, selct_code_file) {
				continue
			}
			selected_code_files = append(selected_code_files, selct_code_file)

			code_file, err := os.ReadFile(selct_code_file)
			if err != nil {
				fmt.Println("文件不存在:", code_file)
				continue
			}
			fmt.Println(selct_code_file + " <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
			print_str := string(code_file)
			print_str_lines := strings.Split(print_str, "\n")

			for i := 0; i < 6; i++ {
				if i >= len(print_str_lines) {
					break
				}
				fmt.Println(print_str_lines[i])
			}
			fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

			fmt.Println("输入0, 返回菜单, 否则继续抽题")
			fmt.Scanf("%s\n", &next_cmd)
			if next_cmd == "0" {
				break
			} else {
				continue
			}
		}

		fmt.Println("======================================================")
	}

}
