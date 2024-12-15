package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	fmt.Printf("Hello world\nos.Args: %v\nArguments: %v\n", args, args[1:])

	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i],"--")
			fmt.Println("Args: %v\n, Value: %v", args[i], args[i++])
	}
}
