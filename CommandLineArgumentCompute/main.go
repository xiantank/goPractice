package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type MathOP func(int, int) int

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func devide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(reflect.TypeOf(os.Args))
	operator := os.Args[1]
	var opFunc MathOP
	switch operator {
	case "a":
		opFunc = add
	case "m":
		opFunc = minus
	case "d":
		opFunc = multiply
	case "n":
		opFunc = devide
	}
	sum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range os.Args[3:] {
		number, _ := strconv.Atoi(value)
		sum = opFunc(sum, number)
	}
	fmt.Println("sum: ", sum)
}
