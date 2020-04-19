package main

import (
	"fmt"
	"github.com/nwehr/monkey/repl"
	"os"
)

func main() {
	fmt.Printf("Monkey REPL\n")
	repl.Start(os.Stdin, os.Stdout)
}
