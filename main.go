package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)

	if len(os.Args) > 1 {
		file := os.Args[1]
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		repl.Start(f, os.Stdout)
		return
	}

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
