package main

import "fmt"

func main() {
//long Declaration
var x string = "Hello, World"
fmt.Println(x)

var y string
y = "Hello, World"
fmt.Println(y)

//Short Declaration
//Type Inference
z:= "Hello,World"
a:= 1234
b:= 10.235
//c:= abc123
fmt.Println(z)
fmt.Printf("Type z: %T\n", z)
fmt.Printf("Type a: %T\n", a)
fmt.Printf("Type b: %T\n", b)
//fmt.Printf("Type c: %T\n", c)
}