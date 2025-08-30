package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: grep PATTERN [FILE...]\n")
		os.Exit(2)
	}

	pattern := os.Args[1]
	if len(os.Args) > 2 {
		grepFile(pattern, os.Args[2])
	} else {
		grepReader(pattern)
	}
}

func grepFile(pattern, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func grepReader(pattern string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
