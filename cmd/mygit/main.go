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

		write := false
		filePath := os.Args[2]

		if os.Args[2] == "-w" {
			if len(os.Args) < 4 {
				fmt.Println("usage: mygit hash-object -w <file>")
				os.Exit(1)
			}
			write = true
			filePath = os.Args[3]
		}

		if err := commands.HashObject(filePath, write); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "cat-file":
		if len(os.Args) < 4 {
			fmt.Println("usage: mygit cat-file -p <hash>")
			os.Exit(1)
		}
		option, err := commands.ParseCatFileOption(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		objectID := os.Args[3]
		if err := commands.CatFile(option, objectID); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s \n", os.Args[1])
		os.Exit(1)
	}
}
