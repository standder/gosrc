package main

import "fmt"

func main(){
  var a int = 10
  var b int = 12
  var c bool = a == b
  var d bool = a != b
  var e bool =  !(c && d)
  var g bool = c || d
  fmt.Println(a+b)
  fmt.Println(c,d)
  fmt.Println(!e,!g)
}
