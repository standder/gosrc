package main

import "fmt"

func main(){
  var a int = 12
  var b int = 12
  var c bool = a == b
  var d bool = a != b
  var e bool = c && d
  var x bool = c || d
  var y bool = ! x
  fmt.Println(c,d)
  fmt.Println(e)
  fmt.Println(x,y)

}
