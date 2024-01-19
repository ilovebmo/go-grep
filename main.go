package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	filter := os.Args[1:]
	for scanner.Scan() {
		name := scanner.Text()
		if strings.Contains(name, strings.Join(filter, "")) {
			fmt.Println(name)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
