package main

import (
	"fmt"
	"os"

	"github.com/yabu1121/mygit/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: mygit <command>")
		return
	}

	switch os.Args[1] {

	case "init":
		if err := commands.Init(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "hash-object":
		if len(os.Args) < 3 {
			fmt.Println("usage: mygit hash-object <file>")
			os.Exit(1)
		}
		if err := commands.HashObject(os.Args[2]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s \n", os.Args[1])
		os.Exit(1)
	}
}
