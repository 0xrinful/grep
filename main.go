package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Flags struct {
	showLineNum bool
}

func getMatchHandler(flags Flags) func(string, int) {
	if flags.showLineNum {
		return func(line string, lineNum int) {
			fmt.Printf("%d:%s\n", lineNum, line)
		}
	}
	return func(line string, _ int) {
		fmt.Println(line)
	}
}

func main() {
	flags := Flags{}
	flag.BoolVar(&flags.showLineNum, "n", false, "print line numbers")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: grep PATTERN [FILE...]\n")
		os.Exit(2)
	}

	pattern := os.Args[1]
	handler := getMatchHandler(flags)

	var found bool
	var err error

	if len(args) > 1 {
		found, err = grepFile(pattern, args[1], handler)
	} else {
		found, err = grepReader(pattern, handler)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	if !found {
		os.Exit(1)
	}
}

func grepFile(pattern, filename string, handleMatch func(string, int)) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()
	return grep(pattern, file, handleMatch)
}

func grepReader(pattern string, handleMatch func(string, int)) (bool, error) {
	return grep(pattern, os.Stdin, handleMatch)
}

func grep(pattern string, r io.Reader, handleMatch func(line string, lineNum int)) (bool, error) {
	found := false
	sc := bufio.NewScanner(r)
	lineNum := 1
	for sc.Scan() {
		line := sc.Text()
		if match(line, pattern) {
			found = true
			handleMatch(line, lineNum)
		}
		lineNum++
	}
	return found, sc.Err()
}
