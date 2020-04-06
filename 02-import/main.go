package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	fmt.Println("Code ran with argument", os.Args[1])
}
