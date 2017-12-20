package main

import (
	"fmt"
	contracts "github.com/ligurio/go-contracts"
)

func test(s string) {
	fmt.Println(s)
}

func main() {
	contracts.Contract(test, "Hello, world!", "Hello", false)
	contracts.Contract(test, "Bye, world!", "Bye", true)
}
