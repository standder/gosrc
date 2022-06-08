# gosrc
此仓库用于存放学习go语言所产生的源文件与可执行二进制文件



## 项目参照对象
该项目文件参照go guide项目进行学习

## 项目目标
通过go语言能够具备全栈开发能力

## 项目目录（持续更新中）
>go语言的基础应用

## 数据类型相关问题
int数据与float64数据做运算时会出现报错
```
package main

import "fmt"

func main(){
  var a int = 2
  var b float64 = 3.5
  fmt.Println(a+b)
}
```
利用go run运行后程序会返回要求选定结果输出数据类型
```
.\ex2.go:10:23: invalid operation: a + b + c (mismatched types int and float64)
```
