package main

import (
	"gogrep/internal/app"
)

//不依赖外部库 纯原生实现 文件搜索工具（mini grep）
//gogrep keyword ./logs
//核心功能
//遍历目录
//读取文件
//搜索关键词
//打印匹配行

func main() {
	app.Run()
}
