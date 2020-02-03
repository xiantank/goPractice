package main

import (
	"fmt"
	"os"
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
	if len(os.Args) < 3 {
		println("need at at least 2 arguments: main <operator> <number> [...numbers]")
		os.Exit(1)
	}

	operator := os.Args[1]
	operands := os.Args[2:]
	var opFunc MathOP
	switch operator {
	case "a", "+":
		opFunc = add
	case "m", "-":
		opFunc = minus
	case "d", "/":
		opFunc = multiply
	case "n", "*":
		opFunc = devide
	default:
		fmt.Println("not exist operator:", operator)
		os.Exit(2)
	}
	sum, err := strconv.Atoi(operands[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range operands[1:] {
		number, _ := strconv.Atoi(value)
		sum = opFunc(sum, number)
	}
	fmt.Println("sum: ", sum)
}
