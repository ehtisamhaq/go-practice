package main

// operators

var a, b int = 10, 20

func operator() {
	var result int = a + b
	println(result)
}

func ArithmeticOperators() (int, int, int, int, int, int) {
	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	mod := a % b
	exp := 1
	for i := 0; i < b; i++ {
		exp *= a
	}
	return sum, diff, prod, quot, mod, exp
}

func RelationalOperators() (bool, bool, bool, bool, bool, bool) {
	eq := a == b
	neq := a != b
	gt := a > b
	lt := a < b
	gte := a >= b
	lte := a <= b
	return eq, neq, gt, lt, gte, lte
}

func LogicalOperators() (bool, bool, bool) {
	and := (a < b) && (a > 5)
	or := (a < b) || (a > 15)
	not := !(a < b)
	return and, or, not
}

func BitwiseOperators() (int, int, int, int, int, int) {
	and := a & b
	or := a | b
	xor := a ^ b
	andNot := a &^ b
	leftShift := a << 2
	rightShift := a >> 2
	return and, or, xor, andNot, leftShift, rightShift
}

func AssignmentOperators() (int, int, int, int, int, int) {
	c := a
	c += b
	c -= b
	c *= b
	c /= b
	c %= b
	return c, c, c, c, c, c
}

func MiscOperators() (int, int) {
	// address operator
	addr := &a
	// dereference operator
	val := *addr
	return addr, val
}	
