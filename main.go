package main

import (
	"fmt"
	"go-noob/greet"
)

func main() {
    const question = "名前を入力してください: "

    var name string

    fmt.Print(question)
    fmt.Scanln(&name)
    
    greet.SayHello(name)
}
