package flag

//展示命令行工具
import (
	"flag"
	"fmt"
	"os"
)

// 遍历目录
// 读取文件
// 搜索关键词
// 打印匹配行
var (
	FileString = flag.String("file", "", "默认查找当前文件下 [path,file,exe 需要有一个存在] ")
	PATH       = flag.String("path", "", "查找目录下的 关键字 [path,file,exe 需要有一个存在]")
	EXE        = flag.String("exe", "", "查找二进制文件下的ASCII [path,file,exe 需要有一个存在]")
	KeyWord    = flag.String("key", "", "关键字  (必须)")
)

type Config struct {
	FilePath string
	KeyWord  string
	EXE      string
	Path     string
}

// 解析参数 并且返回
func readflag() Config {
	return Config{
		FilePath: *FileString,
		KeyWord:  *KeyWord,
		EXE:      *EXE,
		Path:     *PATH,
	}
}

// 检查用户必要输入 并且返回解析后的内容
func Checkflag() Config {
	flag.Parse() //开启flag
	ok := *FileString == "" && *PATH == "" && *EXE == ""
	// 检测  如果没有必要项目 则报错
	if *KeyWord == "" || ok {
		fmt.Fprintln(os.Stderr, "missing required flags:")
		flag.Usage()
	}
	return readflag()
}
