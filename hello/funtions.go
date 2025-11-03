package main

// functions.go

import "fmt"

func greet(name string) string {
	return "Hello, " + name + "!"
}

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(greet("Alice"))
	fmt.Println("Sum:", add(5, 7))
}