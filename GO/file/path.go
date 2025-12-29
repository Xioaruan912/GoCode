package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// 检测目录
func checkPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("%s", "not a path")
	}
	if isdir := info.IsDir(); !isdir {
		return fmt.Errorf("%s", "not a path")
	}
	return nil
}

// 排除exe 文件 传入文件fullpath
func exeout(fullpath os.DirEntry) bool {
	ext := filepath.Ext(fullpath.Name())
	isEXE := ext == ".exe"
	return isEXE
}

/*
【目录处理】
返回 1.错误
*/
func PATH(dir, key string) error {
	//检测目录
	if err := checkPath(dir); err != nil {
		return err
	}
	//获取目录入口
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	//对所有文件操作
	for _, en := range entries {

		//构建完整路径
		filename := en.Name()
		fullpath := filepath.Join(dir, filename)
		//检测 是否为exe
		isExe := exeout(en)
		//如果不是exe则进入处理
		if !isExe {
			//如果不是 目录则进入处理
			if !en.IsDir() {
				//通过调用单文件处理方式 查找关键字 返回 关键字列 文件名字 查找错误
				if _, file, _ := File(fullpath, key); file != "NaN" {
					fmt.Printf("KEYWORD in %s\n", file)
				}
			}
		}
	}
	return nil
}
