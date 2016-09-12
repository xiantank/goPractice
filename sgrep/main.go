package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	pattern := args[0]
	filename := args[1]
	fmt.Println(pattern, filename)

	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	myReader := bufio.NewReader(fi)
	scanner := bufio.NewScanner(myReader)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error", err)
	}
}
