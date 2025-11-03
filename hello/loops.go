package main

// loops.go

import "fmt"

func main() {
	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("For Loop Iteration:", i)
	}

	// while-like loop
	j := 1
	for j <= 5 {
		fmt.Println("While-like Loop Iteration:", j)
		j++
	}

	// infinite loop with break
	k := 1
	for {
		if k > 5 {
			break
		}
		fmt.Println("Infinite Loop Iteration:", k)
		k++
	}

	// range loop
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	
}

func sumOfFirstN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func nestedLoopsExample() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}

func loopWithContinue() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd Number:", i)
	}
}	