package main

import (
	"fmt"
)

func main() {
	var i int
	var j int
	for i = 2; i <= 9; i++ {
		for j = 1; j <= 9; j++ {
			fmt.Printf("%3d*%d = %2d", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
