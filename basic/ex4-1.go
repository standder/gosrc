package main

import(
  "fmt"
  "unsafe"
  )

func main(){
  var a int = 1
  var b int8 = 1
  var c int16 = 1
  var d int32 = 1
  var e int64 = 1
  var g bool = true
  fmt.Println(unsafe.Sizeof(a))
  fmt.Println(unsafe.Sizeof(b))
  fmt.Println(unsafe.Sizeof(c))
  fmt.Println(unsafe.Sizeof(d))
  fmt.Println(unsafe.Sizeof(e))
  fmt.Println(unsafe.Sizeof(g))
}
