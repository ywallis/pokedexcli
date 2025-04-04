package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)

	return split 
}

func main() {
	fmt.Println("Hello, World!")
}
