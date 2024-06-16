package main

import "fmt"

func main() {
	fmt.Println(Sum(2, 2))
}

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Times(a, b int) int {
	return a * b
}

func SumX(a, b int) int {
	return a + b + a
}
