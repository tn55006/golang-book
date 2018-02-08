package main

import "fmt"

func main() {

fmt.Print("Enter a number: ")
var input float64
fmt.Scanf("%f", &input)
var output float64
output = (input-32)*5/9
fmt.Printf("%.2F", output)
}