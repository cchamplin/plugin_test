package main

import "fmt"

type greeting string

//func main() {
//
//}

func (g greeting) Greet() {
    fmt.Println("Hello Universe")
}

// exported as symbol named "Greeter"
var Greeter greeting
