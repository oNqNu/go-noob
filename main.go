package main

import (
	"fmt"
	"go-noob/greet"
)

func main() {
    var name string

    fmt.Print("名前を入力してください: ")
    fmt.Scanln(&name)
    greet.SayHello(name)
}
