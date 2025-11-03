package main


var name string = "Gopher"

func init() {
	name = "Go Programmer"
	age:= 42
	a, b := "hello", 2

	return {
		name, age,
		a,
		b,
	}
}


// block declaration
var (
	city    string = "San Francisco"
	country string = "USA"
)

// scope demonstration
var GlobalVar string = "I am a global var"

func ScopeDemo() string {
	localVar := "I am a local var"
	return localVar
}

