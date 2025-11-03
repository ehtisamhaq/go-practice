// statements.go
package main

import "fmt"

func checkEvenOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}


func main() {
	fmt.Println(checkEvenOdd(10)) // Output: Even
	fmt.Println(checkEvenOdd(7))  // Output: Odd
}