package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	n, err := pp.Println("Hello, world!")
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
}
