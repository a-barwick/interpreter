package main

import (
	"fmt"

	"github.com/a-barwick/interpreter/lexer"
	"github.com/a-barwick/interpreter/token"
)

func main() {
	hello := lexer.Parse()
	world := token.Do()

	fmt.Printf("%s, %s!\n", hello, world)
}
