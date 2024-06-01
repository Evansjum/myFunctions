//CODEWARS Write a function which converts the input string to uppercase.

package main

import "fmt"

func main() {

	input := " hello3   there2re  hello world"
	fmt.Println(MakeUpperCase(input))
}

func MakeUpperCase(str string) string {
	var result string
	for _, i := range str {
		if i >= 'a' && i <= 'z' {
			i -= 32
		}
		result += string(i)
	}

	return result
}

// func MakeUpperCase(str string) string {
// 	var upperStr string
// 	for i := 2; i < len(str)-2; i++ {
// 		if str[i] <= 'z' && str[i] >= 'a' {
// 			upperStr = upperStr + string(str[i]-'a'+'A')
// 		} else {
// 			upperStr = upperStr + string(str[i])
// 		}
// 	}
// 	return upperStr
// }
