# gogrep

一个基于 **Go 标准库**实现的命令行工具，用于在**文件 / 目录 / 可执行文件（EXE）**中搜索指定关键字。

支持：
- 普通文本文件
- 目录递归扫描
- Windows EXE 文件中的 ASCII 字符串匹配

---

## 安装

```bash
go build -o gogrep ./cmd/gogrep
```


# 使用方法
```
gogrep [options]
```
# 参数说明
```
参数	说明
-file	查找指定文件中的关键字
-path	在指定目录下递归查找关键字
-exe	在指定 EXE（二进制文件）的 ASCII 字符串中查找关键字
-key	要查找的关键字（必填）

⚠️ -file / -path / -exe 三者必须至少指定一个
```
# 示例

## 1️⃣ 在单个文件中查找关键字
```
gogrep -file test.txt -key password
```
## 2️⃣ 在目录中递归查找关键字
```
gogrep -path ./logs -key error
```

## 3️⃣ 在 EXE 文件的 ASCII 内容中查找关键字
```
gogrep -exe sample.exe -key http
```
# 说明
EXE 扫描规则

仅提取 可打印 ASCII 字符

默认忽略二进制数据


# 示例输出
```
gogrep.exe 文件中找到关键字 http
```