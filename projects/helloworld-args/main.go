package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Hello, you entered the following cli args: %v", args[1:])
}
