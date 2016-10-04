package main

import (
	"./KeyCount"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var keyCount KeyCount.KeyCount
	keyCount = KeyCount.NewKeyCountArray()
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
