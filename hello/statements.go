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

// switch statement
func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}
