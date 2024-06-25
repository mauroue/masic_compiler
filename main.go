package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please select a file as a parameter.")
	}

	path := args[1]

	file, err := os.Open(path)
	errHandler(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	fmt.Println(line)
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}
