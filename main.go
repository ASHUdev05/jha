package main

import (
	"fmt"
	"jha/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %s, This is the J H A programming language!\n", user.Username)
	fmt.Printf("Feel free to type your commands, below.\n")

	repl.Start(os.Stdin, os.Stdout)
}
