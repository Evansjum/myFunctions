package main

import "fmt"

func main() {
	input := "Hello, world! How are you?"
	fmt.Println(ReverseWords(input))

}

func ReverseWords(str string) []string {

	var result []string

	for j := len(str) - 1; j >= 0; j-- {
		result = append(result, string(str[j]))
	}

	return result

}
