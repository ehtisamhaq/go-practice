package main

import "fmt"

func main() {

	i, j := 42, 2701

	fmt.Println("address of i:", &i, "address of j:", &j)

	p := &i                // point to i
	fmt.Println(p, *p)     // read i through the pointer
	fmt.Printf("%T\n", p) // the type of p is *int

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
