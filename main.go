package main

import (
	"fmt"
	"os"

	"github.com/Telshin/advent-of-code-2025/day01"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	switch os.Args[1] {
	case "1":
		day01.Run()
	}
}
