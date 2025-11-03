package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

type User struct {
	ID   int
	Name string
}

func (u *User) Greet() string {
	return "Hello, " + u.Name + "!" + " Your ID is " + fmt.Sprint(u.ID)
}

func double(n *int) {
	*n = *n * 2
}

func main() {
	// 01 Hello World
	fmt.Println("Hello World!")

	// 02 Variables & Constants
	var name string = "Ehtisam"
	age := 22 // Short declaration
	const country = "Bangladesh"

	fmt.Println(name, age, country)

	// 03 Data Types
	/*
	   Int, Float, String, Bool
	   Arrays, Slices, Maps, Structs
	*/

	// 04 Control Structures
	// 4.1 If-Else
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// 4.2 Loops
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// 4.3 Switch
	day := "Monday"
	switch day {
	case "Thursday", "Friday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// 05 Functions
	sum := add(5, 7)
	fmt.Println("Sum:", sum)

	quotient, err := divide(10, 100.239847)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

	// 06 Structs & Methods
	user := User{ID: 1, Name: "Ehtisam"}
	fmt.Println(user.Greet())

	// 07 Pointers
	num := 100
	fmt.Println("Before:", num)
	double(&num)
	fmt.Println("After:", num)
}
