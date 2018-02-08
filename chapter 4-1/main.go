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

const xx string = "Hello, World"
//cannot assign xx = "Other string"

//Multiple var
var(
	s=5
	t=10
	u=15
	)
fmt.Println(s,t,u)

v1,v2 := "fist", "sec"
fmt.Println(v1,v2)

//swap
v1, v2 = v2, v1
fmt.Println(v1,v2)
}