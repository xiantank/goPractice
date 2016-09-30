package main

import (
	"./KeyCount"
	"bufio"
	"fmt"
	"os"
)

func main() {
	keyCount := KeyCount.NewKeyCount()
	finName := os.Args[1]
	fin, err := os.Open(finName)
	if err != nil {
		panic("open fail")
	}
	myReader := bufio.NewReader(fin)
	scanner := bufio.NewScanner(myReader)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		keyCount.Add(line)
	}
	fmt.Println(keyCount.Total())
	keyCount.Print()

}
