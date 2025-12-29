package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 打开文件检测
func openfile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist: %s", filename)
		}
		return nil, err
	}
	//打开文件获取句柄 返回
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

/*
【单文件处理】
返回 ： 1.关键字所在的列 2.文件名字 3.error
*/
func File(filename string, keyword string) (int, string, error) {
	file, err := openfile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		//os.Exit(1)
	}
	defer file.Close()
	//读入buf中
	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			return i, filename, nil
		}
	}
	return -1, "NaN", scanner.Err()
}
