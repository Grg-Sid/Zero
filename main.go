package main

import (
	"fmt"
	"os"
	"os/user"
	"zero/repl"
)

func main() {
	_, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Print("Hello Wolrd!\n")
	fmt.Print("You can write ZERO commands here\n")
	repl.Start(os.Stdin, os.Stdout)
}
