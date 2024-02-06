package main

import (
	"fmt"

	gomodulesayhello "github.com/iqbalmf/go-module-sayhello"
)

func main() {
	sayHello := gomodulesayhello.SayHello("Iqbal")
	fmt.Println(sayHello)
}
