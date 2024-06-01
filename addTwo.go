package main

import "fmt"

func main() {
	isNegative(-12)
	input1 := 27
	input2 := 3
	fmt.Println(addTWo(input1, input2))
	fmt.Println(multiplication(input1, input2))
	isNegative(12)

}



func multiplication(n1, n2 int) int {
	multiply := n1 * n2
	return multiply

}

func isNegative(n int) {
	if n < 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
