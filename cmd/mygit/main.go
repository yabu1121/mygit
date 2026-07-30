package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: mygit <command>")
		return
	}

	switch os.Args[1] {
	case "init":
		if err := commands.Init; err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("unknown command: %s\n", os.Args[1])
	}
}
