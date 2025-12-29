package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 检测是否为二进制文件
func checkbinaryfile(path string) (bool, error) {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	buf := make([]byte, 8000) //构建8kb的 缓冲区
	n, err := file.Read(buf)  //文件读入缓冲区中
	if err != nil || n == 0 {
		return false, err
	}
	for _, b := range buf[:n] {
		//通过比对是否是0 从而判断 因为文本文件一般很少是0
		if b == 0 {
			return true, nil
		}
	}
	return false, nil
}

// 判断字节是否在 ascii 有效字符内
func isAscii(b byte) bool {
	return b >= 0x20 && b <= 0x7E
}

func ascii_in_EXE(path string) ([]string, error) {
	if isExe, err := checkbinaryfile(path); !isExe || err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//设置 byte 读取 和字符读取
	var (
		current []byte
		result  []string
	)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if isAscii(b) {
			current = append(current, b)
		} else {
			if len(current) > 1 {
				//如果 长度大于1 那么就通过string格式写入 result
				result = append(result, string(current))
			}
			//否则清空
			current = current[:0]
		}
		if len(current) >= 1 {
			result = append(result, string(current))
		}
	}
	//处理exe文件
	return result, nil
}

func findEXE(path, keyword string) (bool, error) {
	string_in_exe, err := ascii_in_EXE(path)
	if err != nil {
		return false, err
	}
	for _, s := range string_in_exe {
		if strings.Contains(s, keyword) {
			return true, nil
		}
	}
	return false, nil
}

func EXE(path, keyword string) error {
	if ok, err := findEXE(path, keyword); err != nil {
		return err
	} else if ok != true {
		fmt.Printf("【ERROR】%s 文件中没有找到关键字 %s\n", path, keyword)
	} else {
		fmt.Printf("【SUCCESS】%s 文件中找到关键字 %s\n", path, keyword)
	}
	return nil
}
