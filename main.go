package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/a-barwick/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Welcome to Monkey, %s\n", user.Username)
	fmt.Println("Start typing")

	repl.Start(os.Stdin, os.Stdout)
}
